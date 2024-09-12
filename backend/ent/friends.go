// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Hosi121/SpeakUp/ent/friends"
)

// FRIENDS is the model entity for the FRIENDS schema.
type FRIENDS struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// TargetUserID holds the value of the "target_user_id" field.
	TargetUserID int `json:"target_user_id,omitempty"`
	// Status holds the value of the "status" field.
	Status friends.Status `json:"status,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FRIENDSQuery when eager-loading is set.
	Edges        FRIENDSEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FRIENDSEdges holds the relations/edges for other nodes in the graph.
type FRIENDSEdges struct {
	// Connects holds the value of the connects edge.
	Connects []*USERS `json:"connects,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ConnectsOrErr returns the Connects value or an error if the edge
// was not loaded in eager-loading.
func (e FRIENDSEdges) ConnectsOrErr() ([]*USERS, error) {
	if e.loadedTypes[0] {
		return e.Connects, nil
	}
	return nil, &NotLoadedError{edge: "connects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FRIENDS) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case friends.FieldID, friends.FieldUserID, friends.FieldTargetUserID:
			values[i] = new(sql.NullInt64)
		case friends.FieldStatus:
			values[i] = new(sql.NullString)
		case friends.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FRIENDS fields.
func (f *FRIENDS) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case friends.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case friends.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				f.UserID = int(value.Int64)
			}
		case friends.FieldTargetUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field target_user_id", values[i])
			} else if value.Valid {
				f.TargetUserID = int(value.Int64)
			}
		case friends.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = friends.Status(value.String)
			}
		case friends.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FRIENDS.
// This includes values selected through modifiers, order, etc.
func (f *FRIENDS) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryConnects queries the "connects" edge of the FRIENDS entity.
func (f *FRIENDS) QueryConnects() *USERSQuery {
	return NewFRIENDSClient(f.config).QueryConnects(f)
}

// Update returns a builder for updating this FRIENDS.
// Note that you need to call FRIENDS.Unwrap() before calling this method if this FRIENDS
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *FRIENDS) Update() *FRIENDSUpdateOne {
	return NewFRIENDSClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the FRIENDS entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *FRIENDS) Unwrap() *FRIENDS {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: FRIENDS is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *FRIENDS) String() string {
	var builder strings.Builder
	builder.WriteString("FRIENDS(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", f.UserID))
	builder.WriteString(", ")
	builder.WriteString("target_user_id=")
	builder.WriteString(fmt.Sprintf("%v", f.TargetUserID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", f.Status))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// FRIENDSs is a parsable slice of FRIENDS.
type FRIENDSs []*FRIENDS