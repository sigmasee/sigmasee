package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/appsetup"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/configuration"
	enterpriseappsetup "github.com/sigmasee/sigmasee/shared/enterprise/appsetup"
	"github.com/sigmasee/sigmasee/shared/enterprise/logger"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	logger, sugarLogger := logger.CreateProductionLogger()
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

	httpServer, err := appsetup.NewHttpServer(
		ctx,
		sugarLogger,
		logger,
		config.App,
		config.ApiEndpoints,
		&http.Client{
			Timeout: config.ApiEndpoints.Timeout,
		})
	if err != nil {
		sugarLogger.Fatal(err)
	}

	handler, err := httpServer.GetHandler(ctx)
	if err != nil {
		sugarLogger.Fatal(err)
	}

	lambda.StartWithOptions(httpadapter.NewV2(handler).ProxyWithContext, lambda.WithContext(ctx))
}
