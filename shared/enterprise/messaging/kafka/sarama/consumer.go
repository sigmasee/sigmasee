package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging"
	"go.uber.org/zap"
)

type kafkaMessageConsumer struct {
	logger      *zap.SugaredLogger
	appConfig   configuration.AppConfig
	kafkaClient KafkaClient
}

type consumerGroupHandler struct {
	logger   *zap.SugaredLogger
	callback messaging.MessageCallback
}

func NewSaramaKafkaMessageConsumer(
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
	brokers, config, err := s.kafkaClient.NewClientConfig()
	if err != nil {
		return err
	}

	consumerGroup, err := sarama.NewConsumerGroup(brokers, s.appConfig.GetSource(), config)
	if err != nil {
		return err
	}

	defer func() {
		if err := consumerGroup.Close(); err != nil {
			s.logger.Errorf("Failed to close consumer group. Error: %v", err)
		}
	}()

	for {
		if err := consumerGroup.Consume(ctx, []string{topic}, &consumerGroupHandler{
			logger:   s.logger,
			callback: callback,
		}); err != nil {
			s.logger.Errorf("Faild to call Consume on topic: %s. Error: %v", topic, err)

			continue
		}

		err = ctx.Err()

		if err == context.Canceled {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func (s *consumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (s *consumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (s *consumerGroupHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		s.logger.Infof(
			"Received message (topic: %s, partition: %d, offset: %d)",
			message.Topic,
			message.Partition,
			message.Offset)

		headers := make(map[string][]byte)
		for _, header := range message.Headers {
			headers[string(header.Key)] = header.Value
		}

		if err := s.callback(messaging.Message{
			Topic:     message.Topic,
			Key:       message.Key,
			Payload:   message.Value,
			Timestamp: &message.Timestamp,
			Headers:   headers,
		}); err != nil {
			return err
		}

		session.MarkMessage(message, "")
	}

	return nil
}
