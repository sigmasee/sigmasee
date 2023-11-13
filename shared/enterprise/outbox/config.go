package outbox

import (
	"time"
)

type OutboxConfig struct {
	MaxRetryCount int           `yaml:"maxRetryCount" env:"SIGMASEE_OUTBOX_MAXRETRYCOUNT"`
	RetryDelay    time.Duration `yaml:"retryDelay" env:"SIGMASEE_OUTBOX_RETRYDELAY"`
}
