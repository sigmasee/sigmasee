// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customer"
	"github.com/sigmasee/sigmasee/customer/shared/entities/identity"
)

// IdentityCreate is the builder for creating a Identity entity.
type IdentityCreate struct {
	config
	mutation *IdentityMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ic *IdentityCreate) SetCreatedAt(t time.Time) *IdentityCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetModifiedAt sets the "modified_at" field.
func (ic *IdentityCreate) SetModifiedAt(t time.Time) *IdentityCreate {
	ic.mutation.SetModifiedAt(t)
	return ic
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (ic *IdentityCreate) SetNillableModifiedAt(t *time.Time) *IdentityCreate {
	if t != nil {
		ic.SetModifiedAt(*t)
	}
	return ic
}

// SetDeletedAt sets the "deleted_at" field.
func (ic *IdentityCreate) SetDeletedAt(t time.Time) *IdentityCreate {
	ic.mutation.SetDeletedAt(t)
	return ic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ic *IdentityCreate) SetNillableDeletedAt(t *time.Time) *IdentityCreate {
	if t != nil {
		ic.SetDeletedAt(*t)
	}
	return ic
}

// SetEmail sets the "email" field.
func (ic *IdentityCreate) SetEmail(s string) *IdentityCreate {
	ic.mutation.SetEmail(s)
	return ic
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (ic *IdentityCreate) SetNillableEmail(s *string) *IdentityCreate {
	if s != nil {
		ic.SetEmail(*s)
	}
	return ic
}

// SetEmailVerified sets the "email_verified" field.
func (ic *IdentityCreate) SetEmailVerified(b bool) *IdentityCreate {
	ic.mutation.SetEmailVerified(b)
	return ic
}

// SetNillableEmailVerified sets the "email_verified" field if the given value is not nil.
func (ic *IdentityCreate) SetNillableEmailVerified(b *bool) *IdentityCreate {
	if b != nil {
		ic.SetEmailVerified(*b)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IdentityCreate) SetID(s string) *IdentityCreate {
	ic.mutation.SetID(s)
	return ic
}

// SetCustomerID sets the "customer" edge to the Customer entity by ID.
func (ic *IdentityCreate) SetCustomerID(id string) *IdentityCreate {
	ic.mutation.SetCustomerID(id)
	return ic
}

// SetCustomer sets the "customer" edge to the Customer entity.
func (ic *IdentityCreate) SetCustomer(c *Customer) *IdentityCreate {
	return ic.SetCustomerID(c.ID)
}

// Mutation returns the IdentityMutation object of the builder.
func (ic *IdentityCreate) Mutation() *IdentityMutation {
	return ic.mutation
}

// Save creates the Identity in the database.
func (ic *IdentityCreate) Save(ctx context.Context) (*Identity, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IdentityCreate) SaveX(ctx context.Context) *Identity {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IdentityCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IdentityCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IdentityCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`entities: missing required field "Identity.created_at"`)}
	}
	if _, ok := ic.mutation.CustomerID(); !ok {
		return &ValidationError{Name: "customer", err: errors.New(`entities: missing required edge "Identity.customer"`)}
	}
	return nil
}

func (ic *IdentityCreate) sqlSave(ctx context.Context) (*Identity, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Identity.ID type: %T", _spec.ID.Value)
		}
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IdentityCreate) createSpec() (*Identity, *sqlgraph.CreateSpec) {
	var (
		_node = &Identity{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(identity.Table, sqlgraph.NewFieldSpec(identity.FieldID, field.TypeString))
	)
	_spec.Schema = ic.schemaConfig.Identity
	_spec.OnConflict = ic.conflict
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(identity.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.ModifiedAt(); ok {
		_spec.SetField(identity.FieldModifiedAt, field.TypeTime, value)
		_node.ModifiedAt = value
	}
	if value, ok := ic.mutation.DeletedAt(); ok {
		_spec.SetField(identity.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ic.mutation.Email(); ok {
		_spec.SetField(identity.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ic.mutation.EmailVerified(); ok {
		_spec.SetField(identity.FieldEmailVerified, field.TypeBool, value)
		_node.EmailVerified = value
	}
	if nodes := ic.mutation.CustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   identity.CustomerTable,
			Columns: []string{identity.CustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeString),
			},
		}
		edge.Schema = ic.schemaConfig.Identity
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.customer_identities = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Identity.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.IdentityUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ic *IdentityCreate) OnConflict(opts ...sql.ConflictOption) *IdentityUpsertOne {
	ic.conflict = opts
	return &IdentityUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Identity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ic *IdentityCreate) OnConflictColumns(columns ...string) *IdentityUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &IdentityUpsertOne{
		create: ic,
	}
}

type (
	// IdentityUpsertOne is the builder for "upsert"-ing
	//  one Identity node.
	IdentityUpsertOne struct {
		create *IdentityCreate
	}

	// IdentityUpsert is the "OnConflict" setter.
	IdentityUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *IdentityUpsert) SetCreatedAt(v time.Time) *IdentityUpsert {
	u.Set(identity.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *IdentityUpsert) UpdateCreatedAt() *IdentityUpsert {
	u.SetExcluded(identity.FieldCreatedAt)
	return u
}

// SetModifiedAt sets the "modified_at" field.
func (u *IdentityUpsert) SetModifiedAt(v time.Time) *IdentityUpsert {
	u.Set(identity.FieldModifiedAt, v)
	return u
}

// UpdateModifiedAt sets the "modified_at" field to the value that was provided on create.
func (u *IdentityUpsert) UpdateModifiedAt() *IdentityUpsert {
	u.SetExcluded(identity.FieldModifiedAt)
	return u
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (u *IdentityUpsert) ClearModifiedAt() *IdentityUpsert {
	u.SetNull(identity.FieldModifiedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *IdentityUpsert) SetDeletedAt(v time.Time) *IdentityUpsert {
	u.Set(identity.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *IdentityUpsert) UpdateDeletedAt() *IdentityUpsert {
	u.SetExcluded(identity.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *IdentityUpsert) ClearDeletedAt() *IdentityUpsert {
	u.SetNull(identity.FieldDeletedAt)
	return u
}

// SetEmail sets the "email" field.
func (u *IdentityUpsert) SetEmail(v string) *IdentityUpsert {
	u.Set(identity.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *IdentityUpsert) UpdateEmail() *IdentityUpsert {
	u.SetExcluded(identity.FieldEmail)
	return u
}

// ClearEmail clears the value of the "email" field.
func (u *IdentityUpsert) ClearEmail() *IdentityUpsert {
	u.SetNull(identity.FieldEmail)
	return u
}

// SetEmailVerified sets the "email_verified" field.
func (u *IdentityUpsert) SetEmailVerified(v bool) *IdentityUpsert {
	u.Set(identity.FieldEmailVerified, v)
	return u
}

// UpdateEmailVerified sets the "email_verified" field to the value that was provided on create.
func (u *IdentityUpsert) UpdateEmailVerified() *IdentityUpsert {
	u.SetExcluded(identity.FieldEmailVerified)
	return u
}

// ClearEmailVerified clears the value of the "email_verified" field.
func (u *IdentityUpsert) ClearEmailVerified() *IdentityUpsert {
	u.SetNull(identity.FieldEmailVerified)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Identity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(identity.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *IdentityUpsertOne) UpdateNewValues() *IdentityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(identity.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Identity.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *IdentityUpsertOne) Ignore() *IdentityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IdentityUpsertOne) DoNothing() *IdentityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IdentityCreate.OnConflict
// documentation for more info.
func (u *IdentityUpsertOne) Update(set func(*IdentityUpsert)) *IdentityUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IdentityUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *IdentityUpsertOne) SetCreatedAt(v time.Time) *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *IdentityUpsertOne) UpdateCreatedAt() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetModifiedAt sets the "modified_at" field.
func (u *IdentityUpsertOne) SetModifiedAt(v time.Time) *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.SetModifiedAt(v)
	})
}

// UpdateModifiedAt sets the "modified_at" field to the value that was provided on create.
func (u *IdentityUpsertOne) UpdateModifiedAt() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateModifiedAt()
	})
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (u *IdentityUpsertOne) ClearModifiedAt() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.ClearModifiedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *IdentityUpsertOne) SetDeletedAt(v time.Time) *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *IdentityUpsertOne) UpdateDeletedAt() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *IdentityUpsertOne) ClearDeletedAt() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.ClearDeletedAt()
	})
}

// SetEmail sets the "email" field.
func (u *IdentityUpsertOne) SetEmail(v string) *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *IdentityUpsertOne) UpdateEmail() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *IdentityUpsertOne) ClearEmail() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.ClearEmail()
	})
}

// SetEmailVerified sets the "email_verified" field.
func (u *IdentityUpsertOne) SetEmailVerified(v bool) *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.SetEmailVerified(v)
	})
}

// UpdateEmailVerified sets the "email_verified" field to the value that was provided on create.
func (u *IdentityUpsertOne) UpdateEmailVerified() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateEmailVerified()
	})
}

// ClearEmailVerified clears the value of the "email_verified" field.
func (u *IdentityUpsertOne) ClearEmailVerified() *IdentityUpsertOne {
	return u.Update(func(s *IdentityUpsert) {
		s.ClearEmailVerified()
	})
}

// Exec executes the query.
func (u *IdentityUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for IdentityCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IdentityUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *IdentityUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("entities: IdentityUpsertOne.ID is not supported by MySQL driver. Use IdentityUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *IdentityUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// IdentityCreateBulk is the builder for creating many Identity entities in bulk.
type IdentityCreateBulk struct {
	config
	err      error
	builders []*IdentityCreate
	conflict []sql.ConflictOption
}

// Save creates the Identity entities in the database.
func (icb *IdentityCreateBulk) Save(ctx context.Context) ([]*Identity, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Identity, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IdentityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IdentityCreateBulk) SaveX(ctx context.Context) []*Identity {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IdentityCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IdentityCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Identity.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.IdentityUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (icb *IdentityCreateBulk) OnConflict(opts ...sql.ConflictOption) *IdentityUpsertBulk {
	icb.conflict = opts
	return &IdentityUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Identity.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (icb *IdentityCreateBulk) OnConflictColumns(columns ...string) *IdentityUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &IdentityUpsertBulk{
		create: icb,
	}
}

// IdentityUpsertBulk is the builder for "upsert"-ing
// a bulk of Identity nodes.
type IdentityUpsertBulk struct {
	create *IdentityCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Identity.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(identity.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *IdentityUpsertBulk) UpdateNewValues() *IdentityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(identity.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Identity.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *IdentityUpsertBulk) Ignore() *IdentityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *IdentityUpsertBulk) DoNothing() *IdentityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the IdentityCreateBulk.OnConflict
// documentation for more info.
func (u *IdentityUpsertBulk) Update(set func(*IdentityUpsert)) *IdentityUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&IdentityUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *IdentityUpsertBulk) SetCreatedAt(v time.Time) *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *IdentityUpsertBulk) UpdateCreatedAt() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetModifiedAt sets the "modified_at" field.
func (u *IdentityUpsertBulk) SetModifiedAt(v time.Time) *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.SetModifiedAt(v)
	})
}

// UpdateModifiedAt sets the "modified_at" field to the value that was provided on create.
func (u *IdentityUpsertBulk) UpdateModifiedAt() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateModifiedAt()
	})
}

// ClearModifiedAt clears the value of the "modified_at" field.
func (u *IdentityUpsertBulk) ClearModifiedAt() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.ClearModifiedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *IdentityUpsertBulk) SetDeletedAt(v time.Time) *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *IdentityUpsertBulk) UpdateDeletedAt() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *IdentityUpsertBulk) ClearDeletedAt() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.ClearDeletedAt()
	})
}

// SetEmail sets the "email" field.
func (u *IdentityUpsertBulk) SetEmail(v string) *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *IdentityUpsertBulk) UpdateEmail() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateEmail()
	})
}

// ClearEmail clears the value of the "email" field.
func (u *IdentityUpsertBulk) ClearEmail() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.ClearEmail()
	})
}

// SetEmailVerified sets the "email_verified" field.
func (u *IdentityUpsertBulk) SetEmailVerified(v bool) *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.SetEmailVerified(v)
	})
}

// UpdateEmailVerified sets the "email_verified" field to the value that was provided on create.
func (u *IdentityUpsertBulk) UpdateEmailVerified() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.UpdateEmailVerified()
	})
}

// ClearEmailVerified clears the value of the "email_verified" field.
func (u *IdentityUpsertBulk) ClearEmailVerified() *IdentityUpsertBulk {
	return u.Update(func(s *IdentityUpsert) {
		s.ClearEmailVerified()
	})
}

// Exec executes the query.
func (u *IdentityUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("entities: OnConflict was set for builder %d. Set it on the IdentityCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("entities: missing options for IdentityCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *IdentityUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
