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
	graphqlv1 "github.com/sigmasee/sigmasee/customer/customer-api/graphql/v1"
	"github.com/sigmasee/sigmasee/customer/customer-api/http"
	"github.com/sigmasee/sigmasee/customer/customer-api/mappers"
	openapi "github.com/sigmasee/sigmasee/customer/customer-api/openapi"
	"github.com/sigmasee/sigmasee/customer/customer-api/services"
	sharedmappers "github.com/sigmasee/sigmasee/customer/shared/mappers"
	"github.com/sigmasee/sigmasee/customer/shared/outbox"
	"github.com/sigmasee/sigmasee/customer/shared/publishers"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	configurationenterprise "github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/context"
	kafka "github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	enterpriseoutbox "github.com/sigmasee/sigmasee/shared/enterprise/outbox"
	"github.com/sigmasee/sigmasee/shared/enterprise/security/token"
	"go.uber.org/zap"
)

func NewHttpServer(
	logger *zap.SugaredLogger,
	appConfig configurationenterprise.AppConfig,
	outboxConfig enterpriseoutbox.OutboxConfig,
	intercomConfig configurationenterprise.IntercomConfig,
	entgoClient repositories.EntgoClient,
	tokenService token.TokenService,
	kafkaClient kafka.KafkaClient) (http.HttpServer, error) {
	wire.Build(
		http.NewHttpServer,
		graphqlv1.NewGraphQLServer,
		openapi.NewOpenApiCustomerV1,
		context.NewContextHelper,
		services.NewCustomerService,
		repositories.NewCustomerRepository,
		repositories.NewCustomerSettingsRepository,
		repositories.NewIdentityRepository,
		outbox.NewOutboxPublisher,
		publishers.NewCustomerOutboxPublisher,
		services.NewOutboxService,
		kafka.NewKafkaGoKafkaMessageProducer,
		enterpriseappsetup.NewRandomHelper,
		mappers.NewMapper,
		sharedmappers.NewMapper)

	return nil, nil
}
