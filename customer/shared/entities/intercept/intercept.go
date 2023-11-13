// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/sigmasee/sigmasee/customer/shared/entities"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customer"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customeroutbox"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customersetting"
	"github.com/sigmasee/sigmasee/customer/shared/entities/identity"
	"github.com/sigmasee/sigmasee/customer/shared/entities/predicate"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next entities.Querier) entities.Querier {
	return entities.QuerierFunc(func(ctx context.Context, q entities.Query) (entities.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next entities.Querier) entities.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q entities.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The CustomerFunc type is an adapter to allow the use of ordinary function as a Querier.
type CustomerFunc func(context.Context, *entities.CustomerQuery) (entities.Value, error)

// Query calls f(ctx, q).
func (f CustomerFunc) Query(ctx context.Context, q entities.Query) (entities.Value, error) {
	if q, ok := q.(*entities.CustomerQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *entities.CustomerQuery", q)
}

// The TraverseCustomer type is an adapter to allow the use of ordinary function as Traverser.
type TraverseCustomer func(context.Context, *entities.CustomerQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseCustomer) Intercept(next entities.Querier) entities.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseCustomer) Traverse(ctx context.Context, q entities.Query) error {
	if q, ok := q.(*entities.CustomerQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *entities.CustomerQuery", q)
}

// The CustomerOutboxFunc type is an adapter to allow the use of ordinary function as a Querier.
type CustomerOutboxFunc func(context.Context, *entities.CustomerOutboxQuery) (entities.Value, error)

// Query calls f(ctx, q).
func (f CustomerOutboxFunc) Query(ctx context.Context, q entities.Query) (entities.Value, error) {
	if q, ok := q.(*entities.CustomerOutboxQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *entities.CustomerOutboxQuery", q)
}

// The TraverseCustomerOutbox type is an adapter to allow the use of ordinary function as Traverser.
type TraverseCustomerOutbox func(context.Context, *entities.CustomerOutboxQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseCustomerOutbox) Intercept(next entities.Querier) entities.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseCustomerOutbox) Traverse(ctx context.Context, q entities.Query) error {
	if q, ok := q.(*entities.CustomerOutboxQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *entities.CustomerOutboxQuery", q)
}

// The CustomerSettingFunc type is an adapter to allow the use of ordinary function as a Querier.
type CustomerSettingFunc func(context.Context, *entities.CustomerSettingQuery) (entities.Value, error)

// Query calls f(ctx, q).
func (f CustomerSettingFunc) Query(ctx context.Context, q entities.Query) (entities.Value, error) {
	if q, ok := q.(*entities.CustomerSettingQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *entities.CustomerSettingQuery", q)
}

// The TraverseCustomerSetting type is an adapter to allow the use of ordinary function as Traverser.
type TraverseCustomerSetting func(context.Context, *entities.CustomerSettingQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseCustomerSetting) Intercept(next entities.Querier) entities.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseCustomerSetting) Traverse(ctx context.Context, q entities.Query) error {
	if q, ok := q.(*entities.CustomerSettingQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *entities.CustomerSettingQuery", q)
}

// The IdentityFunc type is an adapter to allow the use of ordinary function as a Querier.
type IdentityFunc func(context.Context, *entities.IdentityQuery) (entities.Value, error)

// Query calls f(ctx, q).
func (f IdentityFunc) Query(ctx context.Context, q entities.Query) (entities.Value, error) {
	if q, ok := q.(*entities.IdentityQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *entities.IdentityQuery", q)
}

// The TraverseIdentity type is an adapter to allow the use of ordinary function as Traverser.
type TraverseIdentity func(context.Context, *entities.IdentityQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseIdentity) Intercept(next entities.Querier) entities.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseIdentity) Traverse(ctx context.Context, q entities.Query) error {
	if q, ok := q.(*entities.IdentityQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *entities.IdentityQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q entities.Query) (Query, error) {
	switch q := q.(type) {
	case *entities.CustomerQuery:
		return &query[*entities.CustomerQuery, predicate.Customer, customer.OrderOption]{typ: entities.TypeCustomer, tq: q}, nil
	case *entities.CustomerOutboxQuery:
		return &query[*entities.CustomerOutboxQuery, predicate.CustomerOutbox, customeroutbox.OrderOption]{typ: entities.TypeCustomerOutbox, tq: q}, nil
	case *entities.CustomerSettingQuery:
		return &query[*entities.CustomerSettingQuery, predicate.CustomerSetting, customersetting.OrderOption]{typ: entities.TypeCustomerSetting, tq: q}, nil
	case *entities.IdentityQuery:
		return &query[*entities.IdentityQuery, predicate.Identity, identity.OrderOption]{typ: entities.TypeIdentity, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}