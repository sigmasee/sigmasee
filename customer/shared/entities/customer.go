// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customer"
	"github.com/sigmasee/sigmasee/customer/shared/entities/customersetting"
)

// Customer is the model entity for the Customer schema.
type Customer struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Designation holds the value of the "designation" field.
	Designation string `json:"designation,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// GivenName holds the value of the "given_name" field.
	GivenName string `json:"given_name,omitempty"`
	// MiddleName holds the value of the "middle_name" field.
	MiddleName string `json:"middle_name,omitempty"`
	// FamilyName holds the value of the "family_name" field.
	FamilyName string `json:"family_name,omitempty"`
	// PhotoURL holds the value of the "photo_url" field.
	PhotoURL string `json:"photo_url,omitempty"`
	// PhotoURL24 holds the value of the "photo_url_24" field.
	PhotoURL24 string `json:"photo_url_24,omitempty"`
	// PhotoURL32 holds the value of the "photo_url_32" field.
	PhotoURL32 string `json:"photo_url_32,omitempty"`
	// PhotoURL48 holds the value of the "photo_url_48" field.
	PhotoURL48 string `json:"photo_url_48,omitempty"`
	// PhotoURL72 holds the value of the "photo_url_72" field.
	PhotoURL72 string `json:"photo_url_72,omitempty"`
	// PhotoURL192 holds the value of the "photo_url_192" field.
	PhotoURL192 string `json:"photo_url_192,omitempty"`
	// PhotoURL512 holds the value of the "photo_url_512" field.
	PhotoURL512 string `json:"photo_url_512,omitempty"`
	// Timezone holds the value of the "timezone" field.
	Timezone string `json:"timezone,omitempty"`
	// Locale holds the value of the "locale" field.
	Locale string `json:"locale,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerQuery when eager-loading is set.
	Edges        CustomerEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CustomerEdges holds the relations/edges for other nodes in the graph.
type CustomerEdges struct {
	// Identities holds the value of the identities edge.
	Identities []*Identity `json:"identities,omitempty"`
	// CustomerSettings holds the value of the customer_settings edge.
	CustomerSettings *CustomerSetting `json:"customer_settings,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedIdentities map[string][]*Identity
}

// IdentitiesOrErr returns the Identities value or an error if the edge
// was not loaded in eager-loading.
func (e CustomerEdges) IdentitiesOrErr() ([]*Identity, error) {
	if e.loadedTypes[0] {
		return e.Identities, nil
	}
	return nil, &NotLoadedError{edge: "identities"}
}

// CustomerSettingsOrErr returns the CustomerSettings value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerEdges) CustomerSettingsOrErr() (*CustomerSetting, error) {
	if e.loadedTypes[1] {
		if e.CustomerSettings == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customersetting.Label}
		}
		return e.CustomerSettings, nil
	}
	return nil, &NotLoadedError{edge: "customer_settings"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Customer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customer.FieldID, customer.FieldDesignation, customer.FieldTitle, customer.FieldName, customer.FieldGivenName, customer.FieldMiddleName, customer.FieldFamilyName, customer.FieldPhotoURL, customer.FieldPhotoURL24, customer.FieldPhotoURL32, customer.FieldPhotoURL48, customer.FieldPhotoURL72, customer.FieldPhotoURL192, customer.FieldPhotoURL512, customer.FieldTimezone, customer.FieldLocale:
			values[i] = new(sql.NullString)
		case customer.FieldCreatedAt, customer.FieldModifiedAt, customer.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Customer fields.
func (c *Customer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customer.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case customer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case customer.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				c.ModifiedAt = value.Time
			}
		case customer.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case customer.FieldDesignation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field designation", values[i])
			} else if value.Valid {
				c.Designation = value.String
			}
		case customer.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case customer.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case customer.FieldGivenName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field given_name", values[i])
			} else if value.Valid {
				c.GivenName = value.String
			}
		case customer.FieldMiddleName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field middle_name", values[i])
			} else if value.Valid {
				c.MiddleName = value.String
			}
		case customer.FieldFamilyName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field family_name", values[i])
			} else if value.Valid {
				c.FamilyName = value.String
			}
		case customer.FieldPhotoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url", values[i])
			} else if value.Valid {
				c.PhotoURL = value.String
			}
		case customer.FieldPhotoURL24:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url_24", values[i])
			} else if value.Valid {
				c.PhotoURL24 = value.String
			}
		case customer.FieldPhotoURL32:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url_32", values[i])
			} else if value.Valid {
				c.PhotoURL32 = value.String
			}
		case customer.FieldPhotoURL48:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url_48", values[i])
			} else if value.Valid {
				c.PhotoURL48 = value.String
			}
		case customer.FieldPhotoURL72:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url_72", values[i])
			} else if value.Valid {
				c.PhotoURL72 = value.String
			}
		case customer.FieldPhotoURL192:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url_192", values[i])
			} else if value.Valid {
				c.PhotoURL192 = value.String
			}
		case customer.FieldPhotoURL512:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url_512", values[i])
			} else if value.Valid {
				c.PhotoURL512 = value.String
			}
		case customer.FieldTimezone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field timezone", values[i])
			} else if value.Valid {
				c.Timezone = value.String
			}
		case customer.FieldLocale:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field locale", values[i])
			} else if value.Valid {
				c.Locale = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Customer.
// This includes values selected through modifiers, order, etc.
func (c *Customer) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryIdentities queries the "identities" edge of the Customer entity.
func (c *Customer) QueryIdentities() *IdentityQuery {
	return NewCustomerClient(c.config).QueryIdentities(c)
}

// QueryCustomerSettings queries the "customer_settings" edge of the Customer entity.
func (c *Customer) QueryCustomerSettings() *CustomerSettingQuery {
	return NewCustomerClient(c.config).QueryCustomerSettings(c)
}

// Update returns a builder for updating this Customer.
// Note that you need to call Customer.Unwrap() before calling this method if this Customer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Customer) Update() *CustomerUpdateOne {
	return NewCustomerClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Customer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Customer) Unwrap() *Customer {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("entities: Customer is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Customer) String() string {
	var builder strings.Builder
	builder.WriteString("Customer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(c.ModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("designation=")
	builder.WriteString(c.Designation)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(c.Title)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("given_name=")
	builder.WriteString(c.GivenName)
	builder.WriteString(", ")
	builder.WriteString("middle_name=")
	builder.WriteString(c.MiddleName)
	builder.WriteString(", ")
	builder.WriteString("family_name=")
	builder.WriteString(c.FamilyName)
	builder.WriteString(", ")
	builder.WriteString("photo_url=")
	builder.WriteString(c.PhotoURL)
	builder.WriteString(", ")
	builder.WriteString("photo_url_24=")
	builder.WriteString(c.PhotoURL24)
	builder.WriteString(", ")
	builder.WriteString("photo_url_32=")
	builder.WriteString(c.PhotoURL32)
	builder.WriteString(", ")
	builder.WriteString("photo_url_48=")
	builder.WriteString(c.PhotoURL48)
	builder.WriteString(", ")
	builder.WriteString("photo_url_72=")
	builder.WriteString(c.PhotoURL72)
	builder.WriteString(", ")
	builder.WriteString("photo_url_192=")
	builder.WriteString(c.PhotoURL192)
	builder.WriteString(", ")
	builder.WriteString("photo_url_512=")
	builder.WriteString(c.PhotoURL512)
	builder.WriteString(", ")
	builder.WriteString("timezone=")
	builder.WriteString(c.Timezone)
	builder.WriteString(", ")
	builder.WriteString("locale=")
	builder.WriteString(c.Locale)
	builder.WriteByte(')')
	return builder.String()
}

// NamedIdentities returns the Identities named value or an error if the edge was not
// loaded in eager-loading with this name.
func (c *Customer) NamedIdentities(name string) ([]*Identity, error) {
	if c.Edges.namedIdentities == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := c.Edges.namedIdentities[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (c *Customer) appendNamedIdentities(name string, edges ...*Identity) {
	if c.Edges.namedIdentities == nil {
		c.Edges.namedIdentities = make(map[string][]*Identity)
	}
	if len(edges) == 0 {
		c.Edges.namedIdentities[name] = []*Identity{}
	} else {
		c.Edges.namedIdentities[name] = append(c.Edges.namedIdentities[name], edges...)
	}
}

// Customers is a parsable slice of Customer.
type Customers []*Customer
