// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customeroutbox"
	"github.com/sigmasee/sigmasee/customer/shared/entities/predicate"

	"github.com/sigmasee/sigmasee/customer/shared/entities/internal"
)

// CustomerOutboxQuery is the builder for querying CustomerOutbox entities.
type CustomerOutboxQuery struct {
	config
	ctx        *QueryContext
	order      []customeroutbox.OrderOption
	inters     []Interceptor
	predicates []predicate.CustomerOutbox
	loadTotal  []func(context.Context, []*CustomerOutbox) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomerOutboxQuery builder.
func (coq *CustomerOutboxQuery) Where(ps ...predicate.CustomerOutbox) *CustomerOutboxQuery {
	coq.predicates = append(coq.predicates, ps...)
	return coq
}

// Limit the number of records to be returned by this query.
func (coq *CustomerOutboxQuery) Limit(limit int) *CustomerOutboxQuery {
	coq.ctx.Limit = &limit
	return coq
}

// Offset to start from.
func (coq *CustomerOutboxQuery) Offset(offset int) *CustomerOutboxQuery {
	coq.ctx.Offset = &offset
	return coq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (coq *CustomerOutboxQuery) Unique(unique bool) *CustomerOutboxQuery {
	coq.ctx.Unique = &unique
	return coq
}

// Order specifies how the records should be ordered.
func (coq *CustomerOutboxQuery) Order(o ...customeroutbox.OrderOption) *CustomerOutboxQuery {
	coq.order = append(coq.order, o...)
	return coq
}

// First returns the first CustomerOutbox entity from the query.
// Returns a *NotFoundError when no CustomerOutbox was found.
func (coq *CustomerOutboxQuery) First(ctx context.Context) (*CustomerOutbox, error) {
	nodes, err := coq.Limit(1).All(setContextOp(ctx, coq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customeroutbox.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (coq *CustomerOutboxQuery) FirstX(ctx context.Context) *CustomerOutbox {
	node, err := coq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CustomerOutbox ID from the query.
// Returns a *NotFoundError when no CustomerOutbox ID was found.
func (coq *CustomerOutboxQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = coq.Limit(1).IDs(setContextOp(ctx, coq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customeroutbox.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (coq *CustomerOutboxQuery) FirstIDX(ctx context.Context) string {
	id, err := coq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CustomerOutbox entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CustomerOutbox entity is found.
// Returns a *NotFoundError when no CustomerOutbox entities are found.
func (coq *CustomerOutboxQuery) Only(ctx context.Context) (*CustomerOutbox, error) {
	nodes, err := coq.Limit(2).All(setContextOp(ctx, coq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customeroutbox.Label}
	default:
		return nil, &NotSingularError{customeroutbox.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (coq *CustomerOutboxQuery) OnlyX(ctx context.Context) *CustomerOutbox {
	node, err := coq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CustomerOutbox ID in the query.
// Returns a *NotSingularError when more than one CustomerOutbox ID is found.
// Returns a *NotFoundError when no entities are found.
func (coq *CustomerOutboxQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = coq.Limit(2).IDs(setContextOp(ctx, coq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customeroutbox.Label}
	default:
		err = &NotSingularError{customeroutbox.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (coq *CustomerOutboxQuery) OnlyIDX(ctx context.Context) string {
	id, err := coq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CustomerOutboxes.
func (coq *CustomerOutboxQuery) All(ctx context.Context) ([]*CustomerOutbox, error) {
	ctx = setContextOp(ctx, coq.ctx, "All")
	if err := coq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CustomerOutbox, *CustomerOutboxQuery]()
	return withInterceptors[[]*CustomerOutbox](ctx, coq, qr, coq.inters)
}

// AllX is like All, but panics if an error occurs.
func (coq *CustomerOutboxQuery) AllX(ctx context.Context) []*CustomerOutbox {
	nodes, err := coq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CustomerOutbox IDs.
func (coq *CustomerOutboxQuery) IDs(ctx context.Context) (ids []string, err error) {
	if coq.ctx.Unique == nil && coq.path != nil {
		coq.Unique(true)
	}
	ctx = setContextOp(ctx, coq.ctx, "IDs")
	if err = coq.Select(customeroutbox.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (coq *CustomerOutboxQuery) IDsX(ctx context.Context) []string {
	ids, err := coq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (coq *CustomerOutboxQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, coq.ctx, "Count")
	if err := coq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, coq, querierCount[*CustomerOutboxQuery](), coq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (coq *CustomerOutboxQuery) CountX(ctx context.Context) int {
	count, err := coq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (coq *CustomerOutboxQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, coq.ctx, "Exist")
	switch _, err := coq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entities: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (coq *CustomerOutboxQuery) ExistX(ctx context.Context) bool {
	exist, err := coq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomerOutboxQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (coq *CustomerOutboxQuery) Clone() *CustomerOutboxQuery {
	if coq == nil {
		return nil
	}
	return &CustomerOutboxQuery{
		config:     coq.config,
		ctx:        coq.ctx.Clone(),
		order:      append([]customeroutbox.OrderOption{}, coq.order...),
		inters:     append([]Interceptor{}, coq.inters...),
		predicates: append([]predicate.CustomerOutbox{}, coq.predicates...),
		// clone intermediate query.
		sql:  coq.sql.Clone(),
		path: coq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Timestamp time.Time `json:"timestamp,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CustomerOutbox.Query().
//		GroupBy(customeroutbox.FieldTimestamp).
//		Aggregate(entities.Count()).
//		Scan(ctx, &v)
func (coq *CustomerOutboxQuery) GroupBy(field string, fields ...string) *CustomerOutboxGroupBy {
	coq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CustomerOutboxGroupBy{build: coq}
	grbuild.flds = &coq.ctx.Fields
	grbuild.label = customeroutbox.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Timestamp time.Time `json:"timestamp,omitempty"`
//	}
//
//	client.CustomerOutbox.Query().
//		Select(customeroutbox.FieldTimestamp).
//		Scan(ctx, &v)
func (coq *CustomerOutboxQuery) Select(fields ...string) *CustomerOutboxSelect {
	coq.ctx.Fields = append(coq.ctx.Fields, fields...)
	sbuild := &CustomerOutboxSelect{CustomerOutboxQuery: coq}
	sbuild.label = customeroutbox.Label
	sbuild.flds, sbuild.scan = &coq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CustomerOutboxSelect configured with the given aggregations.
func (coq *CustomerOutboxQuery) Aggregate(fns ...AggregateFunc) *CustomerOutboxSelect {
	return coq.Select().Aggregate(fns...)
}

func (coq *CustomerOutboxQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range coq.inters {
		if inter == nil {
			return fmt.Errorf("entities: uninitialized interceptor (forgotten import entities/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, coq); err != nil {
				return err
			}
		}
	}
	for _, f := range coq.ctx.Fields {
		if !customeroutbox.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
		}
	}
	if coq.path != nil {
		prev, err := coq.path(ctx)
		if err != nil {
			return err
		}
		coq.sql = prev
	}
	return nil
}

func (coq *CustomerOutboxQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CustomerOutbox, error) {
	var (
		nodes = []*CustomerOutbox{}
		_spec = coq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CustomerOutbox).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CustomerOutbox{config: coq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = coq.schemaConfig.CustomerOutbox
	ctx = internal.NewSchemaConfigContext(ctx, coq.schemaConfig)
	if len(coq.modifiers) > 0 {
		_spec.Modifiers = coq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, coq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range coq.loadTotal {
		if err := coq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (coq *CustomerOutboxQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := coq.querySpec()
	_spec.Node.Schema = coq.schemaConfig.CustomerOutbox
	ctx = internal.NewSchemaConfigContext(ctx, coq.schemaConfig)
	if len(coq.modifiers) > 0 {
		_spec.Modifiers = coq.modifiers
	}
	_spec.Node.Columns = coq.ctx.Fields
	if len(coq.ctx.Fields) > 0 {
		_spec.Unique = coq.ctx.Unique != nil && *coq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, coq.driver, _spec)
}

func (coq *CustomerOutboxQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(customeroutbox.Table, customeroutbox.Columns, sqlgraph.NewFieldSpec(customeroutbox.FieldID, field.TypeString))
	_spec.From = coq.sql
	if unique := coq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if coq.path != nil {
		_spec.Unique = true
	}
	if fields := coq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customeroutbox.FieldID)
		for i := range fields {
			if fields[i] != customeroutbox.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := coq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := coq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := coq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := coq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (coq *CustomerOutboxQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(coq.driver.Dialect())
	t1 := builder.Table(customeroutbox.Table)
	columns := coq.ctx.Fields
	if len(columns) == 0 {
		columns = customeroutbox.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if coq.sql != nil {
		selector = coq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if coq.ctx.Unique != nil && *coq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(coq.schemaConfig.CustomerOutbox)
	ctx = internal.NewSchemaConfigContext(ctx, coq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range coq.modifiers {
		m(selector)
	}
	for _, p := range coq.predicates {
		p(selector)
	}
	for _, p := range coq.order {
		p(selector)
	}
	if offset := coq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := coq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (coq *CustomerOutboxQuery) ForUpdate(opts ...sql.LockOption) *CustomerOutboxQuery {
	if coq.driver.Dialect() == dialect.Postgres {
		coq.Unique(false)
	}
	coq.modifiers = append(coq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return coq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (coq *CustomerOutboxQuery) ForShare(opts ...sql.LockOption) *CustomerOutboxQuery {
	if coq.driver.Dialect() == dialect.Postgres {
		coq.Unique(false)
	}
	coq.modifiers = append(coq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return coq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (coq *CustomerOutboxQuery) Modify(modifiers ...func(s *sql.Selector)) *CustomerOutboxSelect {
	coq.modifiers = append(coq.modifiers, modifiers...)
	return coq.Select()
}

// CustomerOutboxGroupBy is the group-by builder for CustomerOutbox entities.
type CustomerOutboxGroupBy struct {
	selector
	build *CustomerOutboxQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cogb *CustomerOutboxGroupBy) Aggregate(fns ...AggregateFunc) *CustomerOutboxGroupBy {
	cogb.fns = append(cogb.fns, fns...)
	return cogb
}

// Scan applies the selector query and scans the result into the given value.
func (cogb *CustomerOutboxGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cogb.build.ctx, "GroupBy")
	if err := cogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerOutboxQuery, *CustomerOutboxGroupBy](ctx, cogb.build, cogb, cogb.build.inters, v)
}

func (cogb *CustomerOutboxGroupBy) sqlScan(ctx context.Context, root *CustomerOutboxQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cogb.fns))
	for _, fn := range cogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cogb.flds)+len(cogb.fns))
		for _, f := range *cogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CustomerOutboxSelect is the builder for selecting fields of CustomerOutbox entities.
type CustomerOutboxSelect struct {
	*CustomerOutboxQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cos *CustomerOutboxSelect) Aggregate(fns ...AggregateFunc) *CustomerOutboxSelect {
	cos.fns = append(cos.fns, fns...)
	return cos
}

// Scan applies the selector query and scans the result into the given value.
func (cos *CustomerOutboxSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cos.ctx, "Select")
	if err := cos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerOutboxQuery, *CustomerOutboxSelect](ctx, cos.CustomerOutboxQuery, cos, cos.inters, v)
}

func (cos *CustomerOutboxSelect) sqlScan(ctx context.Context, root *CustomerOutboxQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cos.fns))
	for _, fn := range cos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cos *CustomerOutboxSelect) Modify(modifiers ...func(s *sql.Selector)) *CustomerOutboxSelect {
	cos.modifiers = append(cos.modifiers, modifiers...)
	return cos
}
