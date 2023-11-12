package repositories

import (
	"context"

	"github.com/sigmasee/sigmasee/apex/shared/entities"
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

	GetApexCustomerClient(tx *entities.Tx) *entities.ApexCustomerClient
	GetApexCustomerIdentityClient(tx *entities.Tx) *entities.ApexCustomerIdentityClient
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

func (s *entgoClient) GetApexCustomerClient(tx *entities.Tx) *entities.ApexCustomerClient {
	if tx == nil {
		return s.GetClient().ApexCustomer
	} else {
		return tx.ApexCustomer
	}
}

func (s *entgoClient) GetApexCustomerIdentityClient(tx *entities.Tx) *entities.ApexCustomerIdentityClient {
	if tx == nil {
		return s.GetClient().ApexCustomerIdentity
	} else {
		return tx.ApexCustomerIdentity
	}
}
