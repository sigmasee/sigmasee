package token

import (
	"context"

	"github.com/life4/genesis/slices"
)

type TokenService interface {
	GetVerifiableToken(ctx context.Context, token string) (string, error)
}

type tokenVerifierService interface {
	GetVerifiableToken(ctx context.Context, token string) (string, error)
}

type tokenService struct {
	tokenVerifierServices []tokenVerifierService
}

func NewTokenService(
	cognitoTokenService CognitoTokenService,
	googleTokenService GoogleTokenService,
	slackTokenService SlackTokenService) (TokenService, error) {
	return &tokenService{
		tokenVerifierServices: []tokenVerifierService{
			// TODO: 20231114 - morteza - uncomment when cognito and google resources are deployed
			// cognitoTokenService,
			// googleTokenService,
			slackTokenService,
		},
	}, nil
}

func (s *tokenService) GetVerifiableToken(ctx context.Context, token string) (string, error) {
	verifiableToken := slices.Reduce(s.tokenVerifierServices, "", func(item tokenVerifierService, lastFoundVerifiableToken string) string {
		if len(lastFoundVerifiableToken) > 0 {
			return lastFoundVerifiableToken
		}

		lastFoundVerifiableToken, err := item.GetVerifiableToken(ctx, token)
		if err != nil {
			return ""
		}

		return lastFoundVerifiableToken
	})

	return verifiableToken, nil
}
