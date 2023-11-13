package graphqlv1

import (
	"context"
	"net/http"
	"net/url"

	"github.com/life4/genesis/slices"
	"github.com/wundergraph/graphql-go-tools/pkg/playground"
)

type GraphQLServer interface {
	GetHandlers(ctx context.Context) (map[string]http.Handler, error)
}

type graphQLServer struct {
	datasourcePoller DatasourcePoller
	gateway          Gateway
}

func NewGraphQLServer(
	datasourcePoller DatasourcePoller,
	gateway Gateway) (GraphQLServer, error) {

	return &graphQLServer{
		datasourcePoller: datasourcePoller,
		gateway:          gateway,
	}, nil
}

func (s *graphQLServer) GetHandlers(ctx context.Context) (map[string]http.Handler, error) {
	pathPrefix := "/gateway/api/v1"
	playgroundURLPrefix := "/playground"

	graphqlEndpoint, err := url.JoinPath(pathPrefix, "/graphql")
	if err != nil {
		return nil, err
	}

	playgroundHandlers, err := playground.
		New(playground.Config{
			PathPrefix:                      pathPrefix,
			PlaygroundPath:                  playgroundURLPrefix,
			GraphqlEndpointPath:             graphqlEndpoint,
			GraphQLSubscriptionEndpointPath: graphqlEndpoint,
		}).Handlers()
	if err != nil {
		return nil, err
	}

	handlers := slices.Reduce(
		playgroundHandlers,
		make(map[string]http.Handler),
		func(handlerConfig playground.HandlerConfig, acc map[string]http.Handler) map[string]http.Handler {
			acc[handlerConfig.Path] = handlerConfig.Handler

			return acc
		})

	handlers[graphqlEndpoint] = s.gateway

	return handlers, nil
}
