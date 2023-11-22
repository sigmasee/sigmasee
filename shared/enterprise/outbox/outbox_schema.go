package outbox

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type Outbox struct {
	mixin.Schema
}

func (Outbox) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Annotations(entgql.OrderField("id")),
		field.Time("timestamp"),
		field.String("topic"),
		field.Bytes("key"),
		field.Bytes("payload"),
		field.JSON("headers", map[string][]byte{}),
		field.Int("retry_count"),
		field.Enum("status").
			NamedValues(
				"PENDING", "PENDING",
				"FAILED", "FAILED").
			Annotations(entgql.OrderField("status")),
		field.Time("last_retry").Optional(),
		field.Strings("processing_errors").Optional(),
	}
}

func (Outbox) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("last_retry", "status"),
	}
}
