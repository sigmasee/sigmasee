package messaging

import "time"

type Message struct {
	Topic     string
	Key       []byte
	Headers   map[string][]byte
	Payload   []byte
	Timestamp *time.Time
}
