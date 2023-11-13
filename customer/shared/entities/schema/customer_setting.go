package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/sigmasee/sigmasee/shared/enterprise/entities"
)

type CustomerSetting struct {
	ent.Schema
}

func (CustomerSetting) Fields() []ent.Field {
	return append(
		[]ent.Field{
			field.String("id").Annotations(entgql.OrderField("id")),
		},
		entities.BaseEntity{}.Fields()...,
	)
}

func (CustomerSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("customer_settings").Unique().Required(),
	}
}

func (CustomerSetting) Indexes() []ent.Index {
	return append(
		[]ent.Index{
			index.Edges("customer").Unique(),
		},
		entities.BaseEntity{}.Indexes()...,
	)
}
