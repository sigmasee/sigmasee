// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomer"
	"github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomeridentity"
	"github.com/sigmasee/sigmasee/apex/shared/entities/predicate"

	"github.com/sigmasee/sigmasee/apex/shared/entities/internal"
)

// ApexCustomerIdentityUpdate is the builder for updating ApexCustomerIdentity entities.
type ApexCustomerIdentityUpdate struct {
	config
	hooks     []Hook
	mutation  *ApexCustomerIdentityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ApexCustomerIdentityUpdate builder.
func (aciu *ApexCustomerIdentityUpdate) Where(ps ...predicate.ApexCustomerIdentity) *ApexCustomerIdentityUpdate {
	aciu.mutation.Where(ps...)
	return aciu
}

// SetCreatedAt sets the "created_at" field.
func (aciu *ApexCustomerIdentityUpdate) SetCreatedAt(t time.Time) *ApexCustomerIdentityUpdate {
	aciu.mutation.SetCreatedAt(t)
	return aciu
}

// SetModifiedAt sets the "modified_at" field.
func (aciu *ApexCustomerIdentityUpdate) SetModifiedAt(t time.Time) *ApexCustomerIdentityUpdate {
	aciu.mutation.SetModifiedAt(t)
	return aciu
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (aciu *ApexCustomerIdentityUpdate) SetNillableModifiedAt(t *time.Time) *ApexCustomerIdentityUpdate {
	if t != nil {
		aciu.SetModifiedAt(*t)
	}
	return aciu
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (aciu *ApexCustomerIdentityUpdate) ClearModifiedAt() *ApexCustomerIdentityUpdate {
	aciu.mutation.ClearModifiedAt()
	return aciu
}

// SetDeletedAt sets the "deleted_at" field.
func (aciu *ApexCustomerIdentityUpdate) SetDeletedAt(t time.Time) *ApexCustomerIdentityUpdate {
	aciu.mutation.SetDeletedAt(t)
	return aciu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aciu *ApexCustomerIdentityUpdate) SetNillableDeletedAt(t *time.Time) *ApexCustomerIdentityUpdate {
	if t != nil {
		aciu.SetDeletedAt(*t)
	}
	return aciu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (aciu *ApexCustomerIdentityUpdate) ClearDeletedAt() *ApexCustomerIdentityUpdate {
	aciu.mutation.ClearDeletedAt()
	return aciu
}

// SetEmail sets the "email" field.
func (aciu *ApexCustomerIdentityUpdate) SetEmail(s string) *ApexCustomerIdentityUpdate {
	aciu.mutation.SetEmail(s)
	return aciu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (aciu *ApexCustomerIdentityUpdate) SetNillableEmail(s *string) *ApexCustomerIdentityUpdate {
	if s != nil {
		aciu.SetEmail(*s)
	}
	return aciu
}

// ClearEmail clears the value of the "email" field.
func (aciu *ApexCustomerIdentityUpdate) ClearEmail() *ApexCustomerIdentityUpdate {
	aciu.mutation.ClearEmail()
	return aciu
}

// SetEmailVerified sets the "email_verified" field.
func (aciu *ApexCustomerIdentityUpdate) SetEmailVerified(b bool) *ApexCustomerIdentityUpdate {
	aciu.mutation.SetEmailVerified(b)
	return aciu
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (aciu *ApexCustomerIdentityUpdate) SetNillableEmailVerified(b *bool) *ApexCustomerIdentityUpdate {
	if b != nil {
		aciu.SetEmailVerified(*b)
	}
	return aciu
}

// ClearEmailVerified clears the value of the "email_verified" field.
func (aciu *ApexCustomerIdentityUpdate) ClearEmailVerified() *ApexCustomerIdentityUpdate {
	aciu.mutation.ClearEmailVerified()
	return aciu
}

// SetCustomerID sets the "customer" edge to the ApexCustomer entity by ID.
func (aciu *ApexCustomerIdentityUpdate) SetCustomerID(id string) *ApexCustomerIdentityUpdate {
	aciu.mutation.SetCustomerID(id)
	return aciu
}

// SetCustomer sets the "customer" edge to the ApexCustomer entity.
func (aciu *ApexCustomerIdentityUpdate) SetCustomer(a *ApexCustomer) *ApexCustomerIdentityUpdate {
	return aciu.SetCustomerID(a.ID)
}

// Mutation returns the ApexCustomerIdentityMutation object of the builder.
func (aciu *ApexCustomerIdentityUpdate) Mutation() *ApexCustomerIdentityMutation {
	return aciu.mutation
}

// ClearCustomer clears the "customer" edge to the ApexCustomer entity.
func (aciu *ApexCustomerIdentityUpdate) ClearCustomer() *ApexCustomerIdentityUpdate {
	aciu.mutation.ClearCustomer()
	return aciu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aciu *ApexCustomerIdentityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, aciu.sqlSave, aciu.mutation, aciu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aciu *ApexCustomerIdentityUpdate) SaveX(ctx context.Context) int {
	affected, err := aciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aciu *ApexCustomerIdentityUpdate) Exec(ctx context.Context) error {
	_, err := aciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aciu *ApexCustomerIdentityUpdate) ExecX(ctx context.Context) {
	if err := aciu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aciu *ApexCustomerIdentityUpdate) check() error {
	if _, ok := aciu.mutation.CustomerID(); aciu.mutation.CustomerCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "ApexCustomerIdentity.customer"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aciu *ApexCustomerIdentityUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApexCustomerIdentityUpdate {
	aciu.modifiers = append(aciu.modifiers, modifiers...)
	return aciu
}

func (aciu *ApexCustomerIdentityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aciu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(apexcustomeridentity.Table, apexcustomeridentity.Columns, sqlgraph.NewFieldSpec(apexcustomeridentity.FieldID, field.TypeString))
	if ps := aciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aciu.mutation.CreatedAt(); ok {
		_spec.SetField(apexcustomeridentity.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := aciu.mutation.ModifiedAt(); ok {
		_spec.SetField(apexcustomeridentity.FieldModifiedAt, field.TypeTime, value)
	}
	if aciu.mutation.ModifiedAtCleared() {
		_spec.ClearField(apexcustomeridentity.FieldModifiedAt, field.TypeTime)
	}
	if value, ok := aciu.mutation.DeletedAt(); ok {
		_spec.SetField(apexcustomeridentity.FieldDeletedAt, field.TypeTime, value)
	}
	if aciu.mutation.DeletedAtCleared() {
		_spec.ClearField(apexcustomeridentity.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := aciu.mutation.Email(); ok {
		_spec.SetField(apexcustomeridentity.FieldEmail, field.TypeString, value)
	}
	if aciu.mutation.EmailCleared() {
		_spec.ClearField(apexcustomeridentity.FieldEmail, field.TypeString)
	}
	if value, ok := aciu.mutation.EmailVerified(); ok {
		_spec.SetField(apexcustomeridentity.FieldEmailVerified, field.TypeBool, value)
	}
	if aciu.mutation.EmailVerifiedCleared() {
		_spec.ClearField(apexcustomeridentity.FieldEmailVerified, field.TypeBool)
	}
	if aciu.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apexcustomeridentity.CustomerTable,
			Columns: []string{apexcustomeridentity.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apexcustomer.FieldID, field.TypeString),
			},
		}
		edge.Schema = aciu.schemaConfig.ApexCustomerIdentity
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aciu.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apexcustomeridentity.CustomerTable,
			Columns: []string{apexcustomeridentity.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apexcustomer.FieldID, field.TypeString),
			},
		}
		edge.Schema = aciu.schemaConfig.ApexCustomerIdentity
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = aciu.schemaConfig.ApexCustomerIdentity
	ctx = internal.NewSchemaConfigContext(ctx, aciu.schemaConfig)
	_spec.AddModifiers(aciu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, aciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apexcustomeridentity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aciu.mutation.done = true
	return n, nil
}

// ApexCustomerIdentityUpdateOne is the builder for updating a single ApexCustomerIdentity entity.
type ApexCustomerIdentityUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ApexCustomerIdentityMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (aciuo *ApexCustomerIdentityUpdateOne) SetCreatedAt(t time.Time) *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.SetCreatedAt(t)
	return aciuo
}

// SetModifiedAt sets the "modified_at" field.
func (aciuo *ApexCustomerIdentityUpdateOne) SetModifiedAt(t time.Time) *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.SetModifiedAt(t)
	return aciuo
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (aciuo *ApexCustomerIdentityUpdateOne) SetNillableModifiedAt(t *time.Time) *ApexCustomerIdentityUpdateOne {
	if t != nil {
		aciuo.SetModifiedAt(*t)
	}
	return aciuo
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (aciuo *ApexCustomerIdentityUpdateOne) ClearModifiedAt() *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.ClearModifiedAt()
	return aciuo
}

// SetDeletedAt sets the "deleted_at" field.
func (aciuo *ApexCustomerIdentityUpdateOne) SetDeletedAt(t time.Time) *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.SetDeletedAt(t)
	return aciuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aciuo *ApexCustomerIdentityUpdateOne) SetNillableDeletedAt(t *time.Time) *ApexCustomerIdentityUpdateOne {
	if t != nil {
		aciuo.SetDeletedAt(*t)
	}
	return aciuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (aciuo *ApexCustomerIdentityUpdateOne) ClearDeletedAt() *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.ClearDeletedAt()
	return aciuo
}

// SetEmail sets the "email" field.
func (aciuo *ApexCustomerIdentityUpdateOne) SetEmail(s string) *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.SetEmail(s)
	return aciuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (aciuo *ApexCustomerIdentityUpdateOne) SetNillableEmail(s *string) *ApexCustomerIdentityUpdateOne {
	if s != nil {
		aciuo.SetEmail(*s)
	}
	return aciuo
}

// ClearEmail clears the value of the "email" field.
func (aciuo *ApexCustomerIdentityUpdateOne) ClearEmail() *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.ClearEmail()
	return aciuo
}

// SetEmailVerified sets the "email_verified" field.
func (aciuo *ApexCustomerIdentityUpdateOne) SetEmailVerified(b bool) *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.SetEmailVerified(b)
	return aciuo
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (aciuo *ApexCustomerIdentityUpdateOne) SetNillableEmailVerified(b *bool) *ApexCustomerIdentityUpdateOne {
	if b != nil {
		aciuo.SetEmailVerified(*b)
	}
	return aciuo
}

// ClearEmailVerified clears the value of the "email_verified" field.
func (aciuo *ApexCustomerIdentityUpdateOne) ClearEmailVerified() *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.ClearEmailVerified()
	return aciuo
}

// SetCustomerID sets the "customer" edge to the ApexCustomer entity by ID.
func (aciuo *ApexCustomerIdentityUpdateOne) SetCustomerID(id string) *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.SetCustomerID(id)
	return aciuo
}

// SetCustomer sets the "customer" edge to the ApexCustomer entity.
func (aciuo *ApexCustomerIdentityUpdateOne) SetCustomer(a *ApexCustomer) *ApexCustomerIdentityUpdateOne {
	return aciuo.SetCustomerID(a.ID)
}

// Mutation returns the ApexCustomerIdentityMutation object of the builder.
func (aciuo *ApexCustomerIdentityUpdateOne) Mutation() *ApexCustomerIdentityMutation {
	return aciuo.mutation
}

// ClearCustomer clears the "customer" edge to the ApexCustomer entity.
func (aciuo *ApexCustomerIdentityUpdateOne) ClearCustomer() *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.ClearCustomer()
	return aciuo
}

// Where appends a list predicates to the ApexCustomerIdentityUpdate builder.
func (aciuo *ApexCustomerIdentityUpdateOne) Where(ps ...predicate.ApexCustomerIdentity) *ApexCustomerIdentityUpdateOne {
	aciuo.mutation.Where(ps...)
	return aciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aciuo *ApexCustomerIdentityUpdateOne) Select(field string, fields ...string) *ApexCustomerIdentityUpdateOne {
	aciuo.fields = append([]string{field}, fields...)
	return aciuo
}

// Save executes the query and returns the updated ApexCustomerIdentity entity.
func (aciuo *ApexCustomerIdentityUpdateOne) Save(ctx context.Context) (*ApexCustomerIdentity, error) {
	return withHooks(ctx, aciuo.sqlSave, aciuo.mutation, aciuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aciuo *ApexCustomerIdentityUpdateOne) SaveX(ctx context.Context) *ApexCustomerIdentity {
	node, err := aciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aciuo *ApexCustomerIdentityUpdateOne) Exec(ctx context.Context) error {
	_, err := aciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aciuo *ApexCustomerIdentityUpdateOne) ExecX(ctx context.Context) {
	if err := aciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aciuo *ApexCustomerIdentityUpdateOne) check() error {
	if _, ok := aciuo.mutation.CustomerID(); aciuo.mutation.CustomerCleared() && !ok {
		return errors.New(`entities: clearing a required unique edge "ApexCustomerIdentity.customer"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aciuo *ApexCustomerIdentityUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ApexCustomerIdentityUpdateOne {
	aciuo.modifiers = append(aciuo.modifiers, modifiers...)
	return aciuo
}

func (aciuo *ApexCustomerIdentityUpdateOne) sqlSave(ctx context.Context) (_node *ApexCustomerIdentity, err error) {
	if err := aciuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(apexcustomeridentity.Table, apexcustomeridentity.Columns, sqlgraph.NewFieldSpec(apexcustomeridentity.FieldID, field.TypeString))
	id, ok := aciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entities: missing "ApexCustomerIdentity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apexcustomeridentity.FieldID)
		for _, f := range fields {
			if !apexcustomeridentity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
			}
			if f != apexcustomeridentity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aciuo.mutation.CreatedAt(); ok {
		_spec.SetField(apexcustomeridentity.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := aciuo.mutation.ModifiedAt(); ok {
		_spec.SetField(apexcustomeridentity.FieldModifiedAt, field.TypeTime, value)
	}
	if aciuo.mutation.ModifiedAtCleared() {
		_spec.ClearField(apexcustomeridentity.FieldModifiedAt, field.TypeTime)
	}
	if value, ok := aciuo.mutation.DeletedAt(); ok {
		_spec.SetField(apexcustomeridentity.FieldDeletedAt, field.TypeTime, value)
	}
	if aciuo.mutation.DeletedAtCleared() {
		_spec.ClearField(apexcustomeridentity.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := aciuo.mutation.Email(); ok {
		_spec.SetField(apexcustomeridentity.FieldEmail, field.TypeString, value)
	}
	if aciuo.mutation.EmailCleared() {
		_spec.ClearField(apexcustomeridentity.FieldEmail, field.TypeString)
	}
	if value, ok := aciuo.mutation.EmailVerified(); ok {
		_spec.SetField(apexcustomeridentity.FieldEmailVerified, field.TypeBool, value)
	}
	if aciuo.mutation.EmailVerifiedCleared() {
		_spec.ClearField(apexcustomeridentity.FieldEmailVerified, field.TypeBool)
	}
	if aciuo.mutation.CustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apexcustomeridentity.CustomerTable,
			Columns: []string{apexcustomeridentity.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apexcustomer.FieldID, field.TypeString),
			},
		}
		edge.Schema = aciuo.schemaConfig.ApexCustomerIdentity
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aciuo.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apexcustomeridentity.CustomerTable,
			Columns: []string{apexcustomeridentity.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apexcustomer.FieldID, field.TypeString),
			},
		}
		edge.Schema = aciuo.schemaConfig.ApexCustomerIdentity
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = aciuo.schemaConfig.ApexCustomerIdentity
	ctx = internal.NewSchemaConfigContext(ctx, aciuo.schemaConfig)
	_spec.AddModifiers(aciuo.modifiers...)
	_node = &ApexCustomerIdentity{config: aciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apexcustomeridentity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aciuo.mutation.done = true
	return _node, nil
}
