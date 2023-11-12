package kafka

import (
	"fmt"
	"strings"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type KafkaClient interface {
	NewClientConfig() ([]string, *sarama.Config, error)
}

type kafkaClient struct {
	logger      *zap.SugaredLogger
	kafkaConfig SaramaKafkaConfig
}

func NewSaramaKafkaClient(
	logger *zap.SugaredLogger,
	kafkaConfig SaramaKafkaConfig) (KafkaClient, error) {

	return &kafkaClient{
		logger:      logger,
		kafkaConfig: kafkaConfig,
	}, nil
}

func (s *kafkaClient) NewClientConfig() ([]string, *sarama.Config, error) {
	brokers := strings.Split(s.kafkaConfig.BootstrapServers, ",")

	config := sarama.NewConfig()
	config.ClientID = s.kafkaConfig.ClientId
	config.Metadata.Full = true
	config.Net.SASL.Enable = s.kafkaConfig.EnableSasl

	if s.kafkaConfig.EnableSasl {
		config.Net.SASL.Handshake = s.kafkaConfig.EnableSaslHandshake
		config.Net.SASL.User = s.kafkaConfig.Username
		config.Net.SASL.Password = s.kafkaConfig.Password

		if s.kafkaConfig.SaslAgorithm == "sha512" {
			config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &xdgScramClient{HashGeneratorFcn: SHA512} }
			config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
		} else if s.kafkaConfig.SaslAgorithm == "sha256" {
			config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &xdgScramClient{HashGeneratorFcn: SHA256} }
			config.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
		} else {
			return nil, nil, fmt.Errorf("invalid SHA algorithm \"%s\": can be either \"sha256\" or \"sha512\"", s.kafkaConfig.SaslAgorithm)
		}
	}

	return brokers, config, nil
}
