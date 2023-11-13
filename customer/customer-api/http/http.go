package http

import (
	"net/http"

	"github.com/rs/cors"
	graphqlv1 "github.com/sigmasee/sigmasee/customer/customer-api/graphql/v1"
	openapi "github.com/sigmasee/sigmasee/customer/customer-api/openapi"
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
	graphQLServerV1   graphqlv1.GraphQLServer
	openApiCustomerV1 openapi.OpenApiCustomerV1
}

func NewHttpServer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	graphQLServerV1 graphqlv1.GraphQLServer,
	openApiCustomerV1 openapi.OpenApiCustomerV1) (HttpServer, error) {
	return &httpServer{
		appConfig:         appConfig,
		graphQLServerV1:   graphQLServerV1,
		logger:            logger,
		openApiCustomerV1: openApiCustomerV1,
	}, nil
}

func (s *httpServer) GetHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/customer/api/v1/graphql", s.graphQLServerV1.GetServer())
	mux.Handle("/", s.setResponseTypeToJSONMiddleware(s.openApiCustomerV1.GetHttpHandler()))

	return s.applyCors(mux)
}

func (s *httpServer) ListenAndServe() error {
	s.logger.Infof("customer-api: Listening on '%s'", s.appConfig.ListeningInterface)

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
