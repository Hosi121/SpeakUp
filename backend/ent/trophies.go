// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Hosi121/SpeakUp/ent/trophies"
)

// TROPHIES is the model entity for the TROPHIES schema.
type TROPHIES struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Requirement holds the value of the "requirement" field.
	Requirement string `json:"requirement,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TROPHIESQuery when eager-loading is set.
	Edges        TROPHIESEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TROPHIESEdges holds the relations/edges for other nodes in the graph.
type TROPHIESEdges struct {
	// Refered holds the value of the refered edge.
	Refered []*ACHIEVEMENTS `json:"refered,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ReferedOrErr returns the Refered value or an error if the edge
// was not loaded in eager-loading.
func (e TROPHIESEdges) ReferedOrErr() ([]*ACHIEVEMENTS, error) {
	if e.loadedTypes[0] {
		return e.Refered, nil
	}
	return nil, &NotLoadedError{edge: "refered"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TROPHIES) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case trophies.FieldID:
			values[i] = new(sql.NullInt64)
		case trophies.FieldTitle, trophies.FieldDescription, trophies.FieldRequirement:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TROPHIES fields.
func (t *TROPHIES) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case trophies.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case trophies.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case trophies.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case trophies.FieldRequirement:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field requirement", values[i])
			} else if value.Valid {
				t.Requirement = value.String
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TROPHIES.
// This includes values selected through modifiers, order, etc.
func (t *TROPHIES) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryRefered queries the "refered" edge of the TROPHIES entity.
func (t *TROPHIES) QueryRefered() *ACHIEVEMENTSQuery {
	return NewTROPHIESClient(t.config).QueryRefered(t)
}

// Update returns a builder for updating this TROPHIES.
// Note that you need to call TROPHIES.Unwrap() before calling this method if this TROPHIES
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *TROPHIES) Update() *TROPHIESUpdateOne {
	return NewTROPHIESClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the TROPHIES entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *TROPHIES) Unwrap() *TROPHIES {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: TROPHIES is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *TROPHIES) String() string {
	var builder strings.Builder
	builder.WriteString("TROPHIES(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("title=")
	builder.WriteString(t.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("requirement=")
	builder.WriteString(t.Requirement)
	builder.WriteByte(')')
	return builder.String()
}

// TROPHIESs is a parsable slice of TROPHIES.
type TROPHIESs []*TROPHIES
