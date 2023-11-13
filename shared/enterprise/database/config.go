package database

type DatabaseConfig struct {
	Debug bool `yaml:"debug" env:"SIGMASEE_DATABASE_DEBUG"`
}
