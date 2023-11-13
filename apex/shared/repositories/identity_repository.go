package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/apex/shared/entities"
	apexcustomeridentity "github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomeridentity"
	"github.com/sigmasee/sigmasee/apex/shared/mappers"
	"github.com/sigmasee/sigmasee/apex/shared/models"
	"github.com/sigmasee/sigmasee/shared/enterprise/dataloader/cache"
)

type IdentityRepository interface {
	Upsert(
		ctx context.Context,
		tx *entities.Tx,
		customerID string,
		identity *models.Identity,
		updateFields bool) error

	DeleteByIds(
		ctx context.Context,
		tx *entities.Tx,
		ids []string) error

	GetById(
		ctx context.Context,
		tx *entities.Tx,
		useCache bool,
		id string) (*models.Identity, error)
}

type identityRepository struct {
	entgoClient EntgoClient
	mapper      mappers.Mapper

	byIdCache      *cache.Cache[DatabaseCacheKey[string], *models.Identity]
	byIdDataloader *dataloader.Loader[DatabaseCacheKey[string], *models.Identity]
}

func NewIdentityRepository(
	entgoClient EntgoClient,
	mapper mappers.Mapper) (IdentityRepository, error) {
	service := &identityRepository{
		entgoClient: entgoClient,
		mapper:      mapper,
	}

	var err error

	service.byIdCache, err = cache.NewCache[DatabaseCacheKey[string], *models.Identity](10240)
	if err != nil {
		return nil, err
	}

	service.byIdDataloader = dataloader.NewBatchedLoader(
		service.getByIdBatchLoader,
		dataloader.WithCache[DatabaseCacheKey[string], *models.Identity](service.byIdCache))

	return service, nil
}

func (s *identityRepository) Upsert(
	ctx context.Context,
	tx *entities.Tx,
	customerID string,
	identity *models.Identity,
	updateFields bool) error {
	now := time.Now().UTC()

	query := s.entgoClient.
		GetApexCustomerIdentityClient(tx).
		Create().
		SetID(identity.ID).
		SetCreatedAt(now).
		SetModifiedAt(now).
		SetNillableDeletedAt(nil).
		SetNillableEmail(identity.Email).
		SetNillableEmailVerified(identity.EmailVerified).
		SetCustomerID(customerID).
		OnConflictColumns(apexcustomeridentity.FieldID)

	if updateFields {
		query = query.Update(func(upsert *entities.ApexCustomerIdentityUpsert) {
			upsert.
				UpdateModifiedAt().
				UpdateDeletedAt()

			if identity.Email == nil {
				upsert.ClearEmail()
			} else {
				upsert.UpdateEmail()
			}

			if identity.EmailVerified == nil {
				upsert.ClearEmailVerified()
			} else {
				upsert.UpdateEmailVerified()
			}
		})
	} else {
		query = query.Ignore()
	}

	return query.Exec(ctx)
}

func (s *identityRepository) DeleteByIds(
	ctx context.Context,
	tx *entities.Tx,
	ids []string) error {
	now := time.Now().UTC()

	return s.entgoClient.
		GetApexCustomerIdentityClient(tx).
		Update().
		SetModifiedAt(now).
		SetDeletedAt(now).
		Where(apexcustomeridentity.IDIn(ids...)).
		Exec(ctx)
}

func (s *identityRepository) GetById(
	ctx context.Context,
	tx *entities.Tx,
	useCache bool,
	id string) (*models.Identity, error) {
	if useCache {
		return s.byIdDataloader.Load(
			ctx,
			DatabaseCacheKey[string]{
				Ctx:         ctx,
				Transaction: tx,
				Key:         id,
			})()
	}

	identity, err := s.entgoClient.
		GetApexCustomerIdentityClient(tx).
		Query().
		Where(apexcustomeridentity.IDEQ(id)).
		Only(ctx)
	if entities.IsNotFound(err) {
		return nil, nil
	} else if entities.IsNotSingular(err) {
		return nil, fmt.Errorf("multiple identitys found with given id: %s", id)
	} else if err != nil {
		return nil, err
	}

	return s.mapper.CustomerIdentityEntityToIdentity(identity), nil
}

func (s *identityRepository) getByIdBatchLoader(
	ctx context.Context,
	keys []DatabaseCacheKey[string]) []*dataloader.Result[*models.Identity] {
	return slices.Map(keys, func(item DatabaseCacheKey[string]) *dataloader.Result[*models.Identity] {
		result, err := s.GetById(item.Ctx, item.Transaction, false, item.Key)
		if err != nil {
			return &dataloader.Result[*models.Identity]{Error: err}
		}

		return &dataloader.Result[*models.Identity]{Data: result}
	})
}
