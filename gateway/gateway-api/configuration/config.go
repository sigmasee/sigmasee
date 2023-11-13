package configuration

import (
	"time"

	"github.com/sigmasee/sigmasee/shared/enterprise/configuration"
)

type Config struct {
	App          configuration.AppConfig `yaml:"app"`
	ApiEndpoints ApiEndpoints            `yaml:"apiEndpoints"`
}

type ApiEndpoints struct {
	Timeout  time.Duration `yaml:"timeout" env:"SIGMASEE_GATEWAY_TIMEOUT"`
	Customer string        `yaml:"customer" env:"CUSTOMER_ENDPOINT"`
}
