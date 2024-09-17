// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/events"
	"github.com/Hosi121/SpeakUp/ent/predicate"
	"github.com/Hosi121/SpeakUp/ent/sessions"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// EVENTRECORDSUpdate is the builder for updating EVENT_RECORDS entities.
type EVENTRECORDSUpdate struct {
	config
	hooks    []Hook
	mutation *EVENTRECORDSMutation
}

// Where appends a list predicates to the EVENTRECORDSUpdate builder.
func (eu *EVENTRECORDSUpdate) Where(ps ...predicate.EVENT_RECORDS) *EVENTRECORDSUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUserID sets the "user_id" field.
func (eu *EVENTRECORDSUpdate) SetUserID(i int) *EVENTRECORDSUpdate {
	eu.mutation.ResetUserID()
	eu.mutation.SetUserID(i)
	return eu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (eu *EVENTRECORDSUpdate) SetNillableUserID(i *int) *EVENTRECORDSUpdate {
	if i != nil {
		eu.SetUserID(*i)
	}
	return eu
}

// AddUserID adds i to the "user_id" field.
func (eu *EVENTRECORDSUpdate) AddUserID(i int) *EVENTRECORDSUpdate {
	eu.mutation.AddUserID(i)
	return eu
}

// SetEventID sets the "event_id" field.
func (eu *EVENTRECORDSUpdate) SetEventID(i int) *EVENTRECORDSUpdate {
	eu.mutation.ResetEventID()
	eu.mutation.SetEventID(i)
	return eu
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (eu *EVENTRECORDSUpdate) SetNillableEventID(i *int) *EVENTRECORDSUpdate {
	if i != nil {
		eu.SetEventID(*i)
	}
	return eu
}

// AddEventID adds i to the "event_id" field.
func (eu *EVENTRECORDSUpdate) AddEventID(i int) *EVENTRECORDSUpdate {
	eu.mutation.AddEventID(i)
	return eu
}

// SetRecords sets the "records" field.
func (eu *EVENTRECORDSUpdate) SetRecords(s string) *EVENTRECORDSUpdate {
	eu.mutation.SetRecords(s)
	return eu
}

// SetNillableRecords sets the "records" field if the given value is not nil.
func (eu *EVENTRECORDSUpdate) SetNillableRecords(s *string) *EVENTRECORDSUpdate {
	if s != nil {
		eu.SetRecords(*s)
	}
	return eu
}

// SetMadeID sets the "made" edge to the USERS entity by ID.
func (eu *EVENTRECORDSUpdate) SetMadeID(id int) *EVENTRECORDSUpdate {
	eu.mutation.SetMadeID(id)
	return eu
}

// SetNillableMadeID sets the "made" edge to the USERS entity by ID if the given value is not nil.
func (eu *EVENTRECORDSUpdate) SetNillableMadeID(id *int) *EVENTRECORDSUpdate {
	if id != nil {
		eu = eu.SetMadeID(*id)
	}
	return eu
}

// SetMade sets the "made" edge to the USERS entity.
func (eu *EVENTRECORDSUpdate) SetMade(u *USERS) *EVENTRECORDSUpdate {
	return eu.SetMadeID(u.ID)
}

// SetParticipatesID sets the "participates" edge to the EVENTS entity by ID.
func (eu *EVENTRECORDSUpdate) SetParticipatesID(id int) *EVENTRECORDSUpdate {
	eu.mutation.SetParticipatesID(id)
	return eu
}

// SetNillableParticipatesID sets the "participates" edge to the EVENTS entity by ID if the given value is not nil.
func (eu *EVENTRECORDSUpdate) SetNillableParticipatesID(id *int) *EVENTRECORDSUpdate {
	if id != nil {
		eu = eu.SetParticipatesID(*id)
	}
	return eu
}

// SetParticipates sets the "participates" edge to the EVENTS entity.
func (eu *EVENTRECORDSUpdate) SetParticipates(e *EVENTS) *EVENTRECORDSUpdate {
	return eu.SetParticipatesID(e.ID)
}

// AddHaIDs adds the "has" edge to the SESSIONS entity by IDs.
func (eu *EVENTRECORDSUpdate) AddHaIDs(ids ...int) *EVENTRECORDSUpdate {
	eu.mutation.AddHaIDs(ids...)
	return eu
}

// AddHas adds the "has" edges to the SESSIONS entity.
func (eu *EVENTRECORDSUpdate) AddHas(s ...*SESSIONS) *EVENTRECORDSUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.AddHaIDs(ids...)
}

// Mutation returns the EVENTRECORDSMutation object of the builder.
func (eu *EVENTRECORDSUpdate) Mutation() *EVENTRECORDSMutation {
	return eu.mutation
}

// ClearMade clears the "made" edge to the USERS entity.
func (eu *EVENTRECORDSUpdate) ClearMade() *EVENTRECORDSUpdate {
	eu.mutation.ClearMade()
	return eu
}

// ClearParticipates clears the "participates" edge to the EVENTS entity.
func (eu *EVENTRECORDSUpdate) ClearParticipates() *EVENTRECORDSUpdate {
	eu.mutation.ClearParticipates()
	return eu
}

// ClearHas clears all "has" edges to the SESSIONS entity.
func (eu *EVENTRECORDSUpdate) ClearHas() *EVENTRECORDSUpdate {
	eu.mutation.ClearHas()
	return eu
}

// RemoveHaIDs removes the "has" edge to SESSIONS entities by IDs.
func (eu *EVENTRECORDSUpdate) RemoveHaIDs(ids ...int) *EVENTRECORDSUpdate {
	eu.mutation.RemoveHaIDs(ids...)
	return eu
}

// RemoveHas removes "has" edges to SESSIONS entities.
func (eu *EVENTRECORDSUpdate) RemoveHas(s ...*SESSIONS) *EVENTRECORDSUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return eu.RemoveHaIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EVENTRECORDSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EVENTRECORDSUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EVENTRECORDSUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EVENTRECORDSUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EVENTRECORDSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(event_records.Table, event_records.Columns, sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UserID(); ok {
		_spec.SetField(event_records.FieldUserID, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedUserID(); ok {
		_spec.AddField(event_records.FieldUserID, field.TypeInt, value)
	}
	if value, ok := eu.mutation.EventID(); ok {
		_spec.SetField(event_records.FieldEventID, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedEventID(); ok {
		_spec.AddField(event_records.FieldEventID, field.TypeInt, value)
	}
	if value, ok := eu.mutation.Records(); ok {
		_spec.SetField(event_records.FieldRecords, field.TypeString, value)
	}
	if eu.mutation.MadeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event_records.MadeTable,
			Columns: []string{event_records.MadeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.MadeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event_records.MadeTable,
			Columns: []string{event_records.MadeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ParticipatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event_records.ParticipatesTable,
			Columns: []string{event_records.ParticipatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ParticipatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event_records.ParticipatesTable,
			Columns: []string{event_records.ParticipatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.HasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event_records.HasTable,
			Columns: []string{event_records.HasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedHasIDs(); len(nodes) > 0 && !eu.mutation.HasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event_records.HasTable,
			Columns: []string{event_records.HasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.HasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event_records.HasTable,
			Columns: []string{event_records.HasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event_records.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EVENTRECORDSUpdateOne is the builder for updating a single EVENT_RECORDS entity.
type EVENTRECORDSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EVENTRECORDSMutation
}

// SetUserID sets the "user_id" field.
func (euo *EVENTRECORDSUpdateOne) SetUserID(i int) *EVENTRECORDSUpdateOne {
	euo.mutation.ResetUserID()
	euo.mutation.SetUserID(i)
	return euo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (euo *EVENTRECORDSUpdateOne) SetNillableUserID(i *int) *EVENTRECORDSUpdateOne {
	if i != nil {
		euo.SetUserID(*i)
	}
	return euo
}

// AddUserID adds i to the "user_id" field.
func (euo *EVENTRECORDSUpdateOne) AddUserID(i int) *EVENTRECORDSUpdateOne {
	euo.mutation.AddUserID(i)
	return euo
}

// SetEventID sets the "event_id" field.
func (euo *EVENTRECORDSUpdateOne) SetEventID(i int) *EVENTRECORDSUpdateOne {
	euo.mutation.ResetEventID()
	euo.mutation.SetEventID(i)
	return euo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (euo *EVENTRECORDSUpdateOne) SetNillableEventID(i *int) *EVENTRECORDSUpdateOne {
	if i != nil {
		euo.SetEventID(*i)
	}
	return euo
}

// AddEventID adds i to the "event_id" field.
func (euo *EVENTRECORDSUpdateOne) AddEventID(i int) *EVENTRECORDSUpdateOne {
	euo.mutation.AddEventID(i)
	return euo
}

// SetRecords sets the "records" field.
func (euo *EVENTRECORDSUpdateOne) SetRecords(s string) *EVENTRECORDSUpdateOne {
	euo.mutation.SetRecords(s)
	return euo
}

// SetNillableRecords sets the "records" field if the given value is not nil.
func (euo *EVENTRECORDSUpdateOne) SetNillableRecords(s *string) *EVENTRECORDSUpdateOne {
	if s != nil {
		euo.SetRecords(*s)
	}
	return euo
}

// SetMadeID sets the "made" edge to the USERS entity by ID.
func (euo *EVENTRECORDSUpdateOne) SetMadeID(id int) *EVENTRECORDSUpdateOne {
	euo.mutation.SetMadeID(id)
	return euo
}

// SetNillableMadeID sets the "made" edge to the USERS entity by ID if the given value is not nil.
func (euo *EVENTRECORDSUpdateOne) SetNillableMadeID(id *int) *EVENTRECORDSUpdateOne {
	if id != nil {
		euo = euo.SetMadeID(*id)
	}
	return euo
}

// SetMade sets the "made" edge to the USERS entity.
func (euo *EVENTRECORDSUpdateOne) SetMade(u *USERS) *EVENTRECORDSUpdateOne {
	return euo.SetMadeID(u.ID)
}

// SetParticipatesID sets the "participates" edge to the EVENTS entity by ID.
func (euo *EVENTRECORDSUpdateOne) SetParticipatesID(id int) *EVENTRECORDSUpdateOne {
	euo.mutation.SetParticipatesID(id)
	return euo
}

// SetNillableParticipatesID sets the "participates" edge to the EVENTS entity by ID if the given value is not nil.
func (euo *EVENTRECORDSUpdateOne) SetNillableParticipatesID(id *int) *EVENTRECORDSUpdateOne {
	if id != nil {
		euo = euo.SetParticipatesID(*id)
	}
	return euo
}

// SetParticipates sets the "participates" edge to the EVENTS entity.
func (euo *EVENTRECORDSUpdateOne) SetParticipates(e *EVENTS) *EVENTRECORDSUpdateOne {
	return euo.SetParticipatesID(e.ID)
}

// AddHaIDs adds the "has" edge to the SESSIONS entity by IDs.
func (euo *EVENTRECORDSUpdateOne) AddHaIDs(ids ...int) *EVENTRECORDSUpdateOne {
	euo.mutation.AddHaIDs(ids...)
	return euo
}

// AddHas adds the "has" edges to the SESSIONS entity.
func (euo *EVENTRECORDSUpdateOne) AddHas(s ...*SESSIONS) *EVENTRECORDSUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.AddHaIDs(ids...)
}

// Mutation returns the EVENTRECORDSMutation object of the builder.
func (euo *EVENTRECORDSUpdateOne) Mutation() *EVENTRECORDSMutation {
	return euo.mutation
}

// ClearMade clears the "made" edge to the USERS entity.
func (euo *EVENTRECORDSUpdateOne) ClearMade() *EVENTRECORDSUpdateOne {
	euo.mutation.ClearMade()
	return euo
}

// ClearParticipates clears the "participates" edge to the EVENTS entity.
func (euo *EVENTRECORDSUpdateOne) ClearParticipates() *EVENTRECORDSUpdateOne {
	euo.mutation.ClearParticipates()
	return euo
}

// ClearHas clears all "has" edges to the SESSIONS entity.
func (euo *EVENTRECORDSUpdateOne) ClearHas() *EVENTRECORDSUpdateOne {
	euo.mutation.ClearHas()
	return euo
}

// RemoveHaIDs removes the "has" edge to SESSIONS entities by IDs.
func (euo *EVENTRECORDSUpdateOne) RemoveHaIDs(ids ...int) *EVENTRECORDSUpdateOne {
	euo.mutation.RemoveHaIDs(ids...)
	return euo
}

// RemoveHas removes "has" edges to SESSIONS entities.
func (euo *EVENTRECORDSUpdateOne) RemoveHas(s ...*SESSIONS) *EVENTRECORDSUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return euo.RemoveHaIDs(ids...)
}

// Where appends a list predicates to the EVENTRECORDSUpdate builder.
func (euo *EVENTRECORDSUpdateOne) Where(ps ...predicate.EVENT_RECORDS) *EVENTRECORDSUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EVENTRECORDSUpdateOne) Select(field string, fields ...string) *EVENTRECORDSUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated EVENT_RECORDS entity.
func (euo *EVENTRECORDSUpdateOne) Save(ctx context.Context) (*EVENT_RECORDS, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EVENTRECORDSUpdateOne) SaveX(ctx context.Context) *EVENT_RECORDS {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EVENTRECORDSUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EVENTRECORDSUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EVENTRECORDSUpdateOne) sqlSave(ctx context.Context) (_node *EVENT_RECORDS, err error) {
	_spec := sqlgraph.NewUpdateSpec(event_records.Table, event_records.Columns, sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EVENT_RECORDS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event_records.FieldID)
		for _, f := range fields {
			if !event_records.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event_records.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UserID(); ok {
		_spec.SetField(event_records.FieldUserID, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedUserID(); ok {
		_spec.AddField(event_records.FieldUserID, field.TypeInt, value)
	}
	if value, ok := euo.mutation.EventID(); ok {
		_spec.SetField(event_records.FieldEventID, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedEventID(); ok {
		_spec.AddField(event_records.FieldEventID, field.TypeInt, value)
	}
	if value, ok := euo.mutation.Records(); ok {
		_spec.SetField(event_records.FieldRecords, field.TypeString, value)
	}
	if euo.mutation.MadeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event_records.MadeTable,
			Columns: []string{event_records.MadeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.MadeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event_records.MadeTable,
			Columns: []string{event_records.MadeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ParticipatesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event_records.ParticipatesTable,
			Columns: []string{event_records.ParticipatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ParticipatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event_records.ParticipatesTable,
			Columns: []string{event_records.ParticipatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.HasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event_records.HasTable,
			Columns: []string{event_records.HasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedHasIDs(); len(nodes) > 0 && !euo.mutation.HasCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event_records.HasTable,
			Columns: []string{event_records.HasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.HasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event_records.HasTable,
			Columns: []string{event_records.HasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EVENT_RECORDS{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event_records.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
