package graphqlv1

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"net/url"

	"github.com/life4/genesis/slices"
	customerv1schema "github.com/sigmasee/sigmasee/api-definitions/graphql/sigmasee/customer/v1"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/configuration"
	"github.com/wundergraph/graphql-go-tools/pkg/engine/datasource/graphql_datasource"
)

type DatasourcePoller interface {
}

type datasourcePoller struct {
}

type serviceConfig struct {
	Name   string
	URL    string
	WS     string
	Schema string
}

var customerV1Schema string

func init() {
	customerV1Schema = fmt.Sprintf("%s\n\n%s", customerv1schema.Schema, customerv1schema.EntgoSchema)
}

func NewDatasourcePoller(
	ctx context.Context,
	apiEndpoints configuration.ApiEndpoints,
	gateway Gateway) (DatasourcePoller, error) {

	customerUrl, err := url.JoinPath(apiEndpoints.Customer, "customer/api/v1/graphql")
	if err != nil {
		return nil, err
	}

	services := []serviceConfig{
		{Name: "customer", URL: customerUrl, Schema: customerV1Schema},
	}

	sdlMap := slices.Reduce(services, make(map[string]string), func(serviceConf serviceConfig, acc map[string]string) map[string]string {
		acc[serviceConf.Name] = serviceConf.Schema

		return acc
	})

	gateway.UpdateDataSources(ctx, createDatasourceConfig(services, sdlMap))

	return &datasourcePoller{}, nil
}

func createDatasourceConfig(services []serviceConfig, sdlMap map[string]string) []graphql_datasource.Configuration {
	dataSourceConfigs := make([]graphql_datasource.Configuration, 0, len(services))

	for _, serviceConfig := range services {
		sdl, exists := sdlMap[serviceConfig.Name]
		if !exists {
			continue
		}

		dataSourceConfig := graphql_datasource.Configuration{
			Fetch: graphql_datasource.FetchConfiguration{
				URL:    serviceConfig.URL,
				Method: http.MethodPost,
				Header: http.Header{
					"Authorization":    []string{"{{ .request.headers.Authorization }}"},
					"X-Correlation-Id": []string{"{{ .request.headers.X-Correlation-Id }}"},
				},
			},
			Subscription: graphql_datasource.SubscriptionConfiguration{
				URL: serviceConfig.WS,
			},
			Federation: graphql_datasource.FederationConfiguration{
				Enabled:    true,
				ServiceSDL: sdl,
			},
		}

		dataSourceConfigs = append(dataSourceConfigs, dataSourceConfig)
	}

	return dataSourceConfigs
}
