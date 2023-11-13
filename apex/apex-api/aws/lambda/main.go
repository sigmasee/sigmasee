package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-co-op/gocron"
	"github.com/sigmasee/sigmasee/apex/apex-api/appsetup"
	"github.com/sigmasee/sigmasee/apex/apex-api/configuration"
	apexappsetup "github.com/sigmasee/sigmasee/apex/shared/appsetup"
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

	entgoClient, err := apexappsetup.NewEntgoClient(
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

	httpServer, err := appsetup.NewHttpServer(
		sugarLogger,
		config.App,
		entgoClient,
		tokenService)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	lambda.StartWithOptions(httpadapter.NewV2(httpServer.GetHandler()).ProxyWithContext, lambda.WithContext(ctx))
}
