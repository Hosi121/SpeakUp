// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/calls"
	"github.com/Hosi121/SpeakUp/ent/matchings"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// CALLSUpdate is the builder for updating CALLS entities.
type CALLSUpdate struct {
	config
	hooks    []Hook
	mutation *CALLSMutation
}

// Where appends a list predicates to the CALLSUpdate builder.
func (cu *CALLSUpdate) Where(ps ...predicate.CALLS) *CALLSUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetSessionID sets the "session_id" field.
func (cu *CALLSUpdate) SetSessionID(i int) *CALLSUpdate {
	cu.mutation.ResetSessionID()
	cu.mutation.SetSessionID(i)
	return cu
}

// SetNillableSessionID sets the "session_id" field if the given value is not nil.
func (cu *CALLSUpdate) SetNillableSessionID(i *int) *CALLSUpdate {
	if i != nil {
		cu.SetSessionID(*i)
	}
	return cu
}

// AddSessionID adds i to the "session_id" field.
func (cu *CALLSUpdate) AddSessionID(i int) *CALLSUpdate {
	cu.mutation.AddSessionID(i)
	return cu
}

// SetCallStart sets the "call_start" field.
func (cu *CALLSUpdate) SetCallStart(t time.Time) *CALLSUpdate {
	cu.mutation.SetCallStart(t)
	return cu
}

// SetNillableCallStart sets the "call_start" field if the given value is not nil.
func (cu *CALLSUpdate) SetNillableCallStart(t *time.Time) *CALLSUpdate {
	if t != nil {
		cu.SetCallStart(*t)
	}
	return cu
}

// SetCallEnd sets the "call_end" field.
func (cu *CALLSUpdate) SetCallEnd(t time.Time) *CALLSUpdate {
	cu.mutation.SetCallEnd(t)
	return cu
}

// SetNillableCallEnd sets the "call_end" field if the given value is not nil.
func (cu *CALLSUpdate) SetNillableCallEnd(t *time.Time) *CALLSUpdate {
	if t != nil {
		cu.SetCallEnd(*t)
	}
	return cu
}

// SetRating sets the "rating" field.
func (cu *CALLSUpdate) SetRating(i int) *CALLSUpdate {
	cu.mutation.ResetRating()
	cu.mutation.SetRating(i)
	return cu
}

// SetNillableRating sets the "rating" field if the given value is not nil.
func (cu *CALLSUpdate) SetNillableRating(i *int) *CALLSUpdate {
	if i != nil {
		cu.SetRating(*i)
	}
	return cu
}

// AddRating adds i to the "rating" field.
func (cu *CALLSUpdate) AddRating(i int) *CALLSUpdate {
	cu.mutation.AddRating(i)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CALLSUpdate) SetCreatedAt(t time.Time) *CALLSUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CALLSUpdate) SetNillableCreatedAt(t *time.Time) *CALLSUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetMadeID sets the "made" edge to the MATCHINGS entity by ID.
func (cu *CALLSUpdate) SetMadeID(id int) *CALLSUpdate {
	cu.mutation.SetMadeID(id)
	return cu
}

// SetNillableMadeID sets the "made" edge to the MATCHINGS entity by ID if the given value is not nil.
func (cu *CALLSUpdate) SetNillableMadeID(id *int) *CALLSUpdate {
	if id != nil {
		cu = cu.SetMadeID(*id)
	}
	return cu
}

// SetMade sets the "made" edge to the MATCHINGS entity.
func (cu *CALLSUpdate) SetMade(m *MATCHINGS) *CALLSUpdate {
	return cu.SetMadeID(m.ID)
}

// Mutation returns the CALLSMutation object of the builder.
func (cu *CALLSUpdate) Mutation() *CALLSMutation {
	return cu.mutation
}

// ClearMade clears the "made" edge to the MATCHINGS entity.
func (cu *CALLSUpdate) ClearMade() *CALLSUpdate {
	cu.mutation.ClearMade()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CALLSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CALLSUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CALLSUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CALLSUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CALLSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(calls.Table, calls.Columns, sqlgraph.NewFieldSpec(calls.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.SessionID(); ok {
		_spec.SetField(calls.FieldSessionID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedSessionID(); ok {
		_spec.AddField(calls.FieldSessionID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.CallStart(); ok {
		_spec.SetField(calls.FieldCallStart, field.TypeTime, value)
	}
	if value, ok := cu.mutation.CallEnd(); ok {
		_spec.SetField(calls.FieldCallEnd, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Rating(); ok {
		_spec.SetField(calls.FieldRating, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedRating(); ok {
		_spec.AddField(calls.FieldRating, field.TypeInt, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(calls.FieldCreatedAt, field.TypeTime, value)
	}
	if cu.mutation.MadeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   calls.MadeTable,
			Columns: []string{calls.MadeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(matchings.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MadeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   calls.MadeTable,
			Columns: []string{calls.MadeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(matchings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{calls.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CALLSUpdateOne is the builder for updating a single CALLS entity.
type CALLSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CALLSMutation
}

// SetSessionID sets the "session_id" field.
func (cuo *CALLSUpdateOne) SetSessionID(i int) *CALLSUpdateOne {
	cuo.mutation.ResetSessionID()
	cuo.mutation.SetSessionID(i)
	return cuo
}

// SetNillableSessionID sets the "session_id" field if the given value is not nil.
func (cuo *CALLSUpdateOne) SetNillableSessionID(i *int) *CALLSUpdateOne {
	if i != nil {
		cuo.SetSessionID(*i)
	}
	return cuo
}

// AddSessionID adds i to the "session_id" field.
func (cuo *CALLSUpdateOne) AddSessionID(i int) *CALLSUpdateOne {
	cuo.mutation.AddSessionID(i)
	return cuo
}

// SetCallStart sets the "call_start" field.
func (cuo *CALLSUpdateOne) SetCallStart(t time.Time) *CALLSUpdateOne {
	cuo.mutation.SetCallStart(t)
	return cuo
}

// SetNillableCallStart sets the "call_start" field if the given value is not nil.
func (cuo *CALLSUpdateOne) SetNillableCallStart(t *time.Time) *CALLSUpdateOne {
	if t != nil {
		cuo.SetCallStart(*t)
	}
	return cuo
}

// SetCallEnd sets the "call_end" field.
func (cuo *CALLSUpdateOne) SetCallEnd(t time.Time) *CALLSUpdateOne {
	cuo.mutation.SetCallEnd(t)
	return cuo
}

// SetNillableCallEnd sets the "call_end" field if the given value is not nil.
func (cuo *CALLSUpdateOne) SetNillableCallEnd(t *time.Time) *CALLSUpdateOne {
	if t != nil {
		cuo.SetCallEnd(*t)
	}
	return cuo
}

// SetRating sets the "rating" field.
func (cuo *CALLSUpdateOne) SetRating(i int) *CALLSUpdateOne {
	cuo.mutation.ResetRating()
	cuo.mutation.SetRating(i)
	return cuo
}

// SetNillableRating sets the "rating" field if the given value is not nil.
func (cuo *CALLSUpdateOne) SetNillableRating(i *int) *CALLSUpdateOne {
	if i != nil {
		cuo.SetRating(*i)
	}
	return cuo
}

// AddRating adds i to the "rating" field.
func (cuo *CALLSUpdateOne) AddRating(i int) *CALLSUpdateOne {
	cuo.mutation.AddRating(i)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CALLSUpdateOne) SetCreatedAt(t time.Time) *CALLSUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CALLSUpdateOne) SetNillableCreatedAt(t *time.Time) *CALLSUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetMadeID sets the "made" edge to the MATCHINGS entity by ID.
func (cuo *CALLSUpdateOne) SetMadeID(id int) *CALLSUpdateOne {
	cuo.mutation.SetMadeID(id)
	return cuo
}

// SetNillableMadeID sets the "made" edge to the MATCHINGS entity by ID if the given value is not nil.
func (cuo *CALLSUpdateOne) SetNillableMadeID(id *int) *CALLSUpdateOne {
	if id != nil {
		cuo = cuo.SetMadeID(*id)
	}
	return cuo
}

// SetMade sets the "made" edge to the MATCHINGS entity.
func (cuo *CALLSUpdateOne) SetMade(m *MATCHINGS) *CALLSUpdateOne {
	return cuo.SetMadeID(m.ID)
}

// Mutation returns the CALLSMutation object of the builder.
func (cuo *CALLSUpdateOne) Mutation() *CALLSMutation {
	return cuo.mutation
}

// ClearMade clears the "made" edge to the MATCHINGS entity.
func (cuo *CALLSUpdateOne) ClearMade() *CALLSUpdateOne {
	cuo.mutation.ClearMade()
	return cuo
}

// Where appends a list predicates to the CALLSUpdate builder.
func (cuo *CALLSUpdateOne) Where(ps ...predicate.CALLS) *CALLSUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CALLSUpdateOne) Select(field string, fields ...string) *CALLSUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated CALLS entity.
func (cuo *CALLSUpdateOne) Save(ctx context.Context) (*CALLS, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CALLSUpdateOne) SaveX(ctx context.Context) *CALLS {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CALLSUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CALLSUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CALLSUpdateOne) sqlSave(ctx context.Context) (_node *CALLS, err error) {
	_spec := sqlgraph.NewUpdateSpec(calls.Table, calls.Columns, sqlgraph.NewFieldSpec(calls.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CALLS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, calls.FieldID)
		for _, f := range fields {
			if !calls.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != calls.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.SessionID(); ok {
		_spec.SetField(calls.FieldSessionID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedSessionID(); ok {
		_spec.AddField(calls.FieldSessionID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.CallStart(); ok {
		_spec.SetField(calls.FieldCallStart, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.CallEnd(); ok {
		_spec.SetField(calls.FieldCallEnd, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Rating(); ok {
		_spec.SetField(calls.FieldRating, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedRating(); ok {
		_spec.AddField(calls.FieldRating, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(calls.FieldCreatedAt, field.TypeTime, value)
	}
	if cuo.mutation.MadeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   calls.MadeTable,
			Columns: []string{calls.MadeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(matchings.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MadeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   calls.MadeTable,
			Columns: []string{calls.MadeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(matchings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CALLS{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{calls.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
