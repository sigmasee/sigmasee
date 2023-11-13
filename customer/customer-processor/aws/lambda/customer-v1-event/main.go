package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sigmasee/sigmasee/customer/customer-processor/appsetup"
	"github.com/sigmasee/sigmasee/customer/customer-processor/configuration"
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

	consumer, err := appsetup.NewCustomerAwsLambdaConsumer(
		sugarLogger,
		config.App,
		config.AwsLambda,
		config.Kafka,
		entgoClient)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	lambda.StartWithOptions(consumer.Handle, lambda.WithContext(ctx))
}
