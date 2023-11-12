package messaging

import (
	"context"
)

type MessageProducer interface {
	Produce(ctx context.Context, messages []Message) error
}
