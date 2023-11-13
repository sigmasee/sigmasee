package openapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/gateway/gateway-api/configuration"
	openapiv1 "github.com/sigmasee/sigmasee/gateway/gateway-api/openapi/v1"
	customeropenapiv1 "github.com/sigmasee/sigmasee/shared/clients/openapi/sigmasee/customer/v1"
	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
	"go.uber.org/zap"
)

type OpenApiGatewayV1 interface {
	GetHttpHandler() http.Handler
}

type apiOutboxChangeFeedTriggerFunc func(context.Context) error

type openApiGatewayV1 struct {
	logger                          *zap.SugaredLogger
	apiEndpoints                    configuration.ApiEndpoints
	apiOutboxChangeFeedTriggersFunc []apiOutboxChangeFeedTriggerFunc
}

func NewOpenApiGatewayV1(
	logger *zap.SugaredLogger,
	apiEndpoints configuration.ApiEndpoints) (OpenApiGatewayV1, error) {
	service := &openApiGatewayV1{
		logger:       logger,
		apiEndpoints: apiEndpoints,
	}

	service.apiOutboxChangeFeedTriggersFunc = []apiOutboxChangeFeedTriggerFunc{
		service.customerProcessOutboxChangeFeed,
	}

	return service, nil
}

func (s *openApiGatewayV1) GetHttpHandler() http.Handler {
	return openapiv1.Handler(s)
}

func (s *openApiGatewayV1) Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiGatewayV1) Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiGatewayV1) ProcessOutboxChangeFeed(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func(start time.Time) {
		s.logger.Infof("ProcessOutboxChangeFeed - Execution time: %s", time.Since(start))
	}(start)

	var err error

	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to excecute ProcessOutboxChangeFeed. Error: %v", err)
		}
	}()

	result := slices.MapAsync(s.apiOutboxChangeFeedTriggersFunc, 0, func(item apiOutboxChangeFeedTriggerFunc) error {
		return item(r.Context())
	})

	err = errors.ReduceErrors(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *openApiGatewayV1) customerProcessOutboxChangeFeed(ctx context.Context) (err error) {
	start := time.Now()
	defer func(start time.Time) {
		s.logger.Infof("customerProcessOutboxChangeFeed - Execution time: %s", time.Since(start))
	}(start)

	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to excecute customerProcessOutboxChangeFeed. Error: %v", err)
		}
	}()

	client, err := customeropenapiv1.NewClient(s.apiEndpoints.Customer)
	if err != nil {
		return
	}

	response, err := client.ProcessOutboxChangeFeed(ctx)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("client returned http status code: %d", response.StatusCode)

		return
	}

	return
}
