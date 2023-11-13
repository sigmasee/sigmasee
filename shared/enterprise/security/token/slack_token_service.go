package token

import (
	"context"
	_ "embed"

	"github.com/go-co-op/gocron"
	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
	"go.uber.org/zap"
)

type SlackConfig struct {
	RefreshRsaKeys bool   `yaml:"refreshRsaKeys" env:"SIGMASEE_SECURITY_SLACK_REFRESH_RSA_KEYS"`
	ClientID       string `yaml:"clientId" env:"SIGMASEE_SECURITY_SLACK_CLIENT_ID"`
}

type SlackTokenService interface {
	GetVerifiableToken(ctx context.Context, token string) (string, error)
}

type slackTokenService struct {
	ctx                context.Context
	logger             *zap.SugaredLogger
	config             SlackConfig
	tokenHelperService *tokenHelperService
}

//go:embed slack_jwks_keys.json
var slackJwksKeys string

const slackJwksUrl = "https://slack.com/openid/connect/keys"
const slackIssuer = "https://slack.com"

func NewSlackTokenService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	jobScheduler *gocron.Scheduler,
	config SlackConfig) (SlackTokenService, error) {
	tokenHelperService, err := NewTokenHelperService(
		ctx,
		logger,
		jobScheduler,
		slackJwksUrl,
		slackJwksKeys,
		config.RefreshRsaKeys)
	if err != nil {
		return nil, err
	}

	return &slackTokenService{
		ctx:                ctx,
		logger:             logger,
		config:             config,
		tokenHelperService: tokenHelperService,
	}, nil
}

func (s *slackTokenService) GetVerifiableToken(ctx context.Context, token string) (string, error) {
	claims, err := s.tokenHelperService.VerifyAndGetAllClaims(token)
	if err != nil {
		return "", err
	}

	if value, found := claims["iss"]; found {
		if value != slackIssuer {
			return "", errors.ErrNoVerifiablePieceFoundInClaim
		}
	} else {
		return "", errors.ErrNoVerifiablePieceFoundInClaim
	}

	if value, found := claims["aud"]; found {
		if value != s.config.ClientID {
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
