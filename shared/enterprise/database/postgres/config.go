package postgres

type PostgresConfig struct {
	ConnectionString   string `yaml:"connectionString" env:"sigmasee_POSTGRES_CONNECTIONSTRING"`
	MaxOpenConnections int    `yaml:"maxOpenConnections" env:"sigmasee_POSTGRES_MAX_OPEN_CONNECTIONS"`
}
