package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/apex/shared/entities"
	apexcustomerentity "github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomer"
	apexcustomeridentityentity "github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomeridentity"
	"github.com/sigmasee/sigmasee/apex/shared/mappers"
	"github.com/sigmasee/sigmasee/apex/shared/models"
	"github.com/sigmasee/sigmasee/shared/enterprise/dataloader/cache"
	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
)

type CustomerRepository interface {
	Upsert(
		ctx context.Context,
		tx *entities.Tx,
		customer *models.Customer,
		updateFields bool) error

	GetById(
		ctx context.Context,
		tx *entities.Tx,
		useCache bool,
		id string) (*models.Customer, error)

	GetByVerifiableToken(
		ctx context.Context,
		tx *entities.Tx,
		useCache bool,
		verifiableToken string) (*models.Customer, error)

	GetByEmail(
		ctx context.Context,
		tx *entities.Tx,
		useCache bool,
		email string) (*models.Customer, error)
}

type customerRepository struct {
	entgoClient        EntgoClient
	mapper             mappers.Mapper
	identityRepository IdentityRepository

	byIdCache      *cache.Cache[DatabaseCacheKey[string], *models.Customer]
	byIdDataloader *dataloader.Loader[DatabaseCacheKey[string], *models.Customer]

	byVerifiableTokenCache      *cache.Cache[DatabaseCacheKey[string], *models.Customer]
	byVerifiableTokenDataloader *dataloader.Loader[DatabaseCacheKey[string], *models.Customer]

	byEmailCache      *cache.Cache[DatabaseCacheKey[string], *models.Customer]
	byEmailDataloader *dataloader.Loader[DatabaseCacheKey[string], *models.Customer]
}

func NewCustomerRepository(
	entgoClient EntgoClient,
	mapper mappers.Mapper,
	identityRepository IdentityRepository) (CustomerRepository, error) {
	service := &customerRepository{
		entgoClient:        entgoClient,
		mapper:             mapper,
		identityRepository: identityRepository,
	}

	var err error

	service.byIdCache, err = cache.NewCache[DatabaseCacheKey[string], *models.Customer](10240)
	if err != nil {
		return nil, err
	}

	service.byIdDataloader = dataloader.NewBatchedLoader(
		service.getByIdBatchLoader,
		dataloader.WithCache[DatabaseCacheKey[string], *models.Customer](service.byIdCache))

	service.byVerifiableTokenCache, err = cache.NewCache[DatabaseCacheKey[string], *models.Customer](10240)
	if err != nil {
		return nil, err
	}

	service.byVerifiableTokenDataloader = dataloader.NewBatchedLoader(
		service.getByVerifiableTokenBatchLoader,
		dataloader.WithCache[DatabaseCacheKey[string], *models.Customer](service.byVerifiableTokenCache))

	service.byEmailCache, err = cache.NewCache[DatabaseCacheKey[string], *models.Customer](10240)
	if err != nil {
		return nil, err
	}

	service.byEmailDataloader = dataloader.NewBatchedLoader(
		service.getByEmailBatchLoader,
		dataloader.WithCache[DatabaseCacheKey[string], *models.Customer](service.byEmailCache))

	return service, nil
}

func (s *customerRepository) Upsert(
	ctx context.Context,
	tx *entities.Tx,
	customer *models.Customer,
	updateFields bool) error {
	now := time.Now().UTC()

	query := s.entgoClient.GetApexCustomerClient(tx).
		Create().
		SetID(customer.ID).
		SetCreatedAt(now).
		SetModifiedAt(now).
		SetEventRaisedAt(customer.EventRaisedAt).
		SetNillableName(customer.Name).
		SetNillableGivenName(customer.GivenName).
		SetNillableMiddleName(customer.MiddleName).
		SetNillableFamilyName(customer.FamilyName).
		SetNillablePhotoURL(customer.PhotoURL).
		SetNillablePhotoURL24(customer.PhotoURL24).
		SetNillablePhotoURL32(customer.PhotoURL32).
		SetNillablePhotoURL48(customer.PhotoURL48).
		SetNillablePhotoURL72(customer.PhotoURL72).
		SetNillablePhotoURL192(customer.PhotoURL192).
		SetNillablePhotoURL512(customer.PhotoURL512).
		OnConflictColumns(apexcustomerentity.FieldID)

	if updateFields {
		query = query.Update(func(upsert *entities.ApexCustomerUpsert) {
			upsert.
				UpdateModifiedAt().
				UpdateEventRaisedAt()

			if customer.Name == nil {
				upsert.ClearName()
			} else {
				upsert.UpdateName()
			}

			if customer.GivenName == nil {
				upsert.ClearGivenName()
			} else {
				upsert.UpdateGivenName()
			}

			if customer.MiddleName == nil {
				upsert.ClearMiddleName()
			} else {
				upsert.UpdateMiddleName()
			}

			if customer.FamilyName == nil {
				upsert.ClearFamilyName()
			} else {
				upsert.UpdateFamilyName()
			}

			if customer.PhotoURL == nil {
				upsert.ClearPhotoURL()
			} else {
				upsert.UpdatePhotoURL()
			}

			if customer.PhotoURL24 == nil {
				upsert.ClearPhotoURL24()
			} else {
				upsert.UpdatePhotoURL24()
			}

			if customer.PhotoURL32 == nil {
				upsert.ClearPhotoURL32()
			} else {
				upsert.UpdatePhotoURL32()
			}

			if customer.PhotoURL48 == nil {
				upsert.ClearPhotoURL48()
			} else {
				upsert.UpdatePhotoURL48()
			}

			if customer.PhotoURL72 == nil {
				upsert.ClearPhotoURL72()
			} else {
				upsert.UpdatePhotoURL72()
			}

			if customer.PhotoURL192 == nil {
				upsert.ClearPhotoURL192()
			} else {
				upsert.UpdatePhotoURL192()
			}

			if customer.PhotoURL512 == nil {
				upsert.ClearPhotoURL512()
			} else {
				upsert.UpdatePhotoURL512()
			}
		})
	} else {
		query = query.Ignore()
	}

	if err := query.Exec(ctx); err != nil {
		return err
	}

	if !updateFields {
		return nil
	}

	upsertedCustomer, err := s.GetById(ctx, tx, false, customer.ID)
	if err != nil {
		return err
	}

	identitiesToDelete := slices.Filter(upsertedCustomer.Identities, func(item *models.Identity) bool {
		return slices.All(customer.Identities, func(identity *models.Identity) bool {
			return item.ID != identity.ID
		})
	})

	if err = s.identityRepository.DeleteByIds(
		ctx,
		tx,
		slices.Map(identitiesToDelete, func(item *models.Identity) string {
			return item.ID
		})); err != nil {
		return err
	}

	identitiesToUpsert := slices.Filter(customer.Identities, func(item *models.Identity) bool {
		return slices.All(upsertedCustomer.Identities, func(identity *models.Identity) bool {
			return item.ID != identity.ID
		}) || slices.Any(upsertedCustomer.Identities, func(identity *models.Identity) bool {
			return item.ID == identity.ID
		})
	})

	result := slices.Map(identitiesToUpsert, func(item *models.Identity) error {
		return s.identityRepository.Upsert(ctx, tx, upsertedCustomer.ID, item, true)
	})

	return errors.ReduceErrors(result)
}

func (s *customerRepository) GetById(
	ctx context.Context,
	tx *entities.Tx,
	useCache bool,
	id string) (*models.Customer, error) {
	if useCache {
		return s.byIdDataloader.Load(
			ctx,
			DatabaseCacheKey[string]{
				Ctx:         ctx,
				Transaction: tx,
				Key:         id,
			})()
	}

	customer, err := s.entgoClient.
		GetApexCustomerClient(tx).
		Query().
		Where(apexcustomerentity.IDEQ(id)).
		WithIdentities().
		Only(ctx)
	if entities.IsNotFound(err) {
		return nil, nil
	} else if entities.IsNotSingular(err) {
		return nil, fmt.Errorf("multiple customers found with given Id: %s", id)
	} else if err != nil {
		return nil, err
	}

	return s.mapper.CustomerEntityToCustomer(customer), nil
}

func (s *customerRepository) GetByVerifiableToken(
	ctx context.Context,
	tx *entities.Tx,
	useCache bool,
	verifiableToken string) (*models.Customer, error) {
	if useCache {
		return s.byVerifiableTokenDataloader.Load(
			ctx,
			DatabaseCacheKey[string]{
				Ctx:         ctx,
				Transaction: tx,
				Key:         verifiableToken,
			})()
	}

	customer, err := s.entgoClient.
		GetApexCustomerClient(tx).
		Query().
		Where(apexcustomerentity.HasIdentitiesWith(apexcustomeridentityentity.IDEQ(verifiableToken))).
		WithIdentities().
		Only(ctx)
	if entities.IsNotFound(err) {
		return nil, nil
	} else if entities.IsNotSingular(err) {
		return nil, fmt.Errorf("multiple customers found with given id: %s", verifiableToken)
	} else if err != nil {
		return nil, err
	}

	return s.mapper.CustomerEntityToCustomer(customer), nil
}

func (s *customerRepository) GetByEmail(
	ctx context.Context,
	tx *entities.Tx,
	useCache bool,
	email string) (*models.Customer, error) {
	if useCache {
		return s.byEmailDataloader.Load(
			ctx,
			DatabaseCacheKey[string]{
				Ctx:         ctx,
				Transaction: tx,
				Key:         email,
			})()
	}

	customer, err := s.entgoClient.
		GetApexCustomerClient(tx).
		Query().
		Where(apexcustomerentity.HasIdentitiesWith(apexcustomeridentityentity.EmailEQ(email))).
		WithIdentities().
		Only(ctx)
	if entities.IsNotFound(err) {
		return nil, nil
	} else if entities.IsNotSingular(err) {
		return nil, fmt.Errorf("multiple customers found with given Email: %s", email)
	} else if err != nil {
		return nil, err
	}

	return s.mapper.CustomerEntityToCustomer(customer), nil
}

func (s *customerRepository) getByIdBatchLoader(
	ctx context.Context,
	keys []DatabaseCacheKey[string]) []*dataloader.Result[*models.Customer] {
	return slices.Map(keys, func(item DatabaseCacheKey[string]) *dataloader.Result[*models.Customer] {
		result, err := s.GetById(item.Ctx, item.Transaction, false, item.Key)
		if err != nil {
			return &dataloader.Result[*models.Customer]{Error: err}
		}

		return &dataloader.Result[*models.Customer]{Data: result}
	})
}

func (s *customerRepository) getByVerifiableTokenBatchLoader(
	ctx context.Context,
	keys []DatabaseCacheKey[string]) []*dataloader.Result[*models.Customer] {
	return slices.Map(keys, func(item DatabaseCacheKey[string]) *dataloader.Result[*models.Customer] {
		result, err := s.GetByVerifiableToken(item.Ctx, item.Transaction, false, item.Key)
		if err != nil {
			return &dataloader.Result[*models.Customer]{Error: err}
		}

		return &dataloader.Result[*models.Customer]{Data: result}
	})
}

func (s *customerRepository) getByEmailBatchLoader(
	ctx context.Context,
	keys []DatabaseCacheKey[string]) []*dataloader.Result[*models.Customer] {
	return slices.Map(keys, func(item DatabaseCacheKey[string]) *dataloader.Result[*models.Customer] {
		result, err := s.GetByEmail(item.Ctx, item.Transaction, false, item.Key)
		if err != nil {
			return &dataloader.Result[*models.Customer]{Error: err}
		}

		return &dataloader.Result[*models.Customer]{Data: result}
	})
}
