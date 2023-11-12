package database

type DatabaseConfig struct {
	Debug bool `yaml:"debug" env:"sigmasee_DATABASE_DEBUG"`
}
