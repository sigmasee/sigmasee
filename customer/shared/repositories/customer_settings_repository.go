package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/customer/shared/entities"
	customerentity "github.com/sigmasee/sigmasee/customer/shared/entities/customer"
	customersettingentity "github.com/sigmasee/sigmasee/customer/shared/entities/customersetting"
	"github.com/sigmasee/sigmasee/customer/shared/mappers"
	"github.com/sigmasee/sigmasee/customer/shared/models"
	"github.com/sigmasee/sigmasee/shared/enterprise/dataloader/cache"
)

type CustomerSettingsRepository interface {
	Upsert(
		ctx context.Context,
		tx *entities.Tx,
		customerID string,
		customerSettings *models.CustomerSettings,
		updateFields bool) error

	GetById(
		ctx context.Context,
		tx *entities.Tx,
		useCache bool,
		id string) (*models.CustomerSettings, error)

	GetByCustomerId(
		ctx context.Context,
		tx *entities.Tx,
		useCache bool,
		customerID string) (*models.CustomerSettings, error)
}

type customerSettingsRepository struct {
	entgoClient EntgoClient
	mapper      mappers.Mapper

	byIdCache      *cache.Cache[DatabaseCacheKey[string], *models.CustomerSettings]
	byIdDataloader *dataloader.Loader[DatabaseCacheKey[string], *models.CustomerSettings]

	byCustomerIdCache      *cache.Cache[DatabaseCacheKey[string], *models.CustomerSettings]
	byCustomerIdDataloader *dataloader.Loader[DatabaseCacheKey[string], *models.CustomerSettings]
}

func NewCustomerSettingsRepository(
	entgoClient EntgoClient,
	mapper mappers.Mapper) (CustomerSettingsRepository, error) {
	service := &customerSettingsRepository{
		entgoClient: entgoClient,
		mapper:      mapper,
	}

	var err error

	service.byIdCache, err = cache.NewCache[DatabaseCacheKey[string], *models.CustomerSettings](10240)
	if err != nil {
		return nil, err
	}

	service.byIdDataloader = dataloader.NewBatchedLoader(
		service.getByIdBatchLoader,
		dataloader.WithCache[DatabaseCacheKey[string], *models.CustomerSettings](service.byIdCache))

	service.byCustomerIdCache, err = cache.NewCache[DatabaseCacheKey[string], *models.CustomerSettings](10240)
	if err != nil {
		return nil, err
	}

	service.byCustomerIdDataloader = dataloader.NewBatchedLoader(
		service.getByCustomerIdBatchLoader,
		dataloader.WithCache[DatabaseCacheKey[string], *models.CustomerSettings](service.byCustomerIdCache))

	return service, nil
}

func (s *customerSettingsRepository) Upsert(
	ctx context.Context,
	tx *entities.Tx,
	customerID string,
	customerSettings *models.CustomerSettings,
	updateFields bool) error {
	now := time.Now().UTC()

	query := s.entgoClient.
		GetCustomerSettingClient(tx).
		Create().
		SetID(customerSettings.ID).
		SetCreatedAt(now).
		SetModifiedAt(now).
		SetNillableDeletedAt(nil).
		SetCustomerID(customerID).
		OnConflictColumns(customersettingentity.CustomerColumn)

	if updateFields {
		query = query.Update(func(upsert *entities.CustomerSettingUpsert) {
			upsert.
				UpdateModifiedAt()
		})

	} else {
		query = query.Ignore()
	}

	return query.Exec(ctx)
}

func (s *customerSettingsRepository) GetById(
	ctx context.Context,
	tx *entities.Tx,
	useCache bool,
	id string) (*models.CustomerSettings, error) {
	if useCache {
		return s.byIdDataloader.Load(
			ctx,
			DatabaseCacheKey[string]{
				Ctx:         ctx,
				Transaction: tx,
				Key:         id,
			})()
	}

	customerSettings, err := s.entgoClient.
		GetCustomerSettingClient(tx).
		Query().
		Where(customersettingentity.IDEQ(id)).
		Only(ctx)
	if entities.IsNotFound(err) {
		return nil, nil
	} else if entities.IsNotSingular(err) {
		return nil, fmt.Errorf("multiple customer settingss found with given id: %s", id)
	} else if err != nil {
		return nil, err
	}

	return s.mapper.CustomerSettingsEntityToCustomerSettings(customerSettings), nil
}

func (s *customerSettingsRepository) GetByCustomerId(
	ctx context.Context,
	tx *entities.Tx,
	useCache bool,
	customerID string) (*models.CustomerSettings, error) {
	if useCache {
		return s.byCustomerIdDataloader.Load(
			ctx,
			DatabaseCacheKey[string]{
				Ctx:         ctx,
				Transaction: tx,
				Key:         customerID,
			})()
	}

	customerSettings, err := s.entgoClient.
		GetCustomerSettingClient(tx).
		Query().
		Where(customersettingentity.HasCustomerWith(customerentity.IDEQ(customerID))).
		Only(ctx)
	if entities.IsNotFound(err) {
		return nil, nil
	} else if entities.IsNotSingular(err) {
		return nil, fmt.Errorf("multiple customer settingss found with given customer id: %s", customerID)
	} else if err != nil {
		return nil, err
	}

	return s.mapper.CustomerSettingsEntityToCustomerSettings(customerSettings), nil
}

func (s *customerSettingsRepository) getByIdBatchLoader(
	ctx context.Context,
	keys []DatabaseCacheKey[string]) []*dataloader.Result[*models.CustomerSettings] {
	return slices.Map(keys, func(item DatabaseCacheKey[string]) *dataloader.Result[*models.CustomerSettings] {
		result, err := s.GetById(item.Ctx, item.Transaction, false, item.Key)
		if err != nil {
			return &dataloader.Result[*models.CustomerSettings]{Error: err}
		}

		return &dataloader.Result[*models.CustomerSettings]{Data: result}
	})
}

func (s *customerSettingsRepository) getByCustomerIdBatchLoader(
	ctx context.Context,
	keys []DatabaseCacheKey[string]) []*dataloader.Result[*models.CustomerSettings] {
	return slices.Map(keys, func(item DatabaseCacheKey[string]) *dataloader.Result[*models.CustomerSettings] {
		result, err := s.GetByCustomerId(item.Ctx, item.Transaction, false, item.Key)
		if err != nil {
			return &dataloader.Result[*models.CustomerSettings]{Error: err}
		}

		return &dataloader.Result[*models.CustomerSettings]{Data: result}
	})
}
