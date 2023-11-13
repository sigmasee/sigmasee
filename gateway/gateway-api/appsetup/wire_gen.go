// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package appsetup

import (
	"context"
	configuration2 "github.com/sigmasee/sigmasee/gateway/gateway-api/configuration"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/graphql/v1"
	http2 "github.com/sigmasee/sigmasee/gateway/gateway-api/http"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/openapi"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"go.uber.org/zap"
	"net/http"
)

// Injectors from wire.go:

func NewHttpServer(ctx context.Context, logger *zap.SugaredLogger, zapLogger *zap.Logger, appConfig configuration.AppConfig, graphQLEndpoints configuration2.ApiEndpoints, httpClient *http.Client) (http2.HttpServer, error) {
	gateway, err := graphqlv1.NewGateway(logger, zapLogger, httpClient)
	if err != nil {
		return nil, err
	}
	datasourcePoller, err := graphqlv1.NewDatasourcePoller(ctx, graphQLEndpoints, gateway)
	if err != nil {
		return nil, err
	}
	graphQLServer, err := graphqlv1.NewGraphQLServer(datasourcePoller, gateway)
	if err != nil {
		return nil, err
	}
	openApiGatewayV1, err := openapi.NewOpenApiGatewayV1(logger, graphQLEndpoints)
	if err != nil {
		return nil, err
	}
	httpServer, err := http2.NewHttpServer(logger, appConfig, graphQLServer, openApiGatewayV1)
	if err != nil {
		return nil, err
	}
	return httpServer, nil
}
