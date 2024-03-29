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

// CustomerSetting is the model entity for the CustomerSetting schema.
type CustomerSetting struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerSettingQuery when eager-loading is set.
	Edges                      CustomerSettingEdges `json:"edges"`
	customer_customer_settings *string
	selectValues               sql.SelectValues
}

// CustomerSettingEdges holds the relations/edges for other nodes in the graph.
type CustomerSettingEdges struct {
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerSettingEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[0] {
		if e.Customer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CustomerSetting) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customersetting.FieldID:
			values[i] = new(sql.NullString)
		case customersetting.FieldCreatedAt, customersetting.FieldModifiedAt, customersetting.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case customersetting.ForeignKeys[0]: // customer_customer_settings
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CustomerSetting fields.
func (cs *CustomerSetting) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customersetting.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				cs.ID = value.String
			}
		case customersetting.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cs.CreatedAt = value.Time
			}
		case customersetting.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				cs.ModifiedAt = value.Time
			}
		case customersetting.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				cs.DeletedAt = value.Time
			}
		case customersetting.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field customer_customer_settings", values[i])
			} else if value.Valid {
				cs.customer_customer_settings = new(string)
				*cs.customer_customer_settings = value.String
			}
		default:
			cs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CustomerSetting.
// This includes values selected through modifiers, order, etc.
func (cs *CustomerSetting) Value(name string) (ent.Value, error) {
	return cs.selectValues.Get(name)
}

// QueryCustomer queries the "customer" edge of the CustomerSetting entity.
func (cs *CustomerSetting) QueryCustomer() *CustomerQuery {
	return NewCustomerSettingClient(cs.config).QueryCustomer(cs)
}

// Update returns a builder for updating this CustomerSetting.
// Note that you need to call CustomerSetting.Unwrap() before calling this method if this CustomerSetting
// was returned from a transaction, and the transaction was committed or rolled back.
func (cs *CustomerSetting) Update() *CustomerSettingUpdateOne {
	return NewCustomerSettingClient(cs.config).UpdateOne(cs)
}

// Unwrap unwraps the CustomerSetting entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cs *CustomerSetting) Unwrap() *CustomerSetting {
	_tx, ok := cs.config.driver.(*txDriver)
	if !ok {
		panic("entities: CustomerSetting is not a transactional entity")
	}
	cs.config.driver = _tx.drv
	return cs
}

// String implements the fmt.Stringer.
func (cs *CustomerSetting) String() string {
	var builder strings.Builder
	builder.WriteString("CustomerSetting(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(cs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(cs.ModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(cs.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// CustomerSettings is a parsable slice of CustomerSetting.
type CustomerSettings []*CustomerSetting
