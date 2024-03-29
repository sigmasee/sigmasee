// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomer"
	"github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomeridentity"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ac *ApexCustomerQuery) CollectFields(ctx context.Context, satisfies ...string) (*ApexCustomerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ac, nil
	}
	if err := ac.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ac, nil
}

func (ac *ApexCustomerQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(apexcustomer.Columns))
		selectedFields = []string{apexcustomer.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "identities":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ApexCustomerIdentityClient{config: ac.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			ac.WithNamedIdentities(alias, func(wq *ApexCustomerIdentityQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[apexcustomer.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldCreatedAt)
				fieldSeen[apexcustomer.FieldCreatedAt] = struct{}{}
			}
		case "modifiedAt":
			if _, ok := fieldSeen[apexcustomer.FieldModifiedAt]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldModifiedAt)
				fieldSeen[apexcustomer.FieldModifiedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[apexcustomer.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldDeletedAt)
				fieldSeen[apexcustomer.FieldDeletedAt] = struct{}{}
			}
		case "eventRaisedAt":
			if _, ok := fieldSeen[apexcustomer.FieldEventRaisedAt]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldEventRaisedAt)
				fieldSeen[apexcustomer.FieldEventRaisedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[apexcustomer.FieldName]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldName)
				fieldSeen[apexcustomer.FieldName] = struct{}{}
			}
		case "givenName":
			if _, ok := fieldSeen[apexcustomer.FieldGivenName]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldGivenName)
				fieldSeen[apexcustomer.FieldGivenName] = struct{}{}
			}
		case "middleName":
			if _, ok := fieldSeen[apexcustomer.FieldMiddleName]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldMiddleName)
				fieldSeen[apexcustomer.FieldMiddleName] = struct{}{}
			}
		case "familyName":
			if _, ok := fieldSeen[apexcustomer.FieldFamilyName]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldFamilyName)
				fieldSeen[apexcustomer.FieldFamilyName] = struct{}{}
			}
		case "photoURL":
			if _, ok := fieldSeen[apexcustomer.FieldPhotoURL]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldPhotoURL)
				fieldSeen[apexcustomer.FieldPhotoURL] = struct{}{}
			}
		case "photoURL24":
			if _, ok := fieldSeen[apexcustomer.FieldPhotoURL24]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldPhotoURL24)
				fieldSeen[apexcustomer.FieldPhotoURL24] = struct{}{}
			}
		case "photoURL32":
			if _, ok := fieldSeen[apexcustomer.FieldPhotoURL32]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldPhotoURL32)
				fieldSeen[apexcustomer.FieldPhotoURL32] = struct{}{}
			}
		case "photoURL48":
			if _, ok := fieldSeen[apexcustomer.FieldPhotoURL48]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldPhotoURL48)
				fieldSeen[apexcustomer.FieldPhotoURL48] = struct{}{}
			}
		case "photoURL72":
			if _, ok := fieldSeen[apexcustomer.FieldPhotoURL72]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldPhotoURL72)
				fieldSeen[apexcustomer.FieldPhotoURL72] = struct{}{}
			}
		case "photoURL192":
			if _, ok := fieldSeen[apexcustomer.FieldPhotoURL192]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldPhotoURL192)
				fieldSeen[apexcustomer.FieldPhotoURL192] = struct{}{}
			}
		case "photoURL512":
			if _, ok := fieldSeen[apexcustomer.FieldPhotoURL512]; !ok {
				selectedFields = append(selectedFields, apexcustomer.FieldPhotoURL512)
				fieldSeen[apexcustomer.FieldPhotoURL512] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ac.Select(selectedFields...)
	}
	return nil
}

type apexcustomerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ApexCustomerPaginateOption
}

func newApexCustomerPaginateArgs(rv map[string]any) *apexcustomerPaginateArgs {
	args := &apexcustomerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &ApexCustomerOrder{Field: &ApexCustomerOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithApexCustomerOrder(order))
			}
		case *ApexCustomerOrder:
			if v != nil {
				args.opts = append(args.opts, WithApexCustomerOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*ApexCustomerWhereInput); ok {
		args.opts = append(args.opts, WithApexCustomerFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (aci *ApexCustomerIdentityQuery) CollectFields(ctx context.Context, satisfies ...string) (*ApexCustomerIdentityQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return aci, nil
	}
	if err := aci.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return aci, nil
}

func (aci *ApexCustomerIdentityQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(apexcustomeridentity.Columns))
		selectedFields = []string{apexcustomeridentity.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "customer":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&ApexCustomerClient{config: aci.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			aci.withCustomer = query
		case "createdAt":
			if _, ok := fieldSeen[apexcustomeridentity.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, apexcustomeridentity.FieldCreatedAt)
				fieldSeen[apexcustomeridentity.FieldCreatedAt] = struct{}{}
			}
		case "modifiedAt":
			if _, ok := fieldSeen[apexcustomeridentity.FieldModifiedAt]; !ok {
				selectedFields = append(selectedFields, apexcustomeridentity.FieldModifiedAt)
				fieldSeen[apexcustomeridentity.FieldModifiedAt] = struct{}{}
			}
		case "deletedAt":
			if _, ok := fieldSeen[apexcustomeridentity.FieldDeletedAt]; !ok {
				selectedFields = append(selectedFields, apexcustomeridentity.FieldDeletedAt)
				fieldSeen[apexcustomeridentity.FieldDeletedAt] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[apexcustomeridentity.FieldEmail]; !ok {
				selectedFields = append(selectedFields, apexcustomeridentity.FieldEmail)
				fieldSeen[apexcustomeridentity.FieldEmail] = struct{}{}
			}
		case "emailVerified":
			if _, ok := fieldSeen[apexcustomeridentity.FieldEmailVerified]; !ok {
				selectedFields = append(selectedFields, apexcustomeridentity.FieldEmailVerified)
				fieldSeen[apexcustomeridentity.FieldEmailVerified] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		aci.Select(selectedFields...)
	}
	return nil
}

type apexcustomeridentityPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ApexCustomerIdentityPaginateOption
}

func newApexCustomerIdentityPaginateArgs(rv map[string]any) *apexcustomeridentityPaginateArgs {
	args := &apexcustomeridentityPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &ApexCustomerIdentityOrder{Field: &ApexCustomerIdentityOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithApexCustomerIdentityOrder(order))
			}
		case *ApexCustomerIdentityOrder:
			if v != nil {
				args.opts = append(args.opts, WithApexCustomerIdentityOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*ApexCustomerIdentityWhereInput); ok {
		args.opts = append(args.opts, WithApexCustomerIdentityFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
