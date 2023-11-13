package services

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	customerv1 "github.com/sigmasee/sigmasee/shared/clients/events/sigmasee/customer/v1"
	kafkagoenterprise "github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	"go.uber.org/zap"
)

type TopicService interface {
	CreateTopics(ctx context.Context) error
}

type topicService struct {
	logger      *zap.SugaredLogger
	kafkaConfig kafkagoenterprise.KafkaGoKafkaConfig
}

func NewTopicService(
	logger *zap.SugaredLogger,
	kafkaConfig kafkagoenterprise.KafkaGoKafkaConfig) (TopicService, error) {
	return &topicService{
		logger:      logger,
		kafkaConfig: kafkaConfig,
	}, nil
}

func (s *topicService) CreateTopics(ctx context.Context) error {
	if err := s.createTopic(ctx, customerv1.TopicName); err != nil {
		return err
	}

	for i := 0; i < customerv1.RetryTopicNameCount; i++ {
		if err := s.createTopic(ctx, fmt.Sprintf("%s.%d", customerv1.RetryTopicNamePrefix, i)); err != nil {
			return err
		}
	}

	if err := s.createTopic(ctx, customerv1.DeadLetterTopicName); err != nil {
		return err
	}

	return nil
}

func (s *topicService) createTopic(ctx context.Context, topicName string) error {
	dialer, err := kafka.DialLeader(ctx, "tcp", s.kafkaConfig.BootstrapServers, topicName, 0)
	if err != nil {
		return fmt.Errorf("failed to create topic %s. Error: %v", topicName, err)
	}

	defer dialer.Close()

	return nil
}
