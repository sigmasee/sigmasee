package configuration

import (
	domainconfiguration "github.com/sigmasee/sigmasee/customer/shared/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"github.com/sigmasee/sigmasee/shared/enterprise/database/postgres"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
	"github.com/sigmasee/sigmasee/shared/enterprise/outbox"
)

type Config struct {
	App       configuration.AppConfig         `yaml:"app"`
	AwsLambda configuration.AwsLambdaConfig   `yaml:"awsLambda"`
	Email     domainconfiguration.EmailConfig `yaml:"emailConfig"`
	Database  database.DatabaseConfig         `yaml:"database"`
	Postgres  postgres.PostgresConfig         `yaml:"postgres"`
	Kafka     kafka.KafkaGoKafkaConfig        `yaml:"kafka"`
	Outbox    outbox.OutboxConfig             `yaml:"outbox"`
}
