package token

import (
	"context"
	"fmt"

	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
	"google.golang.org/api/idtoken"
)

type GoogleConfig struct {
	AppID string `yaml:"appId" env:"sigmasee_SECURITY_GOOGLE_APP_ID"`
}

type GoogleTokenService interface {
	GetVerifiableToken(ctx context.Context, token string) (string, error)
}

type googleTokenService struct {
	config GoogleConfig
}

func NewGoogleTokenService(config GoogleConfig) (GoogleTokenService, error) {
	return &googleTokenService{
		config: config,
	}, nil
}

func (s *googleTokenService) GetVerifiableToken(ctx context.Context, token string) (string, error) {
	payload, err := idtoken.Validate(ctx, token, s.config.AppID)
	if err != nil {
		return "", err
	}

	if value, found := payload.Claims["sub"]; found {
		return fmt.Sprintf("%v", value), nil
	}

	return "", errors.ErrNoVerifiablePieceFoundInClaim
}
