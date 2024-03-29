// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"github.com/sigmasee/sigmasee/customer/customerctl/services"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewTopicService(logger *zap.SugaredLogger, kafkaConfig kafka.KafkaGoKafkaConfig) (services.TopicService, error) {
	topicService, err := services.NewTopicService(logger, kafkaConfig)
	if err != nil {
		return nil, err
	}
	return topicService, nil
}
