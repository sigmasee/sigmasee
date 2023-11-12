package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sigmasee/sigmasee/shared/enterprise/entities"
)

type ApexCustomer struct {
	ent.Schema
}

func (ApexCustomer) Fields() []ent.Field {
	return append(
		[]ent.Field{
			field.String("id").Annotations(entgql.OrderField("id")),
			field.Time("event_raised_at").Annotations(entgql.OrderField("eventRaisedAt")),

			field.String("name").Optional().Annotations(entgql.OrderField("name")),
			field.String("given_name").Optional().Annotations(entgql.OrderField("givenName")),
			field.String("middle_name").Optional().Annotations(entgql.OrderField("middleName")),
			field.String("family_name").Optional().Annotations(entgql.OrderField("familyName")),

			field.String("photo_url").Optional().Annotations(entgql.OrderField("photoUrl")),
			field.String("photo_url_24").Optional().Annotations(entgql.OrderField("photoUrl24")),
			field.String("photo_url_32").Optional().Annotations(entgql.OrderField("photoUrl32")),
			field.String("photo_url_48").Optional().Annotations(entgql.OrderField("photoUrl48")),
			field.String("photo_url_72").Optional().Annotations(entgql.OrderField("photoUrl72")),
			field.String("photo_url_192").Optional().Annotations(entgql.OrderField("photoUrl192")),
			field.String("photo_url_512").Optional().Annotations(entgql.OrderField("photoUrl512")),
		},
		entities.BaseEntity{}.Fields()...,
	)
}

func (ApexCustomer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("identities", ApexCustomerIdentity.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (ApexCustomer) Indexes() []ent.Index {
	return append(
		[]ent.Index{},
		entities.BaseEntity{}.Indexes()...,
	)
}
