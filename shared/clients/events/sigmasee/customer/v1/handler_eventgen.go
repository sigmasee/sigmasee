// Code generated by sigmaseectl, DO NOT EDIT.

package v1

import (
	"context"
)

type Subscriber interface {
	Handle(ctx context.Context, topic string, key []byte, headers map[string][]byte, event *Event) error
}
