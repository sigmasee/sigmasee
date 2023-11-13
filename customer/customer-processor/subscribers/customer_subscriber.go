package subscribers

import (
	"context"
	"time"

	"github.com/sigmasee/sigmasee/customer/customer-processor/mappers"
	"github.com/sigmasee/sigmasee/customer/shared/models"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	customerv1 "github.com/sigmasee/sigmasee/shared/clients/events/sigmasee/customer/v1"
	enterpriseconfiguration "github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/random"
	"go.uber.org/zap"
)

type customerSubscriber struct {
	logger                     *zap.SugaredLogger
	appConfig                  enterpriseconfiguration.AppConfig
	customerRepository         repositories.CustomerRepository
	customerSettingsRepository repositories.CustomerSettingsRepository
	mapper                     mappers.Mapper
	randomHelper               random.RandomHelper
}

func NewCustomerSubscriber(
	logger *zap.SugaredLogger,
	appConfig enterpriseconfiguration.AppConfig,
	customerRepository repositories.CustomerRepository,
	customerSettingsRepository repositories.CustomerSettingsRepository,
	mapper mappers.Mapper,
	randomHelper random.RandomHelper) (customerv1.Subscriber, error) {
	return &customerSubscriber{
		logger:                     logger,
		appConfig:                  appConfig,
		customerRepository:         customerRepository,
		customerSettingsRepository: customerSettingsRepository,
		mapper:                     mapper,
		randomHelper:               randomHelper,
	}, nil
}

func (s *customerSubscriber) Handle(
	ctx context.Context,
	topic string,
	key []byte,
	headers map[string][]byte,
	event *customerv1.Event) (err error) {
	start := time.Now()
	defer func(start time.Time) {
		s.logger.Infof("Handle(customerv1.Event) - Event Type: %d, Execution time: %s. Error: %v", event.Metadata.Type, time.Since(start), err)
	}(start)

	if event.Metadata.DomainSource == s.appConfig.DomainSource {
		// Event raised previously by the this domain, ignoring it.
		return
	}

	customer, err := s.mapper.CustomerEventToCustomer(event)
	if err != nil {
		return
	}

	existingCustomer, err := s.customerRepository.GetById(ctx, nil, false, customer.ID)
	if err != nil {
		return
	}

	if existingCustomer != nil && existingCustomer.ModifiedAt.After(customer.ModifiedAt) {
		s.logger.Info("Ignoring Customer event. Event timestamp is older that what is already processed.")

		return
	}

	if err = s.customerRepository.Upsert(ctx, nil, customer, true); err != nil {
		return
	}

	id, err := s.randomHelper.Generate()
	if err != nil {
		return
	}

	err = s.customerSettingsRepository.Upsert(
		ctx,
		nil,
		customer.ID,
		&models.CustomerSettings{
			ID: id,
		},
		false)

	return
}
