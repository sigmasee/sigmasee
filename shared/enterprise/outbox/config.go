package outbox

import (
	"time"
)

type OutboxConfig struct {
	MaxRetryCount int           `yaml:"maxRetryCount" env:"sigmasee_OUTBOX_MAXRETRYCOUNT"`
	RetryDelay    time.Duration `yaml:"retryDelay" env:"sigmasee_OUTBOX_RETRYDELAY"`
}
