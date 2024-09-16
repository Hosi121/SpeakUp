// Code generated by ent, DO NOT EDIT.

package achievements

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the achievements type in the database.
	Label = "achievements"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldTrophyID holds the string denoting the trophy_id field in the database.
	FieldTrophyID = "trophy_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeGranted holds the string denoting the granted edge name in mutations.
	EdgeGranted = "granted"
	// EdgeRefers holds the string denoting the refers edge name in mutations.
	EdgeRefers = "refers"
	// Table holds the table name of the achievements in the database.
	Table = "achievement_ss"
	// GrantedTable is the table that holds the granted relation/edge.
	GrantedTable = "achievement_ss"
	// GrantedInverseTable is the table name for the USERS entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	GrantedInverseTable = "user_ss"
	// GrantedColumn is the table column denoting the granted relation/edge.
	GrantedColumn = "users_acquires"
	// RefersTable is the table that holds the refers relation/edge.
	RefersTable = "achievement_ss"
	// RefersInverseTable is the table name for the TROPHIES entity.
	// It exists in this package in order to avoid circular dependency with the "trophies" package.
	RefersInverseTable = "trophie_ss"
	// RefersColumn is the table column denoting the refers relation/edge.
	RefersColumn = "achievements_refers"
)

// Columns holds all SQL columns for achievements fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldTrophyID,
	FieldCreatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "achievement_ss"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"achievements_refers",
	"users_acquires",
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the ACHIEVEMENTS queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByTrophyID orders the results by the trophy_id field.
func ByTrophyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTrophyID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByGrantedField orders the results by granted field.
func ByGrantedField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGrantedStep(), sql.OrderByField(field, opts...))
	}
}

// ByRefersField orders the results by refers field.
func ByRefersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRefersStep(), sql.OrderByField(field, opts...))
	}
}
func newGrantedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GrantedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GrantedTable, GrantedColumn),
	)
}
func newRefersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RefersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RefersTable, RefersColumn),
	)
}