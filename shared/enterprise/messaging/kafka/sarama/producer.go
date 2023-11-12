package kafka

import (
	"context"
	"time"

	"github.com/IBM/sarama"
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type kafkaMessageProducer struct {
	logger      *zap.SugaredLogger
	kafkaClient KafkaClient
}

func NewSaramaKafkaMessageProducer(
	logger *zap.SugaredLogger,
	kafkaClient KafkaClient) (messaging.MessageProducer, error) {
	return &kafkaMessageProducer{
		logger:      logger,
		kafkaClient: kafkaClient,
	}, nil
}

func (s *kafkaMessageProducer) Produce(ctx context.Context, messages []messaging.Message) error {
	start := time.Now()
	defer func(start time.Time) {
		s.logger.Infof("Produce - Execution time: %s", time.Since(start))
	}(start)

	brokers, config, err := s.kafkaClient.NewClientConfig()
	if err != nil {
		return err
	}

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return err
	}

	defer func() {
		if err := syncProducer.Close(); err != nil {
			s.logger.Errorf("Failed to call Close function for the syncProducer. Error: %v", err)
		}
	}()

	producerMessages := slices.Map(messages, func(message messaging.Message) *sarama.ProducerMessage {
		return &sarama.ProducerMessage{
			Topic:    message.Topic,
			Key:      sarama.ByteEncoder(message.Key),
			Metadata: message.Headers,
			Value:    sarama.ByteEncoder(message.Payload),
		}
	})

	return syncProducer.SendMessages(producerMessages)
}
