package http

import (
	"net/http"

	"github.com/rs/cors"
	openapi "github.com/sigmasee/sigmasee/customer/customer-processor/openapi"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"go.uber.org/zap"
)

type HttpServer interface {
	GetHandler() http.Handler
	ListenAndServe() error
}

type httpServer struct {
	logger            *zap.SugaredLogger
	appConfig         configuration.AppConfig
	openApiCustomerV1 openapi.OpenApiCustomerProcessorV1
}

func NewHttpServer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	openApiCustomerV1 openapi.OpenApiCustomerProcessorV1) (HttpServer, error) {
	return &httpServer{
		appConfig:         appConfig,
		logger:            logger,
		openApiCustomerV1: openApiCustomerV1,
	}, nil
}

func (s *httpServer) GetHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", s.setResponseTypeToJSONMiddleware(s.openApiCustomerV1.GetHttpHandler()))

	return s.applyCors(mux)
}

func (s *httpServer) ListenAndServe() error {
	s.logger.Infof("customer-processor: Listening on '%s'", s.appConfig.ListeningInterface)

	return http.ListenAndServe(s.appConfig.ListeningInterface, s.GetHandler())
}

func (s *httpServer) applyCors(mux *http.ServeMux) http.Handler {
	return cors.AllowAll().Handler(mux)
}

func (s *httpServer) setResponseTypeToJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
