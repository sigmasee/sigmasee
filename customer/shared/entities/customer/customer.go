// Code generated by ent, DO NOT EDIT.

package customer

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDesignation holds the string denoting the designation field in the database.
	FieldDesignation = "designation"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGivenName holds the string denoting the given_name field in the database.
	FieldGivenName = "given_name"
	// FieldMiddleName holds the string denoting the middle_name field in the database.
	FieldMiddleName = "middle_name"
	// FieldFamilyName holds the string denoting the family_name field in the database.
	FieldFamilyName = "family_name"
	// FieldPhotoURL holds the string denoting the photo_url field in the database.
	FieldPhotoURL = "photo_url"
	// FieldPhotoURL24 holds the string denoting the photo_url_24 field in the database.
	FieldPhotoURL24 = "photo_url_24"
	// FieldPhotoURL32 holds the string denoting the photo_url_32 field in the database.
	FieldPhotoURL32 = "photo_url_32"
	// FieldPhotoURL48 holds the string denoting the photo_url_48 field in the database.
	FieldPhotoURL48 = "photo_url_48"
	// FieldPhotoURL72 holds the string denoting the photo_url_72 field in the database.
	FieldPhotoURL72 = "photo_url_72"
	// FieldPhotoURL192 holds the string denoting the photo_url_192 field in the database.
	FieldPhotoURL192 = "photo_url_192"
	// FieldPhotoURL512 holds the string denoting the photo_url_512 field in the database.
	FieldPhotoURL512 = "photo_url_512"
	// FieldTimezone holds the string denoting the timezone field in the database.
	FieldTimezone = "timezone"
	// FieldLocale holds the string denoting the locale field in the database.
	FieldLocale = "locale"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeIdentities holds the string denoting the identities edge name in mutations.
	EdgeIdentities = "identities"
	// EdgeCustomerSettings holds the string denoting the customer_settings edge name in mutations.
	EdgeCustomerSettings = "customer_settings"
	// Table holds the table name of the customer in the database.
	Table = "customers"
	// IdentitiesTable is the table that holds the identities relation/edge.
	IdentitiesTable = "identities"
	// IdentitiesInverseTable is the table name for the Identity entity.
	// It exists in this package in order to avoid circular dependency with the "identity" package.
	IdentitiesInverseTable = "identities"
	// IdentitiesColumn is the table column denoting the identities relation/edge.
	IdentitiesColumn = "customer_identities"
	// CustomerSettingsTable is the table that holds the customer_settings relation/edge.
	CustomerSettingsTable = "customer_settings"
	// CustomerSettingsInverseTable is the table name for the CustomerSetting entity.
	// It exists in this package in order to avoid circular dependency with the "customersetting" package.
	CustomerSettingsInverseTable = "customer_settings"
	// CustomerSettingsColumn is the table column denoting the customer_settings relation/edge.
	CustomerSettingsColumn = "customer_customer_settings"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldDesignation,
	FieldTitle,
	FieldName,
	FieldGivenName,
	FieldMiddleName,
	FieldFamilyName,
	FieldPhotoURL,
	FieldPhotoURL24,
	FieldPhotoURL32,
	FieldPhotoURL48,
	FieldPhotoURL72,
	FieldPhotoURL192,
	FieldPhotoURL512,
	FieldTimezone,
	FieldLocale,
	FieldCreatedAt,
	FieldModifiedAt,
	FieldDeletedAt,
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

// OrderOption defines the ordering options for the Customer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByDesignation orders the results by the designation field.
func ByDesignation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesignation, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByGivenName orders the results by the given_name field.
func ByGivenName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGivenName, opts...).ToFunc()
}

// ByMiddleName orders the results by the middle_name field.
func ByMiddleName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMiddleName, opts...).ToFunc()
}

// ByFamilyName orders the results by the family_name field.
func ByFamilyName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFamilyName, opts...).ToFunc()
}

// ByPhotoURL orders the results by the photo_url field.
func ByPhotoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL, opts...).ToFunc()
}

// ByPhotoURL24 orders the results by the photo_url_24 field.
func ByPhotoURL24(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL24, opts...).ToFunc()
}

// ByPhotoURL32 orders the results by the photo_url_32 field.
func ByPhotoURL32(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL32, opts...).ToFunc()
}

// ByPhotoURL48 orders the results by the photo_url_48 field.
func ByPhotoURL48(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL48, opts...).ToFunc()
}

// ByPhotoURL72 orders the results by the photo_url_72 field.
func ByPhotoURL72(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL72, opts...).ToFunc()
}

// ByPhotoURL192 orders the results by the photo_url_192 field.
func ByPhotoURL192(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL192, opts...).ToFunc()
}

// ByPhotoURL512 orders the results by the photo_url_512 field.
func ByPhotoURL512(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL512, opts...).ToFunc()
}

// ByTimezone orders the results by the timezone field.
func ByTimezone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTimezone, opts...).ToFunc()
}

// ByLocale orders the results by the locale field.
func ByLocale(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocale, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByModifiedAt orders the results by the modified_at field.
func ByModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByIdentitiesCount orders the results by identities count.
func ByIdentitiesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newIdentitiesStep(), opts...)
	}
}

// ByIdentities orders the results by identities terms.
func ByIdentities(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIdentitiesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCustomerSettingsField orders the results by customer_settings field.
func ByCustomerSettingsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerSettingsStep(), sql.OrderByField(field, opts...))
	}
}
func newIdentitiesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IdentitiesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, IdentitiesTable, IdentitiesColumn),
	)
}
func newCustomerSettingsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerSettingsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CustomerSettingsTable, CustomerSettingsColumn),
	)
}
