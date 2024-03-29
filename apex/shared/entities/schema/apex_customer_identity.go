package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/sigmasee/sigmasee/shared/enterprise/entities"
)

type ApexCustomerIdentity struct {
	ent.Schema
}

func (ApexCustomerIdentity) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Annotations(entgql.OrderField("id")),
		field.String("email").Optional().Annotations(entgql.OrderField("email")),
		field.Bool("email_verified").Optional().Annotations(entgql.OrderField("emailVerified")),
	}
}

func (ApexCustomerIdentity) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entities.BaseEntity{},
	}
}

func (ApexCustomerIdentity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", ApexCustomer.Type).Ref("identities").Unique().Required(),
	}
}

func (ApexCustomerIdentity) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
		index.Fields("email_verified"),

		index.Edges("customer"),
	}
}
