package repositories

import (
	"context"

	"github.com/sigmasee/sigmasee/customer/shared/entities"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"go.uber.org/zap"
)

type DatabaseCacheKey[T comparable] struct {
	Ctx         context.Context
	Transaction *entities.Tx
	Key         T
}

type EntgoClient interface {
	Close()
	GetClient() *entities.Client
	CreateTransaction(ctx context.Context) (*entities.Tx, error)
	RollbackTransaction(tx *entities.Tx) error
	CommitTransaction(tx *entities.Tx) error

	GetCustomerClient(tx *entities.Tx) *entities.CustomerClient
	GetCustomerOutboxClient(tx *entities.Tx) *entities.CustomerOutboxClient
	GetIdentityClient(tx *entities.Tx) *entities.IdentityClient
	GetCustomerSettingClient(tx *entities.Tx) *entities.CustomerSettingClient
}

type entgoClient struct {
	logger   *zap.SugaredLogger
	config   database.DatabaseConfig
	database database.Database
	client   *entities.Client
}

func NewEntgoClient(
	logger *zap.SugaredLogger,
	config database.DatabaseConfig,
	database database.Database) (EntgoClient, error) {

	driver := database.GetDriver()

	return &entgoClient{
		logger:   logger,
		config:   config,
		database: database,
		client:   entities.NewClient(entities.Driver(driver)),
	}, nil
}

func (s *entgoClient) Close() {
	if s.client != nil {
		if err := s.client.Close(); err != nil {
			s.logger.Errorf("Failed to close entgo client. Error: %v", err)
		}

		s.client = nil
	}

	s.database.Close()
}

func (s *entgoClient) GetClient() *entities.Client {
	client := s.client
	if !s.config.Debug {
		return client
	}

	return client.Debug()
}

func (s *entgoClient) CreateTransaction(ctx context.Context) (*entities.Tx, error) {
	return s.GetClient().Tx(ctx)
}

func (s *entgoClient) RollbackTransaction(tx *entities.Tx) error {
	return tx.Rollback()
}

func (s *entgoClient) CommitTransaction(tx *entities.Tx) error {
	return tx.Commit()
}

func (s *entgoClient) GetCustomerClient(tx *entities.Tx) *entities.CustomerClient {
	if tx == nil {
		return s.GetClient().Customer
	} else {
		return tx.Customer
	}
}

func (s *entgoClient) GetCustomerOutboxClient(tx *entities.Tx) *entities.CustomerOutboxClient {
	if tx == nil {
		return s.GetClient().CustomerOutbox
	} else {
		return tx.CustomerOutbox
	}
}

func (s *entgoClient) GetIdentityClient(tx *entities.Tx) *entities.IdentityClient {
	if tx == nil {
		return s.GetClient().Identity
	} else {
		return tx.Identity
	}
}

func (s *entgoClient) GetCustomerSettingClient(tx *entities.Tx) *entities.CustomerSettingClient {
	if tx == nil {
		return s.GetClient().CustomerSetting
	} else {
		return tx.CustomerSetting
	}
}
