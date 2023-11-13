package publishers

import (
	"context"
	"time"

	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/customer/shared/mappers"
	"github.com/sigmasee/sigmasee/customer/shared/models"
	customereventv1 "github.com/sigmasee/sigmasee/shared/clients/events/sigmasee/customer/v1"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	enterprisecontext "github.com/sigmasee/sigmasee/shared/enterprise/context"
	"github.com/sigmasee/sigmasee/shared/enterprise/random"
	"github.com/sigmasee/sigmasee/shared/enterprise/tuples"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerPublisher interface {
	Publish(
		ctx context.Context,
		customers []*models.Customer) error
}

type customerPublisher struct {
	logger        *zap.SugaredLogger
	appConfig     configuration.AppConfig
	contextHelper enterprisecontext.ContextHelper
	producer      customereventv1.Producer
	mapper        mappers.Mapper
	randomHelper  random.RandomHelper
}

func NewCustomerPublisher(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	contextHelper enterprisecontext.ContextHelper,
	producer customereventv1.Producer,
	mapper mappers.Mapper,
	randomHelper random.RandomHelper) (CustomerPublisher, error) {
	return &customerPublisher{
		logger:        logger,
		appConfig:     appConfig,
		contextHelper: contextHelper,
		producer:      producer,
		mapper:        mapper,
		randomHelper:  randomHelper,
	}, nil
}

func (s *customerPublisher) Publish(
	ctx context.Context,
	customers []*models.Customer) error {
	time := timestamppb.New(time.Now().UTC())
	correlationID := s.contextHelper.GetCorrelationId(ctx)

	result := slices.Map(customers, func(item *models.Customer) tuples.ValueErrorTuple[*customereventv1.Message] {
		id, err := s.randomHelper.Generate()
		if err != nil {
			return tuples.ValueErrorTuple[*customereventv1.Message]{
				Value: nil,
				Error: err,
			}
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

		return tuples.ValueErrorTuple[*customereventv1.Message]{
			Value: &customereventv1.Message{
				Key:     []byte(item.ID),
				Headers: make(map[string][]byte),
				Event:   &event,
			},
			Error: nil,
		}
	})

	if err := tuples.ReduceErrors(result); err != nil {
		return err
	}

	return s.producer.Produce(
		ctx,
		tuples.GetValues(result))
}
