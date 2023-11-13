package configuration

import (
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"github.com/sigmasee/sigmasee/shared/enterprise/database/postgres"
	"github.com/sigmasee/sigmasee/shared/enterprise/security/token"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Database database.DatabaseConfig `yaml:"database"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`

	CognitoIdentityProvider token.CognitoConfig `yaml:"cognitoIdentityProvider"`
	GoogleIdentityProvider  token.GoogleConfig  `yaml:"googleIdentityProvider"`
	SlackIdentityProvider   token.SlackConfig   `yaml:"slackIdentityProvider"`
}
