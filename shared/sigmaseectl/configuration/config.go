package configuration

import (
	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
	"github.com/sigmasee/sigmasee/shared/enterprise/database"
	"github.com/sigmasee/sigmasee/shared/enterprise/database/postgres"
)

type Config struct {
	App      configuration.AppConfig `yaml:"app"`
	Database database.DatabaseConfig `yaml:"database"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}
