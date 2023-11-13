package schema

import (
	"entgo.io/ent"
	"github.com/sigmasee/sigmasee/shared/enterprise/outbox"
)

type CustomerOutbox struct {
	ent.Schema
}

func (CustomerOutbox) Fields() []ent.Field {
	return outbox.Outbox{}.Fields()
}

func (CustomerOutbox) Edges() []ent.Edge {
	return outbox.Outbox{}.Edges()
}

func (CustomerOutbox) Indexes() []ent.Index {
	return outbox.Outbox{}.Indexes()
}
