// Code generated by ent, DO NOT EDIT.

package entities

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/sigmasee/sigmasee/customer/shared/entities/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customer"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customeroutbox"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customersetting"
	"github.com/sigmasee/sigmasee/customer/shared/entities/identity"

	stdsql "database/sql"

	"github.com/sigmasee/sigmasee/customer/shared/entities/internal"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// CustomerOutbox is the client for interacting with the CustomerOutbox builders.
	CustomerOutbox *CustomerOutboxClient
	// CustomerSetting is the client for interacting with the CustomerSetting builders.
	CustomerSetting *CustomerSettingClient
	// Identity is the client for interacting with the Identity builders.
	Identity *IdentityClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Customer = NewCustomerClient(c.config)
	c.CustomerOutbox = NewCustomerOutboxClient(c.config)
	c.CustomerSetting = NewCustomerSettingClient(c.config)
	c.Identity = NewIdentityClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
		// schemaConfig contains alternative names for all tables.
		schemaConfig SchemaConfig
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("entities: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("entities: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Customer:        NewCustomerClient(cfg),
		CustomerOutbox:  NewCustomerOutboxClient(cfg),
		CustomerSetting: NewCustomerSettingClient(cfg),
		Identity:        NewIdentityClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		Customer:        NewCustomerClient(cfg),
		CustomerOutbox:  NewCustomerOutboxClient(cfg),
		CustomerSetting: NewCustomerSettingClient(cfg),
		Identity:        NewIdentityClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Customer.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Customer.Use(hooks...)
	c.CustomerOutbox.Use(hooks...)
	c.CustomerSetting.Use(hooks...)
	c.Identity.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Customer.Intercept(interceptors...)
	c.CustomerOutbox.Intercept(interceptors...)
	c.CustomerSetting.Intercept(interceptors...)
	c.Identity.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CustomerMutation:
		return c.Customer.mutate(ctx, m)
	case *CustomerOutboxMutation:
		return c.CustomerOutbox.mutate(ctx, m)
	case *CustomerSettingMutation:
		return c.CustomerSetting.mutate(ctx, m)
	case *IdentityMutation:
		return c.Identity.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("entities: unknown mutation type %T", m)
	}
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `customer.Intercept(f(g(h())))`.
func (c *CustomerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Customer = append(c.inters.Customer, interceptors...)
}

// Create returns a builder for creating a Customer entity.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Customer entities.
func (c *CustomerClient) CreateBulk(builders ...*CustomerCreate) *CustomerCreateBulk {
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CustomerClient) MapCreateBulk(slice any, setFunc func(*CustomerCreate, int)) *CustomerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CustomerCreateBulk{err: fmt.Errorf("calling to CustomerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CustomerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id string) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerClient) DeleteOneID(id string) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Query returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCustomer},
		inters: c.Interceptors(),
	}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id string) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id string) *Customer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryIdentities queries the identities edge of a Customer.
func (c *CustomerClient) QueryIdentities(cu *Customer) *IdentityQuery {
	query := (&IdentityClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(identity.Table, identity.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.IdentitiesTable, customer.IdentitiesColumn),
		)
		schemaConfig := cu.schemaConfig
		step.To.Schema = schemaConfig.Identity
		step.Edge.Schema = schemaConfig.Identity
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCustomerSettings queries the customer_settings edge of a Customer.
func (c *CustomerClient) QueryCustomerSettings(cu *Customer) *CustomerSettingQuery {
	query := (&CustomerSettingClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(customersetting.Table, customersetting.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, customer.CustomerSettingsTable, customer.CustomerSettingsColumn),
		)
		schemaConfig := cu.schemaConfig
		step.To.Schema = schemaConfig.CustomerSetting
		step.Edge.Schema = schemaConfig.CustomerSetting
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// Interceptors returns the client interceptors.
func (c *CustomerClient) Interceptors() []Interceptor {
	return c.inters.Customer
}

func (c *CustomerClient) mutate(ctx context.Context, m *CustomerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entities: unknown Customer mutation op: %q", m.Op())
	}
}

// CustomerOutboxClient is a client for the CustomerOutbox schema.
type CustomerOutboxClient struct {
	config
}

// NewCustomerOutboxClient returns a client for the CustomerOutbox from the given config.
func NewCustomerOutboxClient(c config) *CustomerOutboxClient {
	return &CustomerOutboxClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customeroutbox.Hooks(f(g(h())))`.
func (c *CustomerOutboxClient) Use(hooks ...Hook) {
	c.hooks.CustomerOutbox = append(c.hooks.CustomerOutbox, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `customeroutbox.Intercept(f(g(h())))`.
func (c *CustomerOutboxClient) Intercept(interceptors ...Interceptor) {
	c.inters.CustomerOutbox = append(c.inters.CustomerOutbox, interceptors...)
}

// Create returns a builder for creating a CustomerOutbox entity.
func (c *CustomerOutboxClient) Create() *CustomerOutboxCreate {
	mutation := newCustomerOutboxMutation(c.config, OpCreate)
	return &CustomerOutboxCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CustomerOutbox entities.
func (c *CustomerOutboxClient) CreateBulk(builders ...*CustomerOutboxCreate) *CustomerOutboxCreateBulk {
	return &CustomerOutboxCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CustomerOutboxClient) MapCreateBulk(slice any, setFunc func(*CustomerOutboxCreate, int)) *CustomerOutboxCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CustomerOutboxCreateBulk{err: fmt.Errorf("calling to CustomerOutboxClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CustomerOutboxCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CustomerOutboxCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CustomerOutbox.
func (c *CustomerOutboxClient) Update() *CustomerOutboxUpdate {
	mutation := newCustomerOutboxMutation(c.config, OpUpdate)
	return &CustomerOutboxUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerOutboxClient) UpdateOne(co *CustomerOutbox) *CustomerOutboxUpdateOne {
	mutation := newCustomerOutboxMutation(c.config, OpUpdateOne, withCustomerOutbox(co))
	return &CustomerOutboxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerOutboxClient) UpdateOneID(id string) *CustomerOutboxUpdateOne {
	mutation := newCustomerOutboxMutation(c.config, OpUpdateOne, withCustomerOutboxID(id))
	return &CustomerOutboxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CustomerOutbox.
func (c *CustomerOutboxClient) Delete() *CustomerOutboxDelete {
	mutation := newCustomerOutboxMutation(c.config, OpDelete)
	return &CustomerOutboxDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerOutboxClient) DeleteOne(co *CustomerOutbox) *CustomerOutboxDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerOutboxClient) DeleteOneID(id string) *CustomerOutboxDeleteOne {
	builder := c.Delete().Where(customeroutbox.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerOutboxDeleteOne{builder}
}

// Query returns a query builder for CustomerOutbox.
func (c *CustomerOutboxClient) Query() *CustomerOutboxQuery {
	return &CustomerOutboxQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCustomerOutbox},
		inters: c.Interceptors(),
	}
}

// Get returns a CustomerOutbox entity by its id.
func (c *CustomerOutboxClient) Get(ctx context.Context, id string) (*CustomerOutbox, error) {
	return c.Query().Where(customeroutbox.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerOutboxClient) GetX(ctx context.Context, id string) *CustomerOutbox {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CustomerOutboxClient) Hooks() []Hook {
	return c.hooks.CustomerOutbox
}

// Interceptors returns the client interceptors.
func (c *CustomerOutboxClient) Interceptors() []Interceptor {
	return c.inters.CustomerOutbox
}

func (c *CustomerOutboxClient) mutate(ctx context.Context, m *CustomerOutboxMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CustomerOutboxCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CustomerOutboxUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CustomerOutboxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CustomerOutboxDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entities: unknown CustomerOutbox mutation op: %q", m.Op())
	}
}

// CustomerSettingClient is a client for the CustomerSetting schema.
type CustomerSettingClient struct {
	config
}

// NewCustomerSettingClient returns a client for the CustomerSetting from the given config.
func NewCustomerSettingClient(c config) *CustomerSettingClient {
	return &CustomerSettingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customersetting.Hooks(f(g(h())))`.
func (c *CustomerSettingClient) Use(hooks ...Hook) {
	c.hooks.CustomerSetting = append(c.hooks.CustomerSetting, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `customersetting.Intercept(f(g(h())))`.
func (c *CustomerSettingClient) Intercept(interceptors ...Interceptor) {
	c.inters.CustomerSetting = append(c.inters.CustomerSetting, interceptors...)
}

// Create returns a builder for creating a CustomerSetting entity.
func (c *CustomerSettingClient) Create() *CustomerSettingCreate {
	mutation := newCustomerSettingMutation(c.config, OpCreate)
	return &CustomerSettingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CustomerSetting entities.
func (c *CustomerSettingClient) CreateBulk(builders ...*CustomerSettingCreate) *CustomerSettingCreateBulk {
	return &CustomerSettingCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CustomerSettingClient) MapCreateBulk(slice any, setFunc func(*CustomerSettingCreate, int)) *CustomerSettingCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CustomerSettingCreateBulk{err: fmt.Errorf("calling to CustomerSettingClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CustomerSettingCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CustomerSettingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CustomerSetting.
func (c *CustomerSettingClient) Update() *CustomerSettingUpdate {
	mutation := newCustomerSettingMutation(c.config, OpUpdate)
	return &CustomerSettingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerSettingClient) UpdateOne(cs *CustomerSetting) *CustomerSettingUpdateOne {
	mutation := newCustomerSettingMutation(c.config, OpUpdateOne, withCustomerSetting(cs))
	return &CustomerSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerSettingClient) UpdateOneID(id string) *CustomerSettingUpdateOne {
	mutation := newCustomerSettingMutation(c.config, OpUpdateOne, withCustomerSettingID(id))
	return &CustomerSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CustomerSetting.
func (c *CustomerSettingClient) Delete() *CustomerSettingDelete {
	mutation := newCustomerSettingMutation(c.config, OpDelete)
	return &CustomerSettingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerSettingClient) DeleteOne(cs *CustomerSetting) *CustomerSettingDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerSettingClient) DeleteOneID(id string) *CustomerSettingDeleteOne {
	builder := c.Delete().Where(customersetting.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerSettingDeleteOne{builder}
}

// Query returns a query builder for CustomerSetting.
func (c *CustomerSettingClient) Query() *CustomerSettingQuery {
	return &CustomerSettingQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCustomerSetting},
		inters: c.Interceptors(),
	}
}

// Get returns a CustomerSetting entity by its id.
func (c *CustomerSettingClient) Get(ctx context.Context, id string) (*CustomerSetting, error) {
	return c.Query().Where(customersetting.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerSettingClient) GetX(ctx context.Context, id string) *CustomerSetting {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCustomer queries the customer edge of a CustomerSetting.
func (c *CustomerSettingClient) QueryCustomer(cs *CustomerSetting) *CustomerQuery {
	query := (&CustomerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customersetting.Table, customersetting.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, customersetting.CustomerTable, customersetting.CustomerColumn),
		)
		schemaConfig := cs.schemaConfig
		step.To.Schema = schemaConfig.Customer
		step.Edge.Schema = schemaConfig.CustomerSetting
		fromV = sqlgraph.Neighbors(cs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerSettingClient) Hooks() []Hook {
	return c.hooks.CustomerSetting
}

// Interceptors returns the client interceptors.
func (c *CustomerSettingClient) Interceptors() []Interceptor {
	return c.inters.CustomerSetting
}

func (c *CustomerSettingClient) mutate(ctx context.Context, m *CustomerSettingMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CustomerSettingCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CustomerSettingUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CustomerSettingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CustomerSettingDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entities: unknown CustomerSetting mutation op: %q", m.Op())
	}
}

// IdentityClient is a client for the Identity schema.
type IdentityClient struct {
	config
}

// NewIdentityClient returns a client for the Identity from the given config.
func NewIdentityClient(c config) *IdentityClient {
	return &IdentityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `identity.Hooks(f(g(h())))`.
func (c *IdentityClient) Use(hooks ...Hook) {
	c.hooks.Identity = append(c.hooks.Identity, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `identity.Intercept(f(g(h())))`.
func (c *IdentityClient) Intercept(interceptors ...Interceptor) {
	c.inters.Identity = append(c.inters.Identity, interceptors...)
}

// Create returns a builder for creating a Identity entity.
func (c *IdentityClient) Create() *IdentityCreate {
	mutation := newIdentityMutation(c.config, OpCreate)
	return &IdentityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Identity entities.
func (c *IdentityClient) CreateBulk(builders ...*IdentityCreate) *IdentityCreateBulk {
	return &IdentityCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *IdentityClient) MapCreateBulk(slice any, setFunc func(*IdentityCreate, int)) *IdentityCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &IdentityCreateBulk{err: fmt.Errorf("calling to IdentityClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*IdentityCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &IdentityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Identity.
func (c *IdentityClient) Update() *IdentityUpdate {
	mutation := newIdentityMutation(c.config, OpUpdate)
	return &IdentityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IdentityClient) UpdateOne(i *Identity) *IdentityUpdateOne {
	mutation := newIdentityMutation(c.config, OpUpdateOne, withIdentity(i))
	return &IdentityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IdentityClient) UpdateOneID(id string) *IdentityUpdateOne {
	mutation := newIdentityMutation(c.config, OpUpdateOne, withIdentityID(id))
	return &IdentityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Identity.
func (c *IdentityClient) Delete() *IdentityDelete {
	mutation := newIdentityMutation(c.config, OpDelete)
	return &IdentityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *IdentityClient) DeleteOne(i *Identity) *IdentityDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *IdentityClient) DeleteOneID(id string) *IdentityDeleteOne {
	builder := c.Delete().Where(identity.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IdentityDeleteOne{builder}
}

// Query returns a query builder for Identity.
func (c *IdentityClient) Query() *IdentityQuery {
	return &IdentityQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeIdentity},
		inters: c.Interceptors(),
	}
}

// Get returns a Identity entity by its id.
func (c *IdentityClient) Get(ctx context.Context, id string) (*Identity, error) {
	return c.Query().Where(identity.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IdentityClient) GetX(ctx context.Context, id string) *Identity {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCustomer queries the customer edge of a Identity.
func (c *IdentityClient) QueryCustomer(i *Identity) *CustomerQuery {
	query := (&CustomerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(identity.Table, identity.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, identity.CustomerTable, identity.CustomerColumn),
		)
		schemaConfig := i.schemaConfig
		step.To.Schema = schemaConfig.Customer
		step.Edge.Schema = schemaConfig.Identity
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IdentityClient) Hooks() []Hook {
	return c.hooks.Identity
}

// Interceptors returns the client interceptors.
func (c *IdentityClient) Interceptors() []Interceptor {
	return c.inters.Identity
}

func (c *IdentityClient) mutate(ctx context.Context, m *IdentityMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&IdentityCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&IdentityUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&IdentityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&IdentityDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("entities: unknown Identity mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Customer, CustomerOutbox, CustomerSetting, Identity []ent.Hook
	}
	inters struct {
		Customer, CustomerOutbox, CustomerSetting, Identity []ent.Interceptor
	}
)

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig = internal.SchemaConfig

// AlternateSchemas allows alternate schema names to be
// passed into ent operations.
func AlternateSchema(schemaConfig SchemaConfig) Option {
	return func(c *config) {
		c.schemaConfig = schemaConfig
	}
}

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
