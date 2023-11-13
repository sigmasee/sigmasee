// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package appsetup

import (
	"github.com/google/wire"
	"github.com/sigmasee/sigmasee/customer/customer-processor/http"
	"github.com/sigmasee/sigmasee/customer/customer-processor/mappers"
	openapi "github.com/sigmasee/sigmasee/customer/customer-processor/openapi"
	"github.com/sigmasee/sigmasee/customer/customer-processor/subscribers"
	sharedmappers "github.com/sigmasee/sigmasee/customer/shared/mappers"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	customerv1 "github.com/sigmasee/sigmasee/shared/clients/events/sigmasee/customer/v1"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/context"
	kafkagokafka "github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	"go.uber.org/zap"
)

func NewCustomerConsumer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	kafkaConfig kafkagokafka.KafkaGoKafkaConfig,
	entgoClient repositories.EntgoClient) (customerv1.Consumer, error) {
	wire.Build(
		kafkagokafka.NewKafkaGoKafkaMessageConsumer,
		customerv1.NewConsumer,
		subscribers.NewCustomerSubscriber,
		context.NewContextHelper,
		kafkagokafka.NewKafkaGoKafkaMessageProducer,
		kafkagokafka.NewKafkaGoKafkaClient,
		repositories.NewCustomerRepository,
		repositories.NewCustomerSettingsRepository,
		repositories.NewIdentityRepository,
		enterpriseappsetup.NewRandomHelper,
		mappers.NewMapper,
		sharedmappers.NewMapper)

	return nil, nil
}

func NewCustomerAwsLambdaConsumer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	awsLambdaConfig configuration.AwsLambdaConfig,
	kafkaConfig kafkagokafka.KafkaGoKafkaConfig,
	entgoClient repositories.EntgoClient) (customerv1.AwsLambdaConsumer, error) {
	wire.Build(
		customerv1.NewAwsLambdaConsumer,
		subscribers.NewCustomerSubscriber,
		context.NewContextHelper,
		kafkagokafka.NewKafkaGoKafkaMessageProducer,
		kafkagokafka.NewKafkaGoKafkaClient,
		repositories.NewCustomerRepository,
		repositories.NewCustomerSettingsRepository,
		repositories.NewIdentityRepository,
		enterpriseappsetup.NewRandomHelper,
		mappers.NewMapper,
		sharedmappers.NewMapper)

	return nil, nil
}

func NewHttpServer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig) (http.HttpServer, error) {
	wire.Build(
		http.NewHttpServer,
		openapi.NewOpenApiCustomerProcessorV1)

	return nil, nil
}
