package configuration

import (
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"github.com/sigmasee/sigmasee/shared/enterprise/database/postgres"
	"github.com/sigmasee/sigmasee/shared/enterprise/messaging/kafka/kafka-go"
)

type Config struct {
	App      configuration.AppConfig  `yaml:"app"`
	Database database.DatabaseConfig  `yaml:"database"`
	Postgres postgres.PostgresConfig  `yaml:"postgres"`
	Kafka    kafka.KafkaGoKafkaConfig `yaml:"kafka"`
}
