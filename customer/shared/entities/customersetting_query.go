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
	"github.com/sigmasee/sigmasee/customer/shared/entities/customer"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customersetting"
	"github.com/sigmasee/sigmasee/customer/shared/entities/predicate"

	"github.com/sigmasee/sigmasee/customer/shared/entities/internal"
)

// CustomerSettingQuery is the builder for querying CustomerSetting entities.
type CustomerSettingQuery struct {
	config
	ctx          *QueryContext
	order        []customersetting.OrderOption
	inters       []Interceptor
	predicates   []predicate.CustomerSetting
	withCustomer *CustomerQuery
	withFKs      bool
	loadTotal    []func(context.Context, []*CustomerSetting) error
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomerSettingQuery builder.
func (csq *CustomerSettingQuery) Where(ps ...predicate.CustomerSetting) *CustomerSettingQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit the number of records to be returned by this query.
func (csq *CustomerSettingQuery) Limit(limit int) *CustomerSettingQuery {
	csq.ctx.Limit = &limit
	return csq
}

// Offset to start from.
func (csq *CustomerSettingQuery) Offset(offset int) *CustomerSettingQuery {
	csq.ctx.Offset = &offset
	return csq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csq *CustomerSettingQuery) Unique(unique bool) *CustomerSettingQuery {
	csq.ctx.Unique = &unique
	return csq
}

// Order specifies how the records should be ordered.
func (csq *CustomerSettingQuery) Order(o ...customersetting.OrderOption) *CustomerSettingQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// QueryCustomer chains the current query on the "customer" edge.
func (csq *CustomerSettingQuery) QueryCustomer() *CustomerQuery {
	query := (&CustomerClient{config: csq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := csq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customersetting.Table, customersetting.FieldID, selector),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, customersetting.CustomerTable, customersetting.CustomerColumn),
		)
		schemaConfig := csq.schemaConfig
		step.To.Schema = schemaConfig.Customer
		step.Edge.Schema = schemaConfig.CustomerSetting
		fromU = sqlgraph.SetNeighbors(csq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CustomerSetting entity from the query.
// Returns a *NotFoundError when no CustomerSetting was found.
func (csq *CustomerSettingQuery) First(ctx context.Context) (*CustomerSetting, error) {
	nodes, err := csq.Limit(1).All(setContextOp(ctx, csq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customersetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *CustomerSettingQuery) FirstX(ctx context.Context) *CustomerSetting {
	node, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CustomerSetting ID from the query.
// Returns a *NotFoundError when no CustomerSetting ID was found.
func (csq *CustomerSettingQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = csq.Limit(1).IDs(setContextOp(ctx, csq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customersetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csq *CustomerSettingQuery) FirstIDX(ctx context.Context) string {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CustomerSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CustomerSetting entity is found.
// Returns a *NotFoundError when no CustomerSetting entities are found.
func (csq *CustomerSettingQuery) Only(ctx context.Context) (*CustomerSetting, error) {
	nodes, err := csq.Limit(2).All(setContextOp(ctx, csq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customersetting.Label}
	default:
		return nil, &NotSingularError{customersetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *CustomerSettingQuery) OnlyX(ctx context.Context) *CustomerSetting {
	node, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CustomerSetting ID in the query.
// Returns a *NotSingularError when more than one CustomerSetting ID is found.
// Returns a *NotFoundError when no entities are found.
func (csq *CustomerSettingQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = csq.Limit(2).IDs(setContextOp(ctx, csq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customersetting.Label}
	default:
		err = &NotSingularError{customersetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *CustomerSettingQuery) OnlyIDX(ctx context.Context) string {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CustomerSettings.
func (csq *CustomerSettingQuery) All(ctx context.Context) ([]*CustomerSetting, error) {
	ctx = setContextOp(ctx, csq.ctx, "All")
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CustomerSetting, *CustomerSettingQuery]()
	return withInterceptors[[]*CustomerSetting](ctx, csq, qr, csq.inters)
}

// AllX is like All, but panics if an error occurs.
func (csq *CustomerSettingQuery) AllX(ctx context.Context) []*CustomerSetting {
	nodes, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CustomerSetting IDs.
func (csq *CustomerSettingQuery) IDs(ctx context.Context) (ids []string, err error) {
	if csq.ctx.Unique == nil && csq.path != nil {
		csq.Unique(true)
	}
	ctx = setContextOp(ctx, csq.ctx, "IDs")
	if err = csq.Select(customersetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *CustomerSettingQuery) IDsX(ctx context.Context) []string {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *CustomerSettingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, csq.ctx, "Count")
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, csq, querierCount[*CustomerSettingQuery](), csq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (csq *CustomerSettingQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *CustomerSettingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, csq.ctx, "Exist")
	switch _, err := csq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entities: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *CustomerSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomerSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *CustomerSettingQuery) Clone() *CustomerSettingQuery {
	if csq == nil {
		return nil
	}
	return &CustomerSettingQuery{
		config:       csq.config,
		ctx:          csq.ctx.Clone(),
		order:        append([]customersetting.OrderOption{}, csq.order...),
		inters:       append([]Interceptor{}, csq.inters...),
		predicates:   append([]predicate.CustomerSetting{}, csq.predicates...),
		withCustomer: csq.withCustomer.Clone(),
		// clone intermediate query.
		sql:  csq.sql.Clone(),
		path: csq.path,
	}
}

// WithCustomer tells the query-builder to eager-load the nodes that are connected to
// the "customer" edge. The optional arguments are used to configure the query builder of the edge.
func (csq *CustomerSettingQuery) WithCustomer(opts ...func(*CustomerQuery)) *CustomerSettingQuery {
	query := (&CustomerClient{config: csq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	csq.withCustomer = query
	return csq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CustomerSetting.Query().
//		GroupBy(customersetting.FieldCreatedAt).
//		Aggregate(entities.Count()).
//		Scan(ctx, &v)
func (csq *CustomerSettingQuery) GroupBy(field string, fields ...string) *CustomerSettingGroupBy {
	csq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CustomerSettingGroupBy{build: csq}
	grbuild.flds = &csq.ctx.Fields
	grbuild.label = customersetting.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.CustomerSetting.Query().
//		Select(customersetting.FieldCreatedAt).
//		Scan(ctx, &v)
func (csq *CustomerSettingQuery) Select(fields ...string) *CustomerSettingSelect {
	csq.ctx.Fields = append(csq.ctx.Fields, fields...)
	sbuild := &CustomerSettingSelect{CustomerSettingQuery: csq}
	sbuild.label = customersetting.Label
	sbuild.flds, sbuild.scan = &csq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CustomerSettingSelect configured with the given aggregations.
func (csq *CustomerSettingQuery) Aggregate(fns ...AggregateFunc) *CustomerSettingSelect {
	return csq.Select().Aggregate(fns...)
}

func (csq *CustomerSettingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range csq.inters {
		if inter == nil {
			return fmt.Errorf("entities: uninitialized interceptor (forgotten import entities/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, csq); err != nil {
				return err
			}
		}
	}
	for _, f := range csq.ctx.Fields {
		if !customersetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entities: invalid field %q for query", f)}
		}
	}
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *CustomerSettingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CustomerSetting, error) {
	var (
		nodes       = []*CustomerSetting{}
		withFKs     = csq.withFKs
		_spec       = csq.querySpec()
		loadedTypes = [1]bool{
			csq.withCustomer != nil,
		}
	)
	if csq.withCustomer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, customersetting.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CustomerSetting).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CustomerSetting{config: csq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = csq.schemaConfig.CustomerSetting
	ctx = internal.NewSchemaConfigContext(ctx, csq.schemaConfig)
	if len(csq.modifiers) > 0 {
		_spec.Modifiers = csq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := csq.withCustomer; query != nil {
		if err := csq.loadCustomer(ctx, query, nodes, nil,
			func(n *CustomerSetting, e *Customer) { n.Edges.Customer = e }); err != nil {
			return nil, err
		}
	}
	for i := range csq.loadTotal {
		if err := csq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (csq *CustomerSettingQuery) loadCustomer(ctx context.Context, query *CustomerQuery, nodes []*CustomerSetting, init func(*CustomerSetting), assign func(*CustomerSetting, *Customer)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*CustomerSetting)
	for i := range nodes {
		if nodes[i].customer_customer_settings == nil {
			continue
		}
		fk := *nodes[i].customer_customer_settings
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(customer.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "customer_customer_settings" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (csq *CustomerSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	_spec.Node.Schema = csq.schemaConfig.CustomerSetting
	ctx = internal.NewSchemaConfigContext(ctx, csq.schemaConfig)
	if len(csq.modifiers) > 0 {
		_spec.Modifiers = csq.modifiers
	}
	_spec.Node.Columns = csq.ctx.Fields
	if len(csq.ctx.Fields) > 0 {
		_spec.Unique = csq.ctx.Unique != nil && *csq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *CustomerSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(customersetting.Table, customersetting.Columns, sqlgraph.NewFieldSpec(customersetting.FieldID, field.TypeString))
	_spec.From = csq.sql
	if unique := csq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if csq.path != nil {
		_spec.Unique = true
	}
	if fields := csq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customersetting.FieldID)
		for i := range fields {
			if fields[i] != customersetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csq *CustomerSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(customersetting.Table)
	columns := csq.ctx.Fields
	if len(columns) == 0 {
		columns = customersetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csq.ctx.Unique != nil && *csq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(csq.schemaConfig.CustomerSetting)
	ctx = internal.NewSchemaConfigContext(ctx, csq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range csq.modifiers {
		m(selector)
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (csq *CustomerSettingQuery) ForUpdate(opts ...sql.LockOption) *CustomerSettingQuery {
	if csq.driver.Dialect() == dialect.Postgres {
		csq.Unique(false)
	}
	csq.modifiers = append(csq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return csq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (csq *CustomerSettingQuery) ForShare(opts ...sql.LockOption) *CustomerSettingQuery {
	if csq.driver.Dialect() == dialect.Postgres {
		csq.Unique(false)
	}
	csq.modifiers = append(csq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return csq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (csq *CustomerSettingQuery) Modify(modifiers ...func(s *sql.Selector)) *CustomerSettingSelect {
	csq.modifiers = append(csq.modifiers, modifiers...)
	return csq.Select()
}

// CustomerSettingGroupBy is the group-by builder for CustomerSetting entities.
type CustomerSettingGroupBy struct {
	selector
	build *CustomerSettingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *CustomerSettingGroupBy) Aggregate(fns ...AggregateFunc) *CustomerSettingGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the selector query and scans the result into the given value.
func (csgb *CustomerSettingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, csgb.build.ctx, "GroupBy")
	if err := csgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerSettingQuery, *CustomerSettingGroupBy](ctx, csgb.build, csgb, csgb.build.inters, v)
}

func (csgb *CustomerSettingGroupBy) sqlScan(ctx context.Context, root *CustomerSettingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(csgb.fns))
	for _, fn := range csgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*csgb.flds)+len(csgb.fns))
		for _, f := range *csgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*csgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CustomerSettingSelect is the builder for selecting fields of CustomerSetting entities.
type CustomerSettingSelect struct {
	*CustomerSettingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (css *CustomerSettingSelect) Aggregate(fns ...AggregateFunc) *CustomerSettingSelect {
	css.fns = append(css.fns, fns...)
	return css
}

// Scan applies the selector query and scans the result into the given value.
func (css *CustomerSettingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, css.ctx, "Select")
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerSettingQuery, *CustomerSettingSelect](ctx, css.CustomerSettingQuery, css, css.inters, v)
}

func (css *CustomerSettingSelect) sqlScan(ctx context.Context, root *CustomerSettingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(css.fns))
	for _, fn := range css.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*css.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (css *CustomerSettingSelect) Modify(modifiers ...func(s *sql.Selector)) *CustomerSettingSelect {
	css.modifiers = append(css.modifiers, modifiers...)
	return css
}
