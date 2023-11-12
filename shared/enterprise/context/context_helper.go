package context

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/pkg/errors"
)

type ContextHelper interface {
	WaitUntilCancelled(ctx context.Context)
	WithVerifiableToken(ctx context.Context, verifiableToken string) context.Context
	GetVerifiableToken(ctx context.Context) string
	GetCorrelationId(ctx context.Context) string
	WithCorrelationId(ctx context.Context, correlationID string) context.Context
	GetGraphQLAuthorization(ctx context.Context) (string, error)
	GetGraphQLCorrelationId(ctx context.Context) (string, error)
}

var ErrNoAuthorizationHeaderFound = errors.New("no " + authorizationHeader + " header found")
var ErrNoCorrelationIdHeaderFound = errors.New("no " + correlationHeader + " header found")

type contextHelper struct {
}

var authorizationHeader = http.CanonicalHeaderKey("authorization")
var correlationHeader = http.CanonicalHeaderKey("x-correlation-id")

type key string

const correlationIdCtx key = "correlation-id-ctx"
const verifiableTokenCtx key = "verifiable-token-ctx"

func NewContextHelper() (ContextHelper, error) {
	return &contextHelper{}, nil
}

func (s *contextHelper) WaitUntilCancelled(ctx context.Context) {
	for {
		if ctx.Err() == context.Canceled {
			break
		}

		select {
		case <-ctx.Done():
		case <-time.After(time.Second):
		}
	}
}

func (s *contextHelper) WithVerifiableToken(ctx context.Context, verifiableToken string) context.Context {
	return context.WithValue(ctx, verifiableTokenCtx, verifiableToken)
}

func (s *contextHelper) GetVerifiableToken(ctx context.Context) string {
	value := ctx.Value(verifiableTokenCtx)

	if value == nil {
		return ""
	}

	convertedVal := value.(string)

	return convertedVal
}

func (s *contextHelper) GetCorrelationId(ctx context.Context) string {
	value := ctx.Value(correlationIdCtx)

	if value == nil {
		return ""
	}

	convertedVal := value.(string)

	return convertedVal
}

func (s *contextHelper) WithCorrelationId(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, correlationIdCtx, correlationID)
}

func (s *contextHelper) GetGraphQLAuthorization(ctx context.Context) (string, error) {
	operationContext := graphql.GetOperationContext(ctx)
	authorizations, ok := operationContext.Headers[authorizationHeader]

	if !ok || len(authorizations) != 1 {
		return "", ErrNoAuthorizationHeaderFound
	}

	return authorizations[0], nil
}

func (s *contextHelper) GetGraphQLCorrelationId(ctx context.Context) (string, error) {
	operationContext := graphql.GetOperationContext(ctx)
	correlationId, ok := operationContext.Headers[correlationHeader]

	if !ok || len(correlationId) != 1 {
		return "", ErrNoCorrelationIdHeaderFound
	}

	return correlationId[0], nil
}
