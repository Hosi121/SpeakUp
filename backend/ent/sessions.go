// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Hosi121/SpeakUp/ent/calls"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/sessions"
)

// SESSIONS is the model entity for the SESSIONS schema.
type SESSIONS struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// MatchedUserID holds the value of the "matched_user_id" field.
	MatchedUserID int `json:"matched_user_id,omitempty"`
	// RecordID holds the value of the "record_id" field.
	RecordID int `json:"record_id,omitempty"`
	// MatchedAt holds the value of the "matched_at" field.
	MatchedAt time.Time `json:"matched_at,omitempty"`
	// Status holds the value of the "status" field.
	Status sessions.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SESSIONSQuery when eager-loading is set.
	Edges             SESSIONSEdges `json:"edges"`
	event_records_has *int
	selectValues      sql.SelectValues
}

// SESSIONSEdges holds the relations/edges for other nodes in the graph.
type SESSIONSEdges struct {
	// Had holds the value of the had edge.
	Had *EVENT_RECORDS `json:"had,omitempty"`
	// Makes holds the value of the makes edge.
	Makes *CALLS `json:"makes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// HadOrErr returns the Had value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SESSIONSEdges) HadOrErr() (*EVENT_RECORDS, error) {
	if e.Had != nil {
		return e.Had, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: event_records.Label}
	}
	return nil, &NotLoadedError{edge: "had"}
}

// MakesOrErr returns the Makes value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SESSIONSEdges) MakesOrErr() (*CALLS, error) {
	if e.Makes != nil {
		return e.Makes, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: calls.Label}
	}
	return nil, &NotLoadedError{edge: "makes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SESSIONS) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sessions.FieldID, sessions.FieldUserID, sessions.FieldMatchedUserID, sessions.FieldRecordID:
			values[i] = new(sql.NullInt64)
		case sessions.FieldStatus:
			values[i] = new(sql.NullString)
		case sessions.FieldMatchedAt:
			values[i] = new(sql.NullTime)
		case sessions.ForeignKeys[0]: // event_records_has
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SESSIONS fields.
func (s *SESSIONS) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sessions.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case sessions.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = int(value.Int64)
			}
		case sessions.FieldMatchedUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field matched_user_id", values[i])
			} else if value.Valid {
				s.MatchedUserID = int(value.Int64)
			}
		case sessions.FieldRecordID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field record_id", values[i])
			} else if value.Valid {
				s.RecordID = int(value.Int64)
			}
		case sessions.FieldMatchedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field matched_at", values[i])
			} else if value.Valid {
				s.MatchedAt = value.Time
			}
		case sessions.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = sessions.Status(value.String)
			}
		case sessions.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field event_records_has", value)
			} else if value.Valid {
				s.event_records_has = new(int)
				*s.event_records_has = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SESSIONS.
// This includes values selected through modifiers, order, etc.
func (s *SESSIONS) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryHad queries the "had" edge of the SESSIONS entity.
func (s *SESSIONS) QueryHad() *EVENTRECORDSQuery {
	return NewSESSIONSClient(s.config).QueryHad(s)
}

// QueryMakes queries the "makes" edge of the SESSIONS entity.
func (s *SESSIONS) QueryMakes() *CALLSQuery {
	return NewSESSIONSClient(s.config).QueryMakes(s)
}

// Update returns a builder for updating this SESSIONS.
// Note that you need to call SESSIONS.Unwrap() before calling this method if this SESSIONS
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *SESSIONS) Update() *SESSIONSUpdateOne {
	return NewSESSIONSClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the SESSIONS entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *SESSIONS) Unwrap() *SESSIONS {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: SESSIONS is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *SESSIONS) String() string {
	var builder strings.Builder
	builder.WriteString("SESSIONS(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteString(", ")
	builder.WriteString("matched_user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.MatchedUserID))
	builder.WriteString(", ")
	builder.WriteString("record_id=")
	builder.WriteString(fmt.Sprintf("%v", s.RecordID))
	builder.WriteString(", ")
	builder.WriteString("matched_at=")
	builder.WriteString(s.MatchedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteByte(')')
	return builder.String()
}

// SESSIONSs is a parsable slice of SESSIONS.
type SESSIONSs []*SESSIONS
