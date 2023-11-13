package services

import (
	"context"
	"time"

	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/customer/shared/entities"
	outboxentity "github.com/sigmasee/sigmasee/customer/shared/entities/customeroutbox"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging"
	"github.com/sigmasee/sigmasee/shared/enterprise/outbox"
	"go.uber.org/zap"
)

type OutboxService interface {
	HandleOutboxChangeFeedRequest(ctx context.Context) error
}

type outboxService struct {
	logger          *zap.SugaredLogger
	outboxConfig    outbox.OutboxConfig
	entgoClient     repositories.EntgoClient
	messageProducer messaging.MessageProducer
}

func NewOutboxService(
	logger *zap.SugaredLogger,
	outboxConfig outbox.OutboxConfig,
	entgoClient repositories.EntgoClient,
	messageProducer messaging.MessageProducer) (OutboxService, error) {
	return &outboxService{
		logger:          logger,
		outboxConfig:    outboxConfig,
		entgoClient:     entgoClient,
		messageProducer: messageProducer,
	}, nil
}

func (s *outboxService) HandleOutboxChangeFeedRequest(ctx context.Context) error {
	client := s.entgoClient.GetCustomerOutboxClient(nil)
	records, err := client.Query().
		Where(outboxentity.StatusEQ(outboxentity.StatusPENDING)).
		All(ctx)
	if err != nil {
		s.logger.Errorf("Failed to fetch outbox items. Error: %v", err)

		return nil
	}

	numOfRecords := len(records)
	if numOfRecords == 0 {
		return nil
	}

	start := time.Now()
	defer func(start time.Time) {
		s.logger.Infof("Run - Message count %d - Execution time: %s", numOfRecords, time.Since(start))
	}(start)

	messages := slices.Map(records, func(record *entities.CustomerOutbox) messaging.Message {
		return messaging.Message{
			Topic:   record.Topic,
			Key:     record.Key,
			Headers: record.Headers,
			Payload: record.Payload,
		}
	})

	if err = s.messageProducer.Produce(ctx, messages); err != nil {
		s.logger.Errorf("Failed to send messages. Error: %v", err)

		now := time.Now().UTC()
		for _, record := range records {
			if record.RetryCount == s.outboxConfig.MaxRetryCount {
				if err = client.
					UpdateOne(record).
					SetStatus(outboxentity.StatusFAILED).
					SetLastRetry(now).
					SetProcessingErrors(append(record.ProcessingErrors, err.Error())).
					Exec(ctx); err != nil {
					s.logger.Errorf("Failed to update failed outbox item for topic %s. Error: %v", record.Topic, err)
				}
			} else {
				if err = client.
					UpdateOne(record).
					AddRetryCount(1).
					SetLastRetry(now).
					SetProcessingErrors(append(record.ProcessingErrors, err.Error())).
					Exec(ctx); err != nil {
					s.logger.Errorf("Failed to update failed outbox item for topic %s. Error: %v", record.Topic, err)
				}
			}
		}

		return err
	} else {
		recordIDs := slices.Map(records, func(item *entities.CustomerOutbox) string {
			return item.ID
		})

		if len(recordIDs) > 0 {
			if _, err = client.
				Delete().
				Where(outboxentity.IDIn(recordIDs...)).
				Exec(ctx); err != nil {
				s.logger.Errorf("Failed to delete succeeded outbox item. Error: %v", err)

				return err
			}
		}
	}

	return nil
}
