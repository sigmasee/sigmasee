package outbox

import (
	"context"
	"time"

	"github.com/sigmasee/sigmasee/customer/shared/entities"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customeroutbox"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging/schemaregistry"
	"github.com/sigmasee/sigmasee/shared/enterprise/random"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type OutboxPublisher interface {
	Publish(
		ctx context.Context,
		tx *entities.Tx,
		topic string,
		key []byte,
		headers map[string][]byte,
		event proto.Message) error
}

type outboxPublisher struct {
	logger       *zap.SugaredLogger
	entgoClient  repositories.EntgoClient
	randomHelper random.RandomHelper
}

func NewOutboxPublisher(
	logger *zap.SugaredLogger,
	entgoClient repositories.EntgoClient,
	randomHelper random.RandomHelper) (OutboxPublisher, error) {
	return &outboxPublisher{
		logger:       logger,
		entgoClient:  entgoClient,
		randomHelper: randomHelper,
	}, nil
}

func (s *outboxPublisher) Publish(
	ctx context.Context,
	tx *entities.Tx,
	topic string,
	key []byte,
	headers map[string][]byte,
	event proto.Message) error {
	payload, err := proto.Marshal(event)
	if err != nil {
		s.logger.Errorf("Failed to serialize event to Protobuf byte array. Error: %v", err)

		return err
	}

	id, err := s.randomHelper.Generate()
	if err != nil {
		return err
	}

	if _, err := s.entgoClient.GetCustomerOutboxClient(tx).
		Create().
		SetID(id).
		SetTimestamp(time.Now()).
		SetTopic(topic).
		SetKey(key).
		SetPayload(append(schemaregistry.ReservedHeaders, payload...)).
		SetHeaders(headers).
		SetRetryCount(0).
		SetStatus(customeroutbox.StatusPENDING).
		SetNillableLastRetry(nil).
		Save(ctx); err != nil {
		s.logger.Errorf("Failed to save outbox item. Error: %v", err)

		return err
	}

	return nil
}
