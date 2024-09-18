// Code generated by ent, DO NOT EDIT.

package event_records

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the event_records type in the database.
	Label = "event_records"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldEventID holds the string denoting the event_id field in the database.
	FieldEventID = "event_id"
	// FieldRecords holds the string denoting the records field in the database.
	FieldRecords = "records"
	// EdgeMade holds the string denoting the made edge name in mutations.
	EdgeMade = "made"
	// EdgeParticipates holds the string denoting the participates edge name in mutations.
	EdgeParticipates = "participates"
	// EdgeHas holds the string denoting the has edge name in mutations.
	EdgeHas = "has"
	// Table holds the table name of the event_records in the database.
	Table = "event_record_ss"
	// MadeTable is the table that holds the made relation/edge.
	MadeTable = "event_record_ss"
	// MadeInverseTable is the table name for the USERS entity.
	// It exists in this package in order to avoid circular dependency with the "users" package.
	MadeInverseTable = "user_ss"
	// MadeColumn is the table column denoting the made relation/edge.
	MadeColumn = "users_makes"
	// ParticipatesTable is the table that holds the participates relation/edge.
	ParticipatesTable = "event_record_ss"
	// ParticipatesInverseTable is the table name for the EVENTS entity.
	// It exists in this package in order to avoid circular dependency with the "events" package.
	ParticipatesInverseTable = "event_ss"
	// ParticipatesColumn is the table column denoting the participates relation/edge.
	ParticipatesColumn = "events_participated"
	// HasTable is the table that holds the has relation/edge.
	HasTable = "session_ss"
	// HasInverseTable is the table name for the SESSIONS entity.
	// It exists in this package in order to avoid circular dependency with the "sessions" package.
	HasInverseTable = "session_ss"
	// HasColumn is the table column denoting the has relation/edge.
	HasColumn = "event_records_has"
)

// Columns holds all SQL columns for event_records fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldEventID,
	FieldRecords,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "event_record_ss"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"events_participated",
	"users_makes",
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
	// DefaultRecords holds the default value on creation for the "records" field.
	DefaultRecords string
)

// OrderOption defines the ordering options for the EVENT_RECORDS queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByEventID orders the results by the event_id field.
func ByEventID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventID, opts...).ToFunc()
}

// ByRecords orders the results by the records field.
func ByRecords(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecords, opts...).ToFunc()
}

// ByMadeField orders the results by made field.
func ByMadeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMadeStep(), sql.OrderByField(field, opts...))
	}
}

// ByParticipatesField orders the results by participates field.
func ByParticipatesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParticipatesStep(), sql.OrderByField(field, opts...))
	}
}

// ByHasCount orders the results by has count.
func ByHasCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHasStep(), opts...)
	}
}

// ByHas orders the results by has terms.
func ByHas(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHasStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMadeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MadeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MadeTable, MadeColumn),
	)
}
func newParticipatesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParticipatesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ParticipatesTable, ParticipatesColumn),
	)
}
func newHasStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HasInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HasTable, HasColumn),
	)
}