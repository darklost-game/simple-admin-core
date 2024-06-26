// Code generated by ent, DO NOT EDIT.

package logoperation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the logoperation type in the database.
	Label = "log_operation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldMethod holds the string denoting the method field in the database.
	FieldMethod = "method"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldHeaders holds the string denoting the headers field in the database.
	FieldHeaders = "headers"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldStatusCode holds the string denoting the status_code field in the database.
	FieldStatusCode = "status_code"
	// FieldReqTime holds the string denoting the req_time field in the database.
	FieldReqTime = "req_time"
	// FieldResTime holds the string denoting the res_time field in the database.
	FieldResTime = "res_time"
	// FieldCostTime holds the string denoting the cost_time field in the database.
	FieldCostTime = "cost_time"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the logoperation in the database.
	Table = "sys_log_operations"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "sys_log_operations"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "sys_users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "uuid"
)

// Columns holds all SQL columns for logoperation fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUUID,
	FieldMethod,
	FieldPath,
	FieldHeaders,
	FieldBody,
	FieldStatusCode,
	FieldReqTime,
	FieldResTime,
	FieldCostTime,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// HeadersValidator is a validator for the "headers" field. It is called by the builders before save.
	HeadersValidator func(string) error
	// DefaultReqTime holds the default value on creation for the "req_time" field.
	DefaultReqTime func() time.Time
	// DefaultResTime holds the default value on creation for the "res_time" field.
	DefaultResTime func() time.Time
)

// OrderOption defines the ordering options for the LogOperation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByMethod orders the results by the method field.
func ByMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMethod, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByHeaders orders the results by the headers field.
func ByHeaders(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHeaders, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByStatusCode orders the results by the status_code field.
func ByStatusCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatusCode, opts...).ToFunc()
}

// ByReqTime orders the results by the req_time field.
func ByReqTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReqTime, opts...).ToFunc()
}

// ByResTime orders the results by the res_time field.
func ByResTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResTime, opts...).ToFunc()
}

// ByCostTime orders the results by the cost_time field.
func ByCostTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCostTime, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
