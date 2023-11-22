package schema

import (
	"entgo.io/ent"
	"github.com/sigmasee/sigmasee/shared/enterprise/outbox"
)

type CustomerOutbox struct {
	ent.Schema
}

func (CustomerOutbox) Mixin() []ent.Mixin {
	return []ent.Mixin{
		outbox.Outbox{},
	}
}
