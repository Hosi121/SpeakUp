// Code generated by ent, DO NOT EDIT.

package memos

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the memos type in the database.
	Label = "memos"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldMemo1 holds the string denoting the memo1 field in the database.
	FieldMemo1 = "memo1"
	// FieldMemo2 holds the string denoting the memo2 field in the database.
	FieldMemo2 = "memo2"
	// EdgePrepared holds the string denoting the prepared edge name in mutations.
	EdgePrepared = "prepared"
	// Table holds the table name of the memos in the database.
	Table = "memo_ss"
	// PreparedTable is the table that holds the prepared relation/edge.
	PreparedTable = "memo_ss"
	// PreparedInverseTable is the table name for the USERS entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	PreparedInverseTable = "user_ss"
	// PreparedColumn is the table column denoting the prepared relation/edge.
	PreparedColumn = "users_prepares"
)

// Columns holds all SQL columns for memos fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldMemo1,
	FieldMemo2,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "memo_ss"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"users_prepares",
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
	// DefaultMemo1 holds the default value on creation for the "memo1" field.
	DefaultMemo1 string
	// Memo1Validator is a validator for the "memo1" field. It is called by the builders before save.
	Memo1Validator func(string) error
	// DefaultMemo2 holds the default value on creation for the "memo2" field.
	DefaultMemo2 string
	// Memo2Validator is a validator for the "memo2" field. It is called by the builders before save.
	Memo2Validator func(string) error
)

// OrderOption defines the ordering options for the MEMOS queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByMemo1 orders the results by the memo1 field.
func ByMemo1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo1, opts...).ToFunc()
}

// ByMemo2 orders the results by the memo2 field.
func ByMemo2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemo2, opts...).ToFunc()
}

// ByPreparedField orders the results by prepared field.
func ByPreparedField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPreparedStep(), sql.OrderByField(field, opts...))
	}
}
func newPreparedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PreparedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, PreparedTable, PreparedColumn),
	)
}