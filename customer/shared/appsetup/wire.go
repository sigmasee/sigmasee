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
	"context"

	"github.com/go-co-op/gocron"
	"github.com/google/wire"
	"github.com/sigmasee/sigmasee/customer/shared/outbox"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	configurationenterprise "github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"github.com/sigmasee/sigmasee/shared/enterprise/database/postgres"
	kafka "github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	enterpriseOutbox "github.com/sigmasee/sigmasee/shared/enterprise/outbox"
	"go.uber.org/zap"
)

func NewEntgoClient(
	logger *zap.SugaredLogger,
	databaseConfig database.DatabaseConfig,
	postgresConfig postgres.PostgresConfig,
	appConfig configurationenterprise.AppConfig) (repositories.EntgoClient, error) {
	wire.Build(
		postgres.NewPostgres,
		repositories.NewEntgoClient)

	return nil, nil
}

func NewOutboxBackgroundService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	kafkaClient kafka.KafkaClient,
	outboxConfig enterpriseOutbox.OutboxConfig,
	entgoClient repositories.EntgoClient,
	jobScheduler *gocron.Scheduler) (outbox.OutboxBackgroundService, error) {
	wire.Build(
		kafka.NewKafkaGoKafkaMessageProducer,
		outbox.NewOutboxBackgroundService)

	return nil, nil
}
