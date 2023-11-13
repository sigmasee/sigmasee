// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"github.com/sigmasee/sigmasee/customer/customer-processor/http"
	mappers2 "github.com/sigmasee/sigmasee/customer/customer-processor/mappers"
	"github.com/sigmasee/sigmasee/customer/customer-processor/openapi"
	"github.com/sigmasee/sigmasee/customer/customer-processor/subscribers"
	"github.com/sigmasee/sigmasee/customer/shared/mappers"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	"github.com/sigmasee/sigmasee/shared/clients/events/sigmasee/customer/v1"
	"github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/context"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func NewCustomerConsumer(logger *zap.SugaredLogger, appConfig configuration.AppConfig, kafkaConfig kafka.KafkaGoKafkaConfig, entgoClient repositories.EntgoClient) (v1.Consumer, error) {
	contextHelper, err := context.NewContextHelper()
	if err != nil {
		return nil, err
	}
	mapper, err := mappers.NewMapper()
	if err != nil {
		return nil, err
	}
	identityRepository, err := repositories.NewIdentityRepository(entgoClient, mapper)
	if err != nil {
		return nil, err
	}
	customerRepository, err := repositories.NewCustomerRepository(entgoClient, mapper, identityRepository)
	if err != nil {
		return nil, err
	}
	customerSettingsRepository, err := repositories.NewCustomerSettingsRepository(entgoClient, mapper)
	if err != nil {
		return nil, err
	}
	mappersMapper, err := mappers2.NewMapper()
	if err != nil {
		return nil, err
	}
	randomHelper, err := appsetup.NewRandomHelper()
	if err != nil {
		return nil, err
	}
	subscriber, err := subscribers.NewCustomerSubscriber(logger, appConfig, customerRepository, customerSettingsRepository, mappersMapper, randomHelper)
	if err != nil {
		return nil, err
	}
	kafkaClient, err := kafka.NewKafkaGoKafkaClient(logger, kafkaConfig)
	if err != nil {
		return nil, err
	}
	messageConsumer, err := kafka.NewKafkaGoKafkaMessageConsumer(logger, appConfig, kafkaClient)
	if err != nil {
		return nil, err
	}
	messageProducer, err := kafka.NewKafkaGoKafkaMessageProducer(logger, kafkaClient)
	if err != nil {
		return nil, err
	}
	consumer := v1.NewConsumer(logger, appConfig, contextHelper, subscriber, messageConsumer, messageProducer)
	return consumer, nil
}

func NewCustomerAwsLambdaConsumer(logger *zap.SugaredLogger, appConfig configuration.AppConfig, awsLambdaConfig configuration.AwsLambdaConfig, kafkaConfig kafka.KafkaGoKafkaConfig, entgoClient repositories.EntgoClient) (v1.AwsLambdaConsumer, error) {
	contextHelper, err := context.NewContextHelper()
	if err != nil {
		return nil, err
	}
	mapper, err := mappers.NewMapper()
	if err != nil {
		return nil, err
	}
	identityRepository, err := repositories.NewIdentityRepository(entgoClient, mapper)
	if err != nil {
		return nil, err
	}
	customerRepository, err := repositories.NewCustomerRepository(entgoClient, mapper, identityRepository)
	if err != nil {
		return nil, err
	}
	customerSettingsRepository, err := repositories.NewCustomerSettingsRepository(entgoClient, mapper)
	if err != nil {
		return nil, err
	}
	mappersMapper, err := mappers2.NewMapper()
	if err != nil {
		return nil, err
	}
	randomHelper, err := appsetup.NewRandomHelper()
	if err != nil {
		return nil, err
	}
	subscriber, err := subscribers.NewCustomerSubscriber(logger, appConfig, customerRepository, customerSettingsRepository, mappersMapper, randomHelper)
	if err != nil {
		return nil, err
	}
	kafkaClient, err := kafka.NewKafkaGoKafkaClient(logger, kafkaConfig)
	if err != nil {
		return nil, err
	}
	messageProducer, err := kafka.NewKafkaGoKafkaMessageProducer(logger, kafkaClient)
	if err != nil {
		return nil, err
	}
	awsLambdaConsumer := v1.NewAwsLambdaConsumer(logger, appConfig, contextHelper, awsLambdaConfig, subscriber, messageProducer)
	return awsLambdaConsumer, nil
}

func NewHttpServer(logger *zap.SugaredLogger, appConfig configuration.AppConfig) (http.HttpServer, error) {
	openApiCustomerProcessorV1, err := openapi.NewOpenApiCustomerProcessorV1(logger)
	if err != nil {
		return nil, err
	}
	httpServer, err := http.NewHttpServer(logger, appConfig, openApiCustomerProcessorV1)
	if err != nil {
		return nil, err
	}
	return httpServer, nil
}