package configuration

import (
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"github.com/sigmasee/sigmasee/shared/enterprise/database/postgres"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	"github.com/sigmasee/sigmasee/shared/enterprise/outbox"
	"github.com/sigmasee/sigmasee/shared/enterprise/security/token"
)

type Config struct {
	App      configuration.AppConfig      `yaml:"app"`
	Intercom configuration.IntercomConfig `yaml:"intercom"`
	Database database.DatabaseConfig      `yaml:"database"`
	Postgres postgres.PostgresConfig      `yaml:"postgres"`
	Outbox   outbox.OutboxConfig          `yaml:"outbox"`
	Kafka    kafka.KafkaGoKafkaConfig     `yaml:"kafka"`

	CognitoIdentityProvider token.CognitoConfig `yaml:"cognitoIdentityProvider"`
	GoogleIdentityProvider  token.GoogleConfig  `yaml:"googleIdentityProvider"`
	SlackIdentityProvider   token.SlackConfig   `yaml:"slackIdentityProvider"`
}
