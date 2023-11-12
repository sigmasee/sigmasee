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
	"github.com/sigmasee/sigmasee/shared/enterprise/aws/ses"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	configurationenterprise "github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	enterprisecontext "github.com/sigmasee/sigmasee/shared/enterprise/context"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"github.com/sigmasee/sigmasee/shared/enterprise/database/postgres"
	kafkagokafka "github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	saramakafka "github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/sarama"
	"github.com/sigmasee/sigmasee/shared/enterprise/os"
	"github.com/sigmasee/sigmasee/shared/enterprise/random"
	"github.com/sigmasee/sigmasee/shared/enterprise/security/token"
	"github.com/sigmasee/sigmasee/shared/enterprise/time"
	"go.uber.org/zap"
)

func NewTimeHelper(logger *zap.SugaredLogger) (time.TimeHelper, error) {
	wire.Build(
		os.NewOsHelper,
		time.NewTimeHelper)

	return nil, nil
}

func NewConfigurationHelper(logger *zap.SugaredLogger) (configuration.ConfigurationHelper, error) {
	wire.Build(
		os.NewOsHelper,
		configuration.NewConfigurationHelper)

	return nil, nil
}

func NewAwsSesHelper(logger *zap.SugaredLogger) (ses.AwsSesHelper, error) {
	wire.Build(
		os.NewOsHelper,
		ses.NewAwsSesHelper)

	return nil, nil
}

func NewTokenService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	cognitoConfig token.CognitoConfig,
	googleConfig token.GoogleConfig,
	slackConfig token.SlackConfig,
	jobScheduler *gocron.Scheduler) (token.TokenService, error) {
	wire.Build(
		token.NewCognitoTokenService,
		token.NewGoogleTokenService,
		token.NewSlackTokenService,
		token.NewTokenService)

	return nil, nil
}

func NewContextHelper() (enterprisecontext.ContextHelper, error) {
	wire.Build(
		enterprisecontext.NewContextHelper)

	return nil, nil
}

func NewDatabase(
	logger *zap.SugaredLogger,
	postgresConfig postgres.PostgresConfig,
	appConfig configurationenterprise.AppConfig) (database.Database, error) {
	wire.Build(postgres.NewPostgres)

	return nil, nil
}

func NewOsHelper() (os.OsHelper, error) {
	wire.Build(os.NewOsHelper)

	return nil, nil
}

func NewSaramaKafkaClient(
	logger *zap.SugaredLogger,
	kafkaConfig saramakafka.SaramaKafkaConfig) (saramakafka.KafkaClient, error) {
	wire.Build(
		saramakafka.NewSaramaKafkaClient)

	return nil, nil
}

func NewKafkaGoKafkaClient(
	logger *zap.SugaredLogger,
	kafkaConfig kafkagokafka.KafkaGoKafkaConfig) (kafkagokafka.KafkaClient, error) {
	wire.Build(
		kafkagokafka.NewKafkaGoKafkaClient)

	return nil, nil
}

func NewRandomHelper() (random.RandomHelper, error) {
	wire.Build(random.NewRandomHelper)

	return nil, nil
}
