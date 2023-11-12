package kafka

import (
	"context"
	"io"

	"github.com/life4/genesis/slices"
	kafkago "github.com/segmentio/kafka-go"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type kafkaMessageConsumer struct {
	logger      *zap.SugaredLogger
	appConfig   configuration.AppConfig
	kafkaClient KafkaClient
}

func NewKafkaGoKafkaMessageConsumer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	kafkaClient KafkaClient) (messaging.MessageConsumer, error) {
	return &kafkaMessageConsumer{
		logger:      logger,
		appConfig:   appConfig,
		kafkaClient: kafkaClient,
	}, nil
}

func (s *kafkaMessageConsumer) Consume(ctx context.Context, topic string, callback messaging.MessageCallback) error {
	brokers, dialer, err := s.kafkaClient.NewDialer()
	if err != nil {
		return err
	}

	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: brokers,
		GroupID: s.appConfig.GetSource(),
		Topic:   topic,
		Dialer:  dialer,
	})

	defer reader.Close()

	for {
		message, err := reader.FetchMessage(ctx)
		if err == io.EOF {
			return nil
		}

		if err == context.Canceled {
			return nil
		}

		if err != nil {
			s.logger.Errorf("Faild to call ReadMessage on topic: %s. Error: %v", topic, err)

			continue
		}

		headers := slices.Reduce(message.Headers, make(map[string][]byte), func(header kafkago.Header, acc map[string][]byte) map[string][]byte {
			acc[string(header.Key)] = header.Value

			return acc
		})

		err = callback(
			messaging.Message{
				Topic:     message.Topic,
				Key:       message.Key,
				Payload:   message.Value,
				Timestamp: &message.Time,
				Headers:   headers,
			})

		if err == nil {
			err = reader.CommitMessages(ctx, message)
			if err != nil {
				// TODO: 20230711 - Move to retry topic, for now returning error instead of handling it
				return err
			}
		} else {
			// TODO: 20230711 - Move to retry topic, for now returning error instead of handling it
			return err
		}
	}
}
