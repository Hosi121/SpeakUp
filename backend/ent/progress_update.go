// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/predicate"
	"github.com/Hosi121/SpeakUp/ent/progress"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// PROGRESSUpdate is the builder for updating PROGRESS entities.
type PROGRESSUpdate struct {
	config
	hooks    []Hook
	mutation *PROGRESSMutation
}

// Where appends a list predicates to the PROGRESSUpdate builder.
func (pu *PROGRESSUpdate) Where(ps ...predicate.PROGRESS) *PROGRESSUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUserID sets the "user_id" field.
func (pu *PROGRESSUpdate) SetUserID(i int) *PROGRESSUpdate {
	pu.mutation.ResetUserID()
	pu.mutation.SetUserID(i)
	return pu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (pu *PROGRESSUpdate) SetNillableUserID(i *int) *PROGRESSUpdate {
	if i != nil {
		pu.SetUserID(*i)
	}
	return pu
}

// AddUserID adds i to the "user_id" field.
func (pu *PROGRESSUpdate) AddUserID(i int) *PROGRESSUpdate {
	pu.mutation.AddUserID(i)
	return pu
}

// SetLoginDays sets the "login_days" field.
func (pu *PROGRESSUpdate) SetLoginDays(i int) *PROGRESSUpdate {
	pu.mutation.ResetLoginDays()
	pu.mutation.SetLoginDays(i)
	return pu
}

// SetNillableLoginDays sets the "login_days" field if the given value is not nil.
func (pu *PROGRESSUpdate) SetNillableLoginDays(i *int) *PROGRESSUpdate {
	if i != nil {
		pu.SetLoginDays(*i)
	}
	return pu
}

// AddLoginDays adds i to the "login_days" field.
func (pu *PROGRESSUpdate) AddLoginDays(i int) *PROGRESSUpdate {
	pu.mutation.AddLoginDays(i)
	return pu
}

// SetConsecutiveParticipants sets the "consecutive_participants" field.
func (pu *PROGRESSUpdate) SetConsecutiveParticipants(i int) *PROGRESSUpdate {
	pu.mutation.ResetConsecutiveParticipants()
	pu.mutation.SetConsecutiveParticipants(i)
	return pu
}

// SetNillableConsecutiveParticipants sets the "consecutive_participants" field if the given value is not nil.
func (pu *PROGRESSUpdate) SetNillableConsecutiveParticipants(i *int) *PROGRESSUpdate {
	if i != nil {
		pu.SetConsecutiveParticipants(*i)
	}
	return pu
}

// AddConsecutiveParticipants adds i to the "consecutive_participants" field.
func (pu *PROGRESSUpdate) AddConsecutiveParticipants(i int) *PROGRESSUpdate {
	pu.mutation.AddConsecutiveParticipants(i)
	return pu
}

// SetConsecutiveLoginDays sets the "consecutive_login_days" field.
func (pu *PROGRESSUpdate) SetConsecutiveLoginDays(i int) *PROGRESSUpdate {
	pu.mutation.ResetConsecutiveLoginDays()
	pu.mutation.SetConsecutiveLoginDays(i)
	return pu
}

// SetNillableConsecutiveLoginDays sets the "consecutive_login_days" field if the given value is not nil.
func (pu *PROGRESSUpdate) SetNillableConsecutiveLoginDays(i *int) *PROGRESSUpdate {
	if i != nil {
		pu.SetConsecutiveLoginDays(*i)
	}
	return pu
}

// AddConsecutiveLoginDays adds i to the "consecutive_login_days" field.
func (pu *PROGRESSUpdate) AddConsecutiveLoginDays(i int) *PROGRESSUpdate {
	pu.mutation.AddConsecutiveLoginDays(i)
	return pu
}

// SetRecordedID sets the "recorded" edge to the USERS entity by ID.
func (pu *PROGRESSUpdate) SetRecordedID(id int) *PROGRESSUpdate {
	pu.mutation.SetRecordedID(id)
	return pu
}

// SetNillableRecordedID sets the "recorded" edge to the USERS entity by ID if the given value is not nil.
func (pu *PROGRESSUpdate) SetNillableRecordedID(id *int) *PROGRESSUpdate {
	if id != nil {
		pu = pu.SetRecordedID(*id)
	}
	return pu
}

// SetRecorded sets the "recorded" edge to the USERS entity.
func (pu *PROGRESSUpdate) SetRecorded(u *USERS) *PROGRESSUpdate {
	return pu.SetRecordedID(u.ID)
}

// Mutation returns the PROGRESSMutation object of the builder.
func (pu *PROGRESSUpdate) Mutation() *PROGRESSMutation {
	return pu.mutation
}

// ClearRecorded clears the "recorded" edge to the USERS entity.
func (pu *PROGRESSUpdate) ClearRecorded() *PROGRESSUpdate {
	pu.mutation.ClearRecorded()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PROGRESSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PROGRESSUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PROGRESSUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PROGRESSUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PROGRESSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(progress.Table, progress.Columns, sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UserID(); ok {
		_spec.SetField(progress.FieldUserID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedUserID(); ok {
		_spec.AddField(progress.FieldUserID, field.TypeInt, value)
	}
	if value, ok := pu.mutation.LoginDays(); ok {
		_spec.SetField(progress.FieldLoginDays, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedLoginDays(); ok {
		_spec.AddField(progress.FieldLoginDays, field.TypeInt, value)
	}
	if value, ok := pu.mutation.ConsecutiveParticipants(); ok {
		_spec.SetField(progress.FieldConsecutiveParticipants, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedConsecutiveParticipants(); ok {
		_spec.AddField(progress.FieldConsecutiveParticipants, field.TypeInt, value)
	}
	if value, ok := pu.mutation.ConsecutiveLoginDays(); ok {
		_spec.SetField(progress.FieldConsecutiveLoginDays, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedConsecutiveLoginDays(); ok {
		_spec.AddField(progress.FieldConsecutiveLoginDays, field.TypeInt, value)
	}
	if pu.mutation.RecordedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   progress.RecordedTable,
			Columns: []string{progress.RecordedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RecordedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   progress.RecordedTable,
			Columns: []string{progress.RecordedColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{progress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PROGRESSUpdateOne is the builder for updating a single PROGRESS entity.
type PROGRESSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PROGRESSMutation
}

// SetUserID sets the "user_id" field.
func (puo *PROGRESSUpdateOne) SetUserID(i int) *PROGRESSUpdateOne {
	puo.mutation.ResetUserID()
	puo.mutation.SetUserID(i)
	return puo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (puo *PROGRESSUpdateOne) SetNillableUserID(i *int) *PROGRESSUpdateOne {
	if i != nil {
		puo.SetUserID(*i)
	}
	return puo
}

// AddUserID adds i to the "user_id" field.
func (puo *PROGRESSUpdateOne) AddUserID(i int) *PROGRESSUpdateOne {
	puo.mutation.AddUserID(i)
	return puo
}

// SetLoginDays sets the "login_days" field.
func (puo *PROGRESSUpdateOne) SetLoginDays(i int) *PROGRESSUpdateOne {
	puo.mutation.ResetLoginDays()
	puo.mutation.SetLoginDays(i)
	return puo
}

// SetNillableLoginDays sets the "login_days" field if the given value is not nil.
func (puo *PROGRESSUpdateOne) SetNillableLoginDays(i *int) *PROGRESSUpdateOne {
	if i != nil {
		puo.SetLoginDays(*i)
	}
	return puo
}

// AddLoginDays adds i to the "login_days" field.
func (puo *PROGRESSUpdateOne) AddLoginDays(i int) *PROGRESSUpdateOne {
	puo.mutation.AddLoginDays(i)
	return puo
}

// SetConsecutiveParticipants sets the "consecutive_participants" field.
func (puo *PROGRESSUpdateOne) SetConsecutiveParticipants(i int) *PROGRESSUpdateOne {
	puo.mutation.ResetConsecutiveParticipants()
	puo.mutation.SetConsecutiveParticipants(i)
	return puo
}

// SetNillableConsecutiveParticipants sets the "consecutive_participants" field if the given value is not nil.
func (puo *PROGRESSUpdateOne) SetNillableConsecutiveParticipants(i *int) *PROGRESSUpdateOne {
	if i != nil {
		puo.SetConsecutiveParticipants(*i)
	}
	return puo
}

// AddConsecutiveParticipants adds i to the "consecutive_participants" field.
func (puo *PROGRESSUpdateOne) AddConsecutiveParticipants(i int) *PROGRESSUpdateOne {
	puo.mutation.AddConsecutiveParticipants(i)
	return puo
}

// SetConsecutiveLoginDays sets the "consecutive_login_days" field.
func (puo *PROGRESSUpdateOne) SetConsecutiveLoginDays(i int) *PROGRESSUpdateOne {
	puo.mutation.ResetConsecutiveLoginDays()
	puo.mutation.SetConsecutiveLoginDays(i)
	return puo
}

// SetNillableConsecutiveLoginDays sets the "consecutive_login_days" field if the given value is not nil.
func (puo *PROGRESSUpdateOne) SetNillableConsecutiveLoginDays(i *int) *PROGRESSUpdateOne {
	if i != nil {
		puo.SetConsecutiveLoginDays(*i)
	}
	return puo
}

// AddConsecutiveLoginDays adds i to the "consecutive_login_days" field.
func (puo *PROGRESSUpdateOne) AddConsecutiveLoginDays(i int) *PROGRESSUpdateOne {
	puo.mutation.AddConsecutiveLoginDays(i)
	return puo
}

// SetRecordedID sets the "recorded" edge to the USERS entity by ID.
func (puo *PROGRESSUpdateOne) SetRecordedID(id int) *PROGRESSUpdateOne {
	puo.mutation.SetRecordedID(id)
	return puo
}

// SetNillableRecordedID sets the "recorded" edge to the USERS entity by ID if the given value is not nil.
func (puo *PROGRESSUpdateOne) SetNillableRecordedID(id *int) *PROGRESSUpdateOne {
	if id != nil {
		puo = puo.SetRecordedID(*id)
	}
	return puo
}

// SetRecorded sets the "recorded" edge to the USERS entity.
func (puo *PROGRESSUpdateOne) SetRecorded(u *USERS) *PROGRESSUpdateOne {
	return puo.SetRecordedID(u.ID)
}

// Mutation returns the PROGRESSMutation object of the builder.
func (puo *PROGRESSUpdateOne) Mutation() *PROGRESSMutation {
	return puo.mutation
}

// ClearRecorded clears the "recorded" edge to the USERS entity.
func (puo *PROGRESSUpdateOne) ClearRecorded() *PROGRESSUpdateOne {
	puo.mutation.ClearRecorded()
	return puo
}

// Where appends a list predicates to the PROGRESSUpdate builder.
func (puo *PROGRESSUpdateOne) Where(ps ...predicate.PROGRESS) *PROGRESSUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PROGRESSUpdateOne) Select(field string, fields ...string) *PROGRESSUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated PROGRESS entity.
func (puo *PROGRESSUpdateOne) Save(ctx context.Context) (*PROGRESS, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PROGRESSUpdateOne) SaveX(ctx context.Context) *PROGRESS {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PROGRESSUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PROGRESSUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PROGRESSUpdateOne) sqlSave(ctx context.Context) (_node *PROGRESS, err error) {
	_spec := sqlgraph.NewUpdateSpec(progress.Table, progress.Columns, sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PROGRESS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, progress.FieldID)
		for _, f := range fields {
			if !progress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != progress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UserID(); ok {
		_spec.SetField(progress.FieldUserID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedUserID(); ok {
		_spec.AddField(progress.FieldUserID, field.TypeInt, value)
	}
	if value, ok := puo.mutation.LoginDays(); ok {
		_spec.SetField(progress.FieldLoginDays, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedLoginDays(); ok {
		_spec.AddField(progress.FieldLoginDays, field.TypeInt, value)
	}
	if value, ok := puo.mutation.ConsecutiveParticipants(); ok {
		_spec.SetField(progress.FieldConsecutiveParticipants, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedConsecutiveParticipants(); ok {
		_spec.AddField(progress.FieldConsecutiveParticipants, field.TypeInt, value)
	}
	if value, ok := puo.mutation.ConsecutiveLoginDays(); ok {
		_spec.SetField(progress.FieldConsecutiveLoginDays, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedConsecutiveLoginDays(); ok {
		_spec.AddField(progress.FieldConsecutiveLoginDays, field.TypeInt, value)
	}
	if puo.mutation.RecordedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   progress.RecordedTable,
			Columns: []string{progress.RecordedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RecordedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   progress.RecordedTable,
			Columns: []string{progress.RecordedColumn},
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
	_node = &PROGRESS{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{progress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}