package services

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/sigmasee/sigmasee/customer/shared/models"
	"github.com/sigmasee/sigmasee/customer/shared/publishers"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	enterprisecontext "github.com/sigmasee/sigmasee/shared/enterprise/context"
	"github.com/sigmasee/sigmasee/shared/enterprise/converters"
	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
	"go.uber.org/zap"
)

type CustomerService interface {
	Upsert(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	GetMe(ctx context.Context) (*models.Customer, error)
}

type customerService struct {
	logger                     *zap.SugaredLogger
	contextHelper              enterprisecontext.ContextHelper
	entgoClient                repositories.EntgoClient
	intercomConfig             configuration.IntercomConfig
	customerRepository         repositories.CustomerRepository
	customerSettingsRepository repositories.CustomerSettingsRepository
	customerOutboxPublisher    publishers.CustomerOutboxPublisher
}

func NewCustomerService(
	logger *zap.SugaredLogger,
	contextHelper enterprisecontext.ContextHelper,
	entgoClient repositories.EntgoClient,
	intercomConfig configuration.IntercomConfig,
	customerRepository repositories.CustomerRepository,
	customerSettingsRepository repositories.CustomerSettingsRepository,
	customerOutboxPublisher publishers.CustomerOutboxPublisher) (CustomerService, error) {
	return &customerService{
		logger:                     logger,
		contextHelper:              contextHelper,
		entgoClient:                entgoClient,
		intercomConfig:             intercomConfig,
		customerRepository:         customerRepository,
		customerSettingsRepository: customerSettingsRepository,
		customerOutboxPublisher:    customerOutboxPublisher,
	}, nil
}

func (s *customerService) Upsert(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	identity := customer.Identities[0]
	if len(identity.ID) == 0 {
		return nil, fmt.Errorf("identity.ID is empty")
	}

	existingCustomer, err := s.customerRepository.GetByVerifiableToken(ctx, nil, false, identity.ID)
	if err != nil {
		return nil, err
	}

	if existingCustomer == nil && identity.Email != nil && len(*identity.Email) > 0 {
		existingCustomer, err = s.customerRepository.GetByEmail(ctx, nil, false, *identity.Email)
		if err != nil {
			return nil, err
		}

		if existingCustomer != nil {
			customer.ID = existingCustomer.ID
		}
	} else {
		customer.ID = existingCustomer.ID
	}

	tx, err := s.entgoClient.CreateTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			if rollbackErr := s.entgoClient.RollbackTransaction(tx); rollbackErr != nil {
				s.logger.Errorf("Failed to rollback transaction. Error: %v", rollbackErr)
			}
		}
	}()

	if err = s.customerRepository.Upsert(ctx, tx, customer, true); err != nil {
		return nil, err
	}

	if err = s.customerSettingsRepository.Upsert(
		ctx,
		tx,
		customer.ID,
		&models.CustomerSettings{},
		false); err != nil {
		return nil, err
	}

	upsertedCustomer, err := s.customerRepository.GetById(ctx, tx, false, customer.ID)
	if err != nil {
		return nil, err
	}

	if err = s.customerOutboxPublisher.Publish(ctx, tx, []*models.Customer{upsertedCustomer}); err != nil {
		return nil, err
	}

	if err = s.entgoClient.CommitTransaction(tx); err != nil {
		return nil, err
	}

	return upsertedCustomer, nil
}

func (s *customerService) GetMe(ctx context.Context) (*models.Customer, error) {
	verifiableToken := s.contextHelper.GetVerifiableToken(ctx)
	existingCustomer, err := s.customerRepository.GetByVerifiableToken(ctx, nil, true, verifiableToken)
	if err != nil {
		return nil, err
	}

	if existingCustomer == nil {
		return nil, errors.ErrCustomerNotFound
	}

	if len(s.intercomConfig.Secret) > 0 {
		mac := hmac.New(sha256.New, []byte(s.intercomConfig.Secret))
		mac.Write([]byte(existingCustomer.ID))
		existingCustomer.IntercomHash = converters.StrToPointerStr(hex.EncodeToString(mac.Sum(nil)))
	}

	return existingCustomer, nil
}
