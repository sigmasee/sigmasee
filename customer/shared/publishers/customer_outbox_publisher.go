package publishers

import (
	"context"
	"time"

	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/customer/shared/entities"
	"github.com/sigmasee/sigmasee/customer/shared/mappers"
	"github.com/sigmasee/sigmasee/customer/shared/models"
	"github.com/sigmasee/sigmasee/customer/shared/outbox"
	customereventv1 "github.com/sigmasee/sigmasee/shared/clients/events/sigmasee/customer/v1"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	enterprisecontext "github.com/sigmasee/sigmasee/shared/enterprise/context"
	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
	"github.com/sigmasee/sigmasee/shared/enterprise/random"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerOutboxPublisher interface {
	Publish(
		ctx context.Context,
		tx *entities.Tx,
		customers []*models.Customer) error
}

type customerOutboxPublisher struct {
	logger          *zap.SugaredLogger
	appConfig       configuration.AppConfig
	contextHelper   enterprisecontext.ContextHelper
	outboxPublisher outbox.OutboxPublisher
	mapper          mappers.Mapper
	randomHelper    random.RandomHelper
}

func NewCustomerOutboxPublisher(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	contextHelper enterprisecontext.ContextHelper,
	outboxPublisher outbox.OutboxPublisher,
	mapper mappers.Mapper,
	randomHelper random.RandomHelper) (CustomerOutboxPublisher, error) {
	return &customerOutboxPublisher{
		logger:          logger,
		appConfig:       appConfig,
		contextHelper:   contextHelper,
		outboxPublisher: outboxPublisher,
		mapper:          mapper,
		randomHelper:    randomHelper,
	}, nil
}

func (s *customerOutboxPublisher) Publish(
	ctx context.Context,
	tx *entities.Tx,
	customers []*models.Customer) error {
	time := timestamppb.New(time.Now().UTC())
	correlationID := s.contextHelper.GetCorrelationId(ctx)

	result := slices.Map(customers, func(item *models.Customer) error {
		id, err := s.randomHelper.Generate()
		if err != nil {
			return err
		}

		event := customereventv1.Event{
			Metadata: &customereventv1.Metadata{
				Id:            id,
				DomainSource:  s.appConfig.DomainSource,
				AppSource:     s.appConfig.AppSource,
				Time:          time,
				CorrelationId: correlationID,
			},
			Data: &customereventv1.Data{},
		}

		if item.DeletedAt != nil && !item.DeletedAt.IsZero() {
			event.Metadata.Type = customereventv1.Type_CustomerDeleted
		} else {
			event.Metadata.Type = customereventv1.Type_CustomerUpserted
		}

		event.Data.AfterState = s.mapper.CustomerToEventCustomer(item)

		return s.outboxPublisher.Publish(
			ctx,
			tx,
			customereventv1.TopicName,
			[]byte(item.ID),
			make(map[string][]byte),
			&event)
	})

	return errors.ReduceErrors(result)
}
