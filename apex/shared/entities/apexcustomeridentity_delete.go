// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sigmasee/sigmasee/apex/shared/entities/predicate"

	"github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomeridentity"
	"github.com/sigmasee/sigmasee/apex/shared/entities/internal"
)

// ApexCustomerIdentityDelete is the builder for deleting a ApexCustomerIdentity entity.
type ApexCustomerIdentityDelete struct {
	config
	hooks    []Hook
	mutation *ApexCustomerIdentityMutation
}

// Where appends a list predicates to the ApexCustomerIdentityDelete builder.
func (acid *ApexCustomerIdentityDelete) Where(ps ...predicate.ApexCustomerIdentity) *ApexCustomerIdentityDelete {
	acid.mutation.Where(ps...)
	return acid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acid *ApexCustomerIdentityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, acid.sqlExec, acid.mutation, acid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acid *ApexCustomerIdentityDelete) ExecX(ctx context.Context) int {
	n, err := acid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acid *ApexCustomerIdentityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(apexcustomeridentity.Table, sqlgraph.NewFieldSpec(apexcustomeridentity.FieldID, field.TypeString))
	_spec.Node.Schema = acid.schemaConfig.ApexCustomerIdentity
	ctx = internal.NewSchemaConfigContext(ctx, acid.schemaConfig)
	if ps := acid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acid.mutation.done = true
	return affected, err
}

// ApexCustomerIdentityDeleteOne is the builder for deleting a single ApexCustomerIdentity entity.
type ApexCustomerIdentityDeleteOne struct {
	acid *ApexCustomerIdentityDelete
}

// Where appends a list predicates to the ApexCustomerIdentityDelete builder.
func (acido *ApexCustomerIdentityDeleteOne) Where(ps ...predicate.ApexCustomerIdentity) *ApexCustomerIdentityDeleteOne {
	acido.acid.mutation.Where(ps...)
	return acido
}

// Exec executes the deletion query.
func (acido *ApexCustomerIdentityDeleteOne) Exec(ctx context.Context) error {
	n, err := acido.acid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{apexcustomeridentity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acido *ApexCustomerIdentityDeleteOne) ExecX(ctx context.Context) {
	if err := acido.Exec(ctx); err != nil {
		panic(err)
	}
}
