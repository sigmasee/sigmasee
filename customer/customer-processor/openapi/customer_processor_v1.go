package openapi

import (
	"net/http"

	openapiv1 "github.com/sigmasee/sigmasee/customer/customer-processor/openapi/v1"
	"go.uber.org/zap"
)

type OpenApiCustomerProcessorV1 interface {
	GetHttpHandler() http.Handler
}

type openApiCustomerProcessorV1 struct {
	logger *zap.SugaredLogger
}

func NewOpenApiCustomerProcessorV1(
	logger *zap.SugaredLogger) (OpenApiCustomerProcessorV1, error) {
	return &openApiCustomerProcessorV1{
		logger: logger,
	}, nil
}

func (s *openApiCustomerProcessorV1) GetHttpHandler() http.Handler {
	return openapiv1.Handler(s)
}

func (oaiv1 *openApiCustomerProcessorV1) Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (oaiv1 *openApiCustomerProcessorV1) Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
