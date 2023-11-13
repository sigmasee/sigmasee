package graphqlv1

import (
	"context"
	"net/http"

	"github.com/gobwas/ws"
	log "github.com/jensneuse/abstractlogger"
	http2 "github.com/wundergraph/graphql-go-tools/examples/federation/gateway/http"
	"github.com/wundergraph/graphql-go-tools/pkg/engine/datasource/graphql_datasource"
	"github.com/wundergraph/graphql-go-tools/pkg/graphql"
	"go.uber.org/zap"
)

type Gateway interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	UpdateDataSources(ctx context.Context, newDataSourcesConfig []graphql_datasource.Configuration)
}

type gateway struct {
	logger            *zap.SugaredLogger
	abstractLogger    log.Logger
	gqlHandlerFactory HandlerFactory
	httpClient        *http.Client
	gqlHandler        http.Handler
}

type DataSourceObserver interface {
	UpdateDataSources(ctx context.Context, dataSourceConfig []graphql_datasource.Configuration)
}

type DataSourceSubject interface {
	Register(observer DataSourceObserver)
}

type HandlerFactory interface {
	Make(schema *graphql.Schema, engine *graphql.ExecutionEngineV2) http.Handler
}

type HandlerFactoryFn func(schema *graphql.Schema, engine *graphql.ExecutionEngineV2) http.Handler

func (s HandlerFactoryFn) Make(schema *graphql.Schema, engine *graphql.ExecutionEngineV2) http.Handler {
	return s(schema, engine)
}

func NewGateway(
	logger *zap.SugaredLogger,
	zapLogger *zap.Logger,
	httpClient *http.Client) (Gateway, error) {
	abstractLogger := log.NewZapLogger(zapLogger, log.InfoLevel)

	upgrader := &ws.DefaultHTTPUpgrader
	upgrader.Header = http.Header{}
	upgrader.Header.Add("Sec-Websocket-Protocol", "graphql-ws")

	var gqlHandlerFactory HandlerFactoryFn = func(schema *graphql.Schema, engine *graphql.ExecutionEngineV2) http.Handler {
		return http2.NewGraphqlHTTPHandler(schema, engine, upgrader, abstractLogger)
	}

	return &gateway{
		logger:            logger,
		abstractLogger:    abstractLogger,
		gqlHandlerFactory: gqlHandlerFactory,
		httpClient:        httpClient,
	}, nil
}

func (s *gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := s.gqlHandler

	handler.ServeHTTP(w, r)
}

func (s *gateway) UpdateDataSources(ctx context.Context, newDataSourcesConfig []graphql_datasource.Configuration) {
	engineConfigFactory := graphql.NewFederationEngineConfigFactory(
		newDataSourcesConfig,
		graphql_datasource.NewBatchFactory(),
		graphql.WithFederationHttpClient(s.httpClient))

	schema, err := engineConfigFactory.MergedSchema()
	if err != nil {
		s.logger.Errorf("get schema:", err)

		return
	}

	datasourceConfig, err := engineConfigFactory.EngineV2Configuration()
	if err != nil {
		s.logger.Errorf("get engine config: %v", err)

		return
	}

	datasourceConfig.EnableDataLoader(true)

	engine, err := graphql.NewExecutionEngineV2(ctx, s.abstractLogger, datasourceConfig)
	if err != nil {
		s.logger.Errorf("create engine: %v", err)

		return
	}

	s.gqlHandler = s.gqlHandlerFactory.Make(schema, engine)
}
