package kafka

import (
	"context"
	"time"

	"github.com/life4/genesis/slices"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type kafkaMessageProducer struct {
	logger      *zap.SugaredLogger
	kafkaClient KafkaClient
}

func NewKafkaGoKafkaMessageProducer(
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

	brokers, dialer, err := s.kafkaClient.NewDialer()
	if err != nil {
		return err
	}

	writer := kafkago.NewWriter(kafkago.WriterConfig{
		Brokers: brokers,
		Dialer:  dialer,
	})

	defer func() {
		if err := writer.Close(); err != nil {
			s.logger.Errorf("Failed to call Close function for the writer. Error: %v", err)
		}
	}()

	messagesToWrite := slices.Map(messages, func(message messaging.Message) kafkago.Message {
		headers := []kafkago.Header{}
		for key, value := range message.Headers {
			headers = append(headers, kafkago.Header{Key: key, Value: value})
		}

		return kafkago.Message{
			Topic:   message.Topic,
			Key:     message.Key,
			Headers: headers,
			Value:   message.Payload,
		}
	})

	return writer.WriteMessages(ctx, messagesToWrite...)
}
