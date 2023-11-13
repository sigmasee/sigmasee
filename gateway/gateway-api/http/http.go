package http

import (
	"context"
	"net/http"

	"github.com/rs/cors"
	graphqlv1 "github.com/sigmasee/sigmasee/gateway/gateway-api/graphql/v1"
	openapi "github.com/sigmasee/sigmasee/gateway/gateway-api/openapi"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"go.uber.org/zap"
)

type HttpServer interface {
	GetHandler(ctx context.Context) (http.Handler, error)
	ListenAndServe(ctx context.Context) error
}

type httpServer struct {
	logger           *zap.SugaredLogger
	appConfig        configuration.AppConfig
	graphQLServerV1  graphqlv1.GraphQLServer
	openApiGatewayV1 openapi.OpenApiGatewayV1
}

func NewHttpServer(
	logger *zap.SugaredLogger,
	appConfig configuration.AppConfig,
	graphQLServerV1 graphqlv1.GraphQLServer,
	openApiGatewayV1 openapi.OpenApiGatewayV1) (HttpServer, error) {
	return &httpServer{
		appConfig:        appConfig,
		logger:           logger,
		graphQLServerV1:  graphQLServerV1,
		openApiGatewayV1: openApiGatewayV1,
	}, nil
}

func (s *httpServer) GetHandler(ctx context.Context) (http.Handler, error) {
	mux := http.NewServeMux()

	handlers, err := s.graphQLServerV1.GetHandlers(ctx)
	if err != nil {
		return nil, err
	}

	for path, handler := range handlers {
		mux.Handle(path, handler)
	}

	mux.Handle("/", s.setResponseTypeToJSONMiddleware(s.openApiGatewayV1.GetHttpHandler()))

	return s.applyCors(mux), nil
}

func (s *httpServer) ListenAndServe(ctx context.Context) error {
	s.logger.Infof("gateway-api: Listening on '%s'", s.appConfig.ListeningInterface)

	handler, err := s.GetHandler(ctx)
	if err != nil {
		return err
	}

	return http.ListenAndServe(s.appConfig.ListeningInterface, handler)
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
