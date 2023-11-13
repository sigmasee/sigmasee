package postgres

type PostgresConfig struct {
	ConnectionString   string `yaml:"connectionString" env:"SIGMASEE_POSTGRES_CONNECTIONSTRING"`
	MaxOpenConnections int    `yaml:"maxOpenConnections" env:"SIGMASEE_POSTGRES_MAX_OPEN_CONNECTIONS"`
}
