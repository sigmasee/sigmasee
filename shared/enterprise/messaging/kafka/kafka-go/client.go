package kafka

import (
	"crypto/tls"
	"fmt"
	"strings"

	kafkago "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	"go.uber.org/zap"
)

type KafkaClient interface {
	NewDialer() ([]string, *kafkago.Dialer, error)
}

type kafkaClient struct {
	logger      *zap.SugaredLogger
	kafkaConfig KafkaGoKafkaConfig
}

func NewKafkaGoKafkaClient(
	logger *zap.SugaredLogger,
	kafkaConfig KafkaGoKafkaConfig) (KafkaClient, error) {

	return &kafkaClient{
		logger:      logger,
		kafkaConfig: kafkaConfig,
	}, nil
}

func (s *kafkaClient) NewDialer() ([]string, *kafkago.Dialer, error) {
	brokers := strings.Split(s.kafkaConfig.BootstrapServers, ",")

	dialer := kafkago.Dialer{}

	if s.kafkaConfig.EnableTls {
		dialer.TLS = &tls.Config{}
	}

	if s.kafkaConfig.EnableSasl {
		if s.kafkaConfig.SaslAgorithm == "sha512" {
			mechanism, err := scram.Mechanism(scram.SHA512, s.kafkaConfig.Username, s.kafkaConfig.Password)
			if err != nil {
				return nil, nil, err
			}

			dialer.SASLMechanism = mechanism
		} else if s.kafkaConfig.SaslAgorithm == "sha256" {
			mechanism, err := scram.Mechanism(scram.SHA256, s.kafkaConfig.Username, s.kafkaConfig.Password)
			if err != nil {
				return nil, nil, err
			}

			dialer.SASLMechanism = mechanism
		} else {
			return nil, nil, fmt.Errorf("invalid SHA algorithm \"%s\": can be either \"sha256\" or \"sha512\"", s.kafkaConfig.SaslAgorithm)
		}
	}

	return brokers, &dialer, nil
}
