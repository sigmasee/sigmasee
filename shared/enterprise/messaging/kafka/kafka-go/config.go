package kafka

type KafkaGoKafkaConfig struct {
	BootstrapServers string `yaml:"bootstrapServers" env:"SIGMASEE_KAFKA_KAFKAGO_BOOTSTRAP_SERVERS"`
	EnableTls        bool   `yaml:"enableTls" env:"SIGMASEE_KAFKA_KAFKAGO_ENABLE_TLS"`
	EnableSasl       bool   `yaml:"enableSasl" env:"SIGMASEE_KAFKA_KAFKAGO_ENABLE_SASL"`
	SaslAgorithm     string `yaml:"saslAgorithm" env:"SIGMASEE_KAFKA_KAFKAGO_SASL_ALGORITHM"`
	Username         string `yaml:"username" env:"SIGMASEE_KAFKA_KAFKAGO_USERNAME"`
	Password         string `yaml:"password" env:"SIGMASEE_KAFKA_KAFKAGO_PASSWORD"`
}
