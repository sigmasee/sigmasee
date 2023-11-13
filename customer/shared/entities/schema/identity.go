package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/sigmasee/sigmasee/shared/enterprise/entities"
)

type Identity struct {
	ent.Schema
}

func (Identity) Fields() []ent.Field {
	return append(
		[]ent.Field{
			field.String("id").Annotations(entgql.OrderField("id")),
			field.String("email").Optional().Annotations(entgql.OrderField("email")),
			field.Bool("email_verified").Optional().Annotations(entgql.OrderField("emailVerified")),
		},
		entities.BaseEntity{}.Fields()...,
	)
}

func (Identity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("identities").Unique().Required(),
	}
}

func (Identity) Indexes() []ent.Index {
	return append(
		[]ent.Index{
			index.Fields("email"),
			index.Fields("email_verified"),

			index.Edges("customer"),
		},
		entities.BaseEntity{}.Indexes()...,
	)
}
