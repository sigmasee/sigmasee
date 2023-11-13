package graphqlv1

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/sigmasee/sigmasee/customer/customer-api/graphql/v1/generated"
	"github.com/sigmasee/sigmasee/customer/customer-api/mappers"
	"github.com/sigmasee/sigmasee/customer/customer-api/services"
	"github.com/sigmasee/sigmasee/customer/shared/entities"
	"github.com/sigmasee/sigmasee/customer/shared/repositories"
	enterprisecontext "github.com/sigmasee/sigmasee/shared/enterprise/context"
	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
	"github.com/sigmasee/sigmasee/shared/enterprise/security/token"
	"go.uber.org/zap"
)

type Resolver struct {
	logger          *zap.SugaredLogger
	client          *entities.Client
	contextHelper   enterprisecontext.ContextHelper
	customerService services.CustomerService
	tokenService    token.TokenService
	mapper          mappers.Mapper
}

type GraphQLServer interface {
	GetServer() *handler.Server
}

type graphQLServer struct {
	executableSchema graphql.ExecutableSchema
}

const bearerTokenPrefix = "Bearer "

func NewGraphQLServer(
	logger *zap.SugaredLogger,
	entgoClient repositories.EntgoClient,
	contextHelper enterprisecontext.ContextHelper,
	customerService services.CustomerService,
	tokenService token.TokenService,
	mapper mappers.Mapper) (GraphQLServer, error) {
	executableSchema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			logger:          logger,
			client:          entgoClient.GetClient(),
			contextHelper:   contextHelper,
			customerService: customerService,
			tokenService:    tokenService,
			mapper:          mapper,
		},
	})

	return &graphQLServer{
		executableSchema: executableSchema,
	}, nil
}

func (gs *graphQLServer) GetServer() *handler.Server {
	return handler.NewDefaultServer(gs.executableSchema)
}

func (r *Resolver) addRequiredAttributesToContext(ctx context.Context) (context.Context, error) {
	if correlationId, err := r.contextHelper.GetGraphQLCorrelationId(ctx); err == nil {
		ctx = r.contextHelper.WithCorrelationId(ctx, correlationId)
	} else {
		return ctx, err
	}

	authorization, err := r.contextHelper.GetGraphQLAuthorization(ctx)
	if err != nil {
		return ctx, err
	}

	if !strings.HasPrefix(authorization, bearerTokenPrefix) {
		return ctx, errors.ErrTokenNotSupported
	}

	authorization = authorization[len(bearerTokenPrefix):]

	verifiableToken, err := r.tokenService.GetVerifiableToken(ctx, authorization)
	if err != nil {
		return ctx, err
	}

	return r.contextHelper.WithVerifiableToken(ctx, verifiableToken), nil
}
