// Code generated by ent, DO NOT EDIT.

package entities

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomer"
	"github.com/sigmasee/sigmasee/apex/shared/entities/apexcustomeridentity"
)

// ApexCustomerIdentity is the model entity for the ApexCustomerIdentity schema.
type ApexCustomerIdentity struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// EmailVerified holds the value of the "email_verified" field.
	EmailVerified bool `json:"email_verified,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ApexCustomerIdentityQuery when eager-loading is set.
	Edges                    ApexCustomerIdentityEdges `json:"edges"`
	apex_customer_identities *string
	selectValues             sql.SelectValues
}

// ApexCustomerIdentityEdges holds the relations/edges for other nodes in the graph.
type ApexCustomerIdentityEdges struct {
	// Customer holds the value of the customer edge.
	Customer *ApexCustomer `json:"customer,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ApexCustomerIdentityEdges) CustomerOrErr() (*ApexCustomer, error) {
	if e.loadedTypes[0] {
		if e.Customer == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: apexcustomer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApexCustomerIdentity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case apexcustomeridentity.FieldEmailVerified:
			values[i] = new(sql.NullBool)
		case apexcustomeridentity.FieldID, apexcustomeridentity.FieldEmail:
			values[i] = new(sql.NullString)
		case apexcustomeridentity.FieldCreatedAt, apexcustomeridentity.FieldModifiedAt, apexcustomeridentity.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case apexcustomeridentity.ForeignKeys[0]: // apex_customer_identities
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApexCustomerIdentity fields.
func (aci *ApexCustomerIdentity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case apexcustomeridentity.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				aci.ID = value.String
			}
		case apexcustomeridentity.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				aci.Email = value.String
			}
		case apexcustomeridentity.FieldEmailVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_verified", values[i])
			} else if value.Valid {
				aci.EmailVerified = value.Bool
			}
		case apexcustomeridentity.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				aci.CreatedAt = value.Time
			}
		case apexcustomeridentity.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				aci.ModifiedAt = value.Time
			}
		case apexcustomeridentity.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				aci.DeletedAt = value.Time
			}
		case apexcustomeridentity.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field apex_customer_identities", values[i])
			} else if value.Valid {
				aci.apex_customer_identities = new(string)
				*aci.apex_customer_identities = value.String
			}
		default:
			aci.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ApexCustomerIdentity.
// This includes values selected through modifiers, order, etc.
func (aci *ApexCustomerIdentity) Value(name string) (ent.Value, error) {
	return aci.selectValues.Get(name)
}

// QueryCustomer queries the "customer" edge of the ApexCustomerIdentity entity.
func (aci *ApexCustomerIdentity) QueryCustomer() *ApexCustomerQuery {
	return NewApexCustomerIdentityClient(aci.config).QueryCustomer(aci)
}

// Update returns a builder for updating this ApexCustomerIdentity.
// Note that you need to call ApexCustomerIdentity.Unwrap() before calling this method if this ApexCustomerIdentity
// was returned from a transaction, and the transaction was committed or rolled back.
func (aci *ApexCustomerIdentity) Update() *ApexCustomerIdentityUpdateOne {
	return NewApexCustomerIdentityClient(aci.config).UpdateOne(aci)
}

// Unwrap unwraps the ApexCustomerIdentity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (aci *ApexCustomerIdentity) Unwrap() *ApexCustomerIdentity {
	_tx, ok := aci.config.driver.(*txDriver)
	if !ok {
		panic("entities: ApexCustomerIdentity is not a transactional entity")
	}
	aci.config.driver = _tx.drv
	return aci
}

// String implements the fmt.Stringer.
func (aci *ApexCustomerIdentity) String() string {
	var builder strings.Builder
	builder.WriteString("ApexCustomerIdentity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", aci.ID))
	builder.WriteString("email=")
	builder.WriteString(aci.Email)
	builder.WriteString(", ")
	builder.WriteString("email_verified=")
	builder.WriteString(fmt.Sprintf("%v", aci.EmailVerified))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(aci.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(aci.ModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(aci.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ApexCustomerIdentities is a parsable slice of ApexCustomerIdentity.
type ApexCustomerIdentities []*ApexCustomerIdentity
