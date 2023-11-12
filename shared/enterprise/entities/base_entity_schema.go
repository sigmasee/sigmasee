package entities

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type BaseEntity struct {
	ent.Schema
}

func (BaseEntity) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Annotations(entgql.OrderField("createdAt")),
		field.Time("modified_at").Optional().Annotations(entgql.OrderField("modifiedAt")),
		field.Time("deleted_at").Optional().Annotations(entgql.OrderField("deletedAt")),
	}
}

func (BaseEntity) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (BaseEntity) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("deleted_at"),
	}
}
