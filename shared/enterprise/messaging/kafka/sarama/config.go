package kafka

type SaramaKafkaConfig struct {
	ClientId            string `yaml:"clientId" env:"SIGMASEE_KAFKA_SARAMA_CLIENT_ID"`
	BootstrapServers    string `yaml:"bootstrapServers" env:"SIGMASEE_KAFKA_SARAMA_BOOTSTRAP_SERVERS"`
	EnableSasl          bool   `yaml:"enableSasl" env:"SIGMASEE_KAFKA_SARAMA_ENABLE_SASL"`
	EnableSaslHandshake bool   `yaml:"enableSaslHandshake" env:"SIGMASEE_KAFKA_SARAMA_ENABLE_SASL_HANDSHAKE"`
	SaslAgorithm        string `yaml:"saslAgorithm" env:"SIGMASEE_KAFKA_SARAMA_SASL_ALGORITHM"`
	Username            string `yaml:"username" env:"SIGMASEE_KAFKA_SARAMA_USERNAME"`
	Password            string `yaml:"password" env:"SIGMASEE_KAFKA_SARAMA_PASSWORD"`
}
