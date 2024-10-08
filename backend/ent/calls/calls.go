// Code generated by ent, DO NOT EDIT.

package calls

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the calls type in the database.
	Label = "calls"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSessionID holds the string denoting the session_id field in the database.
	FieldSessionID = "session_id"
	// FieldCallStart holds the string denoting the call_start field in the database.
	FieldCallStart = "call_start"
	// FieldCallEnd holds the string denoting the call_end field in the database.
	FieldCallEnd = "call_end"
	// FieldRating holds the string denoting the rating field in the database.
	FieldRating = "rating"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeMade holds the string denoting the made edge name in mutations.
	EdgeMade = "made"
	// Table holds the table name of the calls in the database.
	Table = "call_ss"
	// MadeTable is the table that holds the made relation/edge.
	MadeTable = "call_ss"
	// MadeInverseTable is the table name for the SESSIONS entity.
	// It exists in this package in order to avoid circular dependency with the "sessions" package.
	MadeInverseTable = "session_ss"
	// MadeColumn is the table column denoting the made relation/edge.
	MadeColumn = "sessions_makes"
)

// Columns holds all SQL columns for calls fields.
var Columns = []string{
	FieldID,
	FieldSessionID,
	FieldCallStart,
	FieldCallEnd,
	FieldRating,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "call_ss"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"sessions_makes",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCallStart holds the default value on creation for the "call_start" field.
	DefaultCallStart func() time.Time
	// DefaultCallEnd holds the default value on creation for the "call_end" field.
	DefaultCallEnd func() time.Time
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the CALLS queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySessionID orders the results by the session_id field.
func BySessionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSessionID, opts...).ToFunc()
}

// ByCallStart orders the results by the call_start field.
func ByCallStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCallStart, opts...).ToFunc()
}

// ByCallEnd orders the results by the call_end field.
func ByCallEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCallEnd, opts...).ToFunc()
}

// ByRating orders the results by the rating field.
func ByRating(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRating, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByMadeField orders the results by made field.
func ByMadeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMadeStep(), sql.OrderByField(field, opts...))
	}
}
func newMadeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MadeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, MadeTable, MadeColumn),
	)
}
