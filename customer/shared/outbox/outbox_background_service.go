package outbox

import (
	"context"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/customer/shared/entities"
	outboxentity "github.com/sigmasee/sigmasee/customer/shared/entities/customeroutbox"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging"
	"github.com/sigmasee/sigmasee/shared/enterprise/outbox"
	"go.uber.org/zap"
)

type OutboxBackgroundService interface {
	Run()
	RunAsync()
}

type outboxBackgroundService struct {
	ctx             context.Context
	logger          *zap.SugaredLogger
	outboxConfig    outbox.OutboxConfig
	messageProducer messaging.MessageProducer
	entgoClient     repositories.EntgoClient
	globalMutex     sync.Mutex
}

func NewOutboxBackgroundService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	outboxConfig outbox.OutboxConfig,
	messageProducer messaging.MessageProducer,
	entgoClient repositories.EntgoClient,
	jobScheduler *gocron.Scheduler) (OutboxBackgroundService, error) {
	instance := &outboxBackgroundService{
		ctx:             ctx,
		logger:          logger,
		outboxConfig:    outboxConfig,
		messageProducer: messageProducer,
		entgoClient:     entgoClient,
		globalMutex:     sync.Mutex{},
	}

	if _, err := jobScheduler.Every(1).Seconds().Do(func() {
		instance.Run()
	}); err != nil {
		return nil, err
	}

	return instance, nil
}

func (s *outboxBackgroundService) RunAsync() {
	go s.Run()
}

func (s *outboxBackgroundService) Run() {
	s.globalMutex.Lock()
	defer s.globalMutex.Unlock()

	client := s.entgoClient.GetCustomerOutboxClient(nil)

	now := time.Now().UTC()
	records, err := client.Query().
		Where(
			outboxentity.And(
				outboxentity.StatusEQ(outboxentity.StatusPENDING),
				outboxentity.Or(
					outboxentity.LastRetryLTE(now.Add(-1*s.outboxConfig.RetryDelay)),
					outboxentity.LastRetryIsNil()),
			),
		).
		All(s.ctx)
	if err != nil {
		s.logger.Errorf("Failed to fetch outbox items. Error: %v", err)

		return
	}

	numOfRecords := len(records)
	if numOfRecords == 0 {
		return
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

	if err = s.messageProducer.Produce(s.ctx, messages); err != nil {
		s.logger.Errorf("Failed to send messages. Error: %v", err)

		now = time.Now().UTC()
		for _, record := range records {
			if record.RetryCount == s.outboxConfig.MaxRetryCount {
				if err = client.
					UpdateOne(record).
					SetStatus(outboxentity.StatusFAILED).
					SetLastRetry(now).
					SetProcessingErrors(append(record.ProcessingErrors, err.Error())).
					Exec(s.ctx); err != nil {
					s.logger.Errorf("Failed to update failed outbox item for topic %s. Error: %v", record.Topic, err)
				}
			} else {
				if err = client.
					UpdateOne(record).
					AddRetryCount(1).
					SetLastRetry(now).
					SetProcessingErrors(append(record.ProcessingErrors, err.Error())).
					Exec(s.ctx); err != nil {
					s.logger.Errorf("Failed to update failed outbox item for topic %s. Error: %v", record.Topic, err)
				}
			}
		}
	} else {
		recordIDs := slices.Map(records, func(item *entities.CustomerOutbox) string {
			return item.ID
		})

		if len(recordIDs) > 0 {
			if _, err = client.
				Delete().
				Where(outboxentity.IDIn(recordIDs...)).
				Exec(s.ctx); err != nil {
				s.logger.Errorf("Failed to delete succeeded outbox item. Error: %v", err)
			}
		}
	}
}
