package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/sigmasee/sigmasee/shared/enterprise/entities"
)

type Customer struct {
	ent.Schema
}

func (Customer) Fields() []ent.Field {
	return append(
		[]ent.Field{
			field.String("id").Annotations(entgql.OrderField("id")),

			field.String("designation").Optional().Annotations(entgql.OrderField("designation")),

			field.String("title").Optional().Annotations(entgql.OrderField("title")),
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

			field.String("timezone").Optional().Annotations(entgql.OrderField("timezone")),
			field.String("locale").Optional().Annotations(entgql.OrderField("locale")),
		},
		entities.BaseEntity{}.Fields()...,
	)
}

func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("identities", Identity.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("customer_settings", CustomerSetting.Type).Unique().Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

func (Customer) Indexes() []ent.Index {
	return append(
		[]ent.Index{
			index.Fields("designation"),
			index.Fields("title"),
			index.Fields("name"),
			index.Fields("given_name"),
			index.Fields("middle_name"),
			index.Fields("family_name"),
		},
		entities.BaseEntity{}.Indexes()...,
	)
}
