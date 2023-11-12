package messaging

import (
	"context"
)

type MessageCallback func(Message) error

type MessageConsumer interface {
	Consume(ctx context.Context, topic string, callback MessageCallback) error
}
