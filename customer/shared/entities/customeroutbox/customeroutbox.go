// Code generated by ent, DO NOT EDIT.

package customeroutbox

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the customeroutbox type in the database.
	Label = "customer_outbox"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldTopic holds the string denoting the topic field in the database.
	FieldTopic = "topic"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldPayload holds the string denoting the payload field in the database.
	FieldPayload = "payload"
	// FieldHeaders holds the string denoting the headers field in the database.
	FieldHeaders = "headers"
	// FieldRetryCount holds the string denoting the retry_count field in the database.
	FieldRetryCount = "retry_count"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLastRetry holds the string denoting the last_retry field in the database.
	FieldLastRetry = "last_retry"
	// FieldProcessingErrors holds the string denoting the processing_errors field in the database.
	FieldProcessingErrors = "processing_errors"
	// Table holds the table name of the customeroutbox in the database.
	Table = "customer_outboxes"
)

// Columns holds all SQL columns for customeroutbox fields.
var Columns = []string{
	FieldID,
	FieldTimestamp,
	FieldTopic,
	FieldKey,
	FieldPayload,
	FieldHeaders,
	FieldRetryCount,
	FieldStatus,
	FieldLastRetry,
	FieldProcessingErrors,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusPENDING Status = "PENDING"
	StatusFAILED  Status = "FAILED"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPENDING, StatusFAILED:
		return nil
	default:
		return fmt.Errorf("customeroutbox: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the CustomerOutbox queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTimestamp orders the results by the timestamp field.
func ByTimestamp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimestamp, opts...).ToFunc()
}

// ByTopic orders the results by the topic field.
func ByTopic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTopic, opts...).ToFunc()
}

// ByRetryCount orders the results by the retry_count field.
func ByRetryCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRetryCount, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLastRetry orders the results by the last_retry field.
func ByLastRetry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastRetry, opts...).ToFunc()
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Status) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Status) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Status(str)
	if err := StatusValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}