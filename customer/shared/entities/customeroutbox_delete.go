// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sigmasee/sigmasee/customer/shared/entities/predicate"

	"github.com/sigmasee/sigmasee/customer/shared/entities/customeroutbox"
	"github.com/sigmasee/sigmasee/customer/shared/entities/internal"
)

// CustomerOutboxDelete is the builder for deleting a CustomerOutbox entity.
type CustomerOutboxDelete struct {
	config
	hooks    []Hook
	mutation *CustomerOutboxMutation
}

// Where appends a list predicates to the CustomerOutboxDelete builder.
func (cod *CustomerOutboxDelete) Where(ps ...predicate.CustomerOutbox) *CustomerOutboxDelete {
	cod.mutation.Where(ps...)
	return cod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cod *CustomerOutboxDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cod.sqlExec, cod.mutation, cod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cod *CustomerOutboxDelete) ExecX(ctx context.Context) int {
	n, err := cod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cod *CustomerOutboxDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(customeroutbox.Table, sqlgraph.NewFieldSpec(customeroutbox.FieldID, field.TypeString))
	_spec.Node.Schema = cod.schemaConfig.CustomerOutbox
	ctx = internal.NewSchemaConfigContext(ctx, cod.schemaConfig)
	if ps := cod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cod.mutation.done = true
	return affected, err
}

// CustomerOutboxDeleteOne is the builder for deleting a single CustomerOutbox entity.
type CustomerOutboxDeleteOne struct {
	cod *CustomerOutboxDelete
}

// Where appends a list predicates to the CustomerOutboxDelete builder.
func (codo *CustomerOutboxDeleteOne) Where(ps ...predicate.CustomerOutbox) *CustomerOutboxDeleteOne {
	codo.cod.mutation.Where(ps...)
	return codo
}

// Exec executes the deletion query.
func (codo *CustomerOutboxDeleteOne) Exec(ctx context.Context) error {
	n, err := codo.cod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{customeroutbox.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (codo *CustomerOutboxDeleteOne) ExecX(ctx context.Context) {
	if err := codo.Exec(ctx); err != nil {
		panic(err)
	}
}
