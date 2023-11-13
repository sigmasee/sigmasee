package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-co-op/gocron"
	"github.com/sigmasee/sigmasee/customer/customer-api/appsetup"
	"github.com/sigmasee/sigmasee/customer/customer-api/configuration"
	customerappsetup "github.com/sigmasee/sigmasee/customer/shared/appsetup"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	_, sugarLogger := logger.CreateProductionLogger()
	defer func() {
		_ = sugarLogger.Sync()
	}()

	configurationHelper, err := enterpriseappsetup.NewConfigurationHelper(sugarLogger)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	var config configuration.Config
	if err := configurationHelper.LoadYaml("config.yaml", &config); err != nil {
		sugarLogger.Fatal(err)
	}

	entgoClient, err := customerappsetup.NewEntgoClient(
		sugarLogger,
		config.Database,
		config.Postgres,
		config.App)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	jobScheduler := gocron.NewScheduler(time.UTC)
	jobScheduler.StartAsync()
	defer jobScheduler.Stop()

	tokenService, err := enterpriseappsetup.NewTokenService(
		ctx,
		sugarLogger,
		config.CognitoIdentityProvider,
		config.GoogleIdentityProvider,
		config.SlackIdentityProvider,
		jobScheduler)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	kafkaClient, err := enterpriseappsetup.NewKafkaGoKafkaClient(
		sugarLogger,
		config.Kafka)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	if _, err = customerappsetup.NewOutboxBackgroundService(
		ctx,
		sugarLogger,
		kafkaClient,
		config.Outbox,
		entgoClient,
		jobScheduler); err != nil {
		sugarLogger.Fatal(err)
	}

	httpServer, err := appsetup.NewHttpServer(
		sugarLogger,
		config.App,
		config.Outbox,
		config.Intercom,
		entgoClient,
		tokenService,
		kafkaClient)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	lambda.StartWithOptions(httpadapter.NewV2(httpServer.GetHandler()).ProxyWithContext, lambda.WithContext(ctx))
}
