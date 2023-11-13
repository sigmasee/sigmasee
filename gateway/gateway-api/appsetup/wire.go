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
	nethttp "net/http"

	"github.com/google/wire"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/configuration"
	graphqlv1 "github.com/sigmasee/sigmasee/gateway/gateway-api/graphql/v1"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/http"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/openapi"
	enterpriseconfiguration "github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"go.uber.org/zap"
)

func NewHttpServer(
	ctx context.Context,
	logger *zap.SugaredLogger,
	zapLogger *zap.Logger,
	appConfig enterpriseconfiguration.AppConfig,
	graphQLEndpoints configuration.ApiEndpoints,
	httpClient *nethttp.Client) (http.HttpServer, error) {
	wire.Build(
		http.NewHttpServer,
		graphqlv1.NewGraphQLServer,
		graphqlv1.NewGateway,
		graphqlv1.NewDatasourcePoller,
		openapi.NewOpenApiGatewayV1)

	return nil, nil
}
