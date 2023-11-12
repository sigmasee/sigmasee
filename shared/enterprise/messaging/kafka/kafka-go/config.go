package kafka

type KafkaGoKafkaConfig struct {
	BootstrapServers string `yaml:"bootstrapServers" env:"sigmasee_KAFKA_KAFKAGO_BOOTSTRAP_SERVERS"`
	EnableTls        bool   `yaml:"enableTls" env:"sigmasee_KAFKA_KAFKAGO_ENABLE_TLS"`
	EnableSasl       bool   `yaml:"enableSasl" env:"sigmasee_KAFKA_KAFKAGO_ENABLE_SASL"`
	SaslAgorithm     string `yaml:"saslAgorithm" env:"sigmasee_KAFKA_KAFKAGO_SASL_ALGORITHM"`
	Username         string `yaml:"username" env:"sigmasee_KAFKA_KAFKAGO_USERNAME"`
	Password         string `yaml:"password" env:"sigmasee_KAFKA_KAFKAGO_PASSWORD"`
}
