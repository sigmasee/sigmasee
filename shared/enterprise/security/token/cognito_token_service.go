package token

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
	"go.uber.org/zap"
)

type CognitoConfig struct {
	JwksUrl         string `yaml:"JwksUrl" env:"sigmasee_SECURITY_COGNITO_JWKS_URL"`
	RefreshRsaKeys  bool   `yaml:"refreshRsaKeys" env:"sigmasee_SECURITY_COGNITO_REFRESH_RSA_KEYS"`
	Issuer          string `yaml:"issuer" env:"sigmasee_SECURITY_COGNITO_ISSUER"`
	InitialJwksKeys string `yaml:"initialJwksKeys" env:"sigmasee_SECURITY_COGNITO_INITIAL_JWKS_KEYS"`
}

type CognitoTokenService interface {
	GetVerifiableToken(ctx context.Context, token string) (string, error)
}

type cognitoTokenService struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	config             CognitoConfig
	tokenHelperService *tokenHelperService
}

func NewCognitoTokenService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	jobScheduler *gocron.Scheduler,
	config CognitoConfig) (CognitoTokenService, error) {
	tokenHelperService, err := NewTokenHelperService(
		ctx,
		logger,
		jobScheduler,
		config.JwksUrl,
		config.InitialJwksKeys,
		config.RefreshRsaKeys)
	if err != nil {
		return nil, err
	}

	return &cognitoTokenService{
		ctx:                ctx,
		logger:             logger,
		config:             config,
		tokenHelperService: tokenHelperService,
	}, nil
}

func (s *cognitoTokenService) GetVerifiableToken(ctx context.Context, token string) (string, error) {
	claims, err := s.tokenHelperService.VerifyAndGetAllClaims(token)
	if err != nil {
		return "", err
	}

	if value, found := claims["iss"]; found {
		if value != s.config.Issuer {
			return "", errors.ErrNoVerifiablePieceFoundInClaim
		}
	} else {
		return "", errors.ErrNoVerifiablePieceFoundInClaim
	}

	if value, found := claims["sub"]; found {
		return value, nil
	}

	return "", errors.ErrNoVerifiablePieceFoundInClaim
}
