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
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/predicate"
	"github.com/Hosi121/SpeakUp/ent/sessions"
)

// SESSIONSUpdate is the builder for updating SESSIONS entities.
type SESSIONSUpdate struct {
	config
	hooks    []Hook
	mutation *SESSIONSMutation
}

// Where appends a list predicates to the SESSIONSUpdate builder.
func (su *SESSIONSUpdate) Where(ps ...predicate.SESSIONS) *SESSIONSUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUserID sets the "user_id" field.
func (su *SESSIONSUpdate) SetUserID(i int) *SESSIONSUpdate {
	su.mutation.ResetUserID()
	su.mutation.SetUserID(i)
	return su
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (su *SESSIONSUpdate) SetNillableUserID(i *int) *SESSIONSUpdate {
	if i != nil {
		su.SetUserID(*i)
	}
	return su
}

// AddUserID adds i to the "user_id" field.
func (su *SESSIONSUpdate) AddUserID(i int) *SESSIONSUpdate {
	su.mutation.AddUserID(i)
	return su
}

// SetMatchedUserID sets the "matched_user_id" field.
func (su *SESSIONSUpdate) SetMatchedUserID(i int) *SESSIONSUpdate {
	su.mutation.ResetMatchedUserID()
	su.mutation.SetMatchedUserID(i)
	return su
}

// SetNillableMatchedUserID sets the "matched_user_id" field if the given value is not nil.
func (su *SESSIONSUpdate) SetNillableMatchedUserID(i *int) *SESSIONSUpdate {
	if i != nil {
		su.SetMatchedUserID(*i)
	}
	return su
}

// AddMatchedUserID adds i to the "matched_user_id" field.
func (su *SESSIONSUpdate) AddMatchedUserID(i int) *SESSIONSUpdate {
	su.mutation.AddMatchedUserID(i)
	return su
}

// SetRecordID sets the "record_id" field.
func (su *SESSIONSUpdate) SetRecordID(i int) *SESSIONSUpdate {
	su.mutation.ResetRecordID()
	su.mutation.SetRecordID(i)
	return su
}

// SetNillableRecordID sets the "record_id" field if the given value is not nil.
func (su *SESSIONSUpdate) SetNillableRecordID(i *int) *SESSIONSUpdate {
	if i != nil {
		su.SetRecordID(*i)
	}
	return su
}

// AddRecordID adds i to the "record_id" field.
func (su *SESSIONSUpdate) AddRecordID(i int) *SESSIONSUpdate {
	su.mutation.AddRecordID(i)
	return su
}

// SetMatchedAt sets the "matched_at" field.
func (su *SESSIONSUpdate) SetMatchedAt(t time.Time) *SESSIONSUpdate {
	su.mutation.SetMatchedAt(t)
	return su
}

// SetNillableMatchedAt sets the "matched_at" field if the given value is not nil.
func (su *SESSIONSUpdate) SetNillableMatchedAt(t *time.Time) *SESSIONSUpdate {
	if t != nil {
		su.SetMatchedAt(*t)
	}
	return su
}

// SetStatus sets the "status" field.
func (su *SESSIONSUpdate) SetStatus(s sessions.Status) *SESSIONSUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (su *SESSIONSUpdate) SetNillableStatus(s *sessions.Status) *SESSIONSUpdate {
	if s != nil {
		su.SetStatus(*s)
	}
	return su
}

// SetHadID sets the "had" edge to the EVENT_RECORDS entity by ID.
func (su *SESSIONSUpdate) SetHadID(id int) *SESSIONSUpdate {
	su.mutation.SetHadID(id)
	return su
}

// SetNillableHadID sets the "had" edge to the EVENT_RECORDS entity by ID if the given value is not nil.
func (su *SESSIONSUpdate) SetNillableHadID(id *int) *SESSIONSUpdate {
	if id != nil {
		su = su.SetHadID(*id)
	}
	return su
}

// SetHad sets the "had" edge to the EVENT_RECORDS entity.
func (su *SESSIONSUpdate) SetHad(e *EVENT_RECORDS) *SESSIONSUpdate {
	return su.SetHadID(e.ID)
}

// SetMakesID sets the "makes" edge to the CALLS entity by ID.
func (su *SESSIONSUpdate) SetMakesID(id int) *SESSIONSUpdate {
	su.mutation.SetMakesID(id)
	return su
}

// SetNillableMakesID sets the "makes" edge to the CALLS entity by ID if the given value is not nil.
func (su *SESSIONSUpdate) SetNillableMakesID(id *int) *SESSIONSUpdate {
	if id != nil {
		su = su.SetMakesID(*id)
	}
	return su
}

// SetMakes sets the "makes" edge to the CALLS entity.
func (su *SESSIONSUpdate) SetMakes(c *CALLS) *SESSIONSUpdate {
	return su.SetMakesID(c.ID)
}

// Mutation returns the SESSIONSMutation object of the builder.
func (su *SESSIONSUpdate) Mutation() *SESSIONSMutation {
	return su.mutation
}

// ClearHad clears the "had" edge to the EVENT_RECORDS entity.
func (su *SESSIONSUpdate) ClearHad() *SESSIONSUpdate {
	su.mutation.ClearHad()
	return su
}

// ClearMakes clears the "makes" edge to the CALLS entity.
func (su *SESSIONSUpdate) ClearMakes() *SESSIONSUpdate {
	su.mutation.ClearMakes()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SESSIONSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SESSIONSUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SESSIONSUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SESSIONSUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SESSIONSUpdate) check() error {
	if v, ok := su.mutation.Status(); ok {
		if err := sessions.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SESSIONS.status": %w`, err)}
		}
	}
	return nil
}

func (su *SESSIONSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(sessions.Table, sessions.Columns, sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UserID(); ok {
		_spec.SetField(sessions.FieldUserID, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedUserID(); ok {
		_spec.AddField(sessions.FieldUserID, field.TypeInt, value)
	}
	if value, ok := su.mutation.MatchedUserID(); ok {
		_spec.SetField(sessions.FieldMatchedUserID, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedMatchedUserID(); ok {
		_spec.AddField(sessions.FieldMatchedUserID, field.TypeInt, value)
	}
	if value, ok := su.mutation.RecordID(); ok {
		_spec.SetField(sessions.FieldRecordID, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedRecordID(); ok {
		_spec.AddField(sessions.FieldRecordID, field.TypeInt, value)
	}
	if value, ok := su.mutation.MatchedAt(); ok {
		_spec.SetField(sessions.FieldMatchedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(sessions.FieldStatus, field.TypeEnum, value)
	}
	if su.mutation.HadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sessions.HadTable,
			Columns: []string{sessions.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.HadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sessions.HadTable,
			Columns: []string{sessions.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.MakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   sessions.MakesTable,
			Columns: []string{sessions.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(calls.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.MakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   sessions.MakesTable,
			Columns: []string{sessions.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(calls.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sessions.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SESSIONSUpdateOne is the builder for updating a single SESSIONS entity.
type SESSIONSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SESSIONSMutation
}

// SetUserID sets the "user_id" field.
func (suo *SESSIONSUpdateOne) SetUserID(i int) *SESSIONSUpdateOne {
	suo.mutation.ResetUserID()
	suo.mutation.SetUserID(i)
	return suo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suo *SESSIONSUpdateOne) SetNillableUserID(i *int) *SESSIONSUpdateOne {
	if i != nil {
		suo.SetUserID(*i)
	}
	return suo
}

// AddUserID adds i to the "user_id" field.
func (suo *SESSIONSUpdateOne) AddUserID(i int) *SESSIONSUpdateOne {
	suo.mutation.AddUserID(i)
	return suo
}

// SetMatchedUserID sets the "matched_user_id" field.
func (suo *SESSIONSUpdateOne) SetMatchedUserID(i int) *SESSIONSUpdateOne {
	suo.mutation.ResetMatchedUserID()
	suo.mutation.SetMatchedUserID(i)
	return suo
}

// SetNillableMatchedUserID sets the "matched_user_id" field if the given value is not nil.
func (suo *SESSIONSUpdateOne) SetNillableMatchedUserID(i *int) *SESSIONSUpdateOne {
	if i != nil {
		suo.SetMatchedUserID(*i)
	}
	return suo
}

// AddMatchedUserID adds i to the "matched_user_id" field.
func (suo *SESSIONSUpdateOne) AddMatchedUserID(i int) *SESSIONSUpdateOne {
	suo.mutation.AddMatchedUserID(i)
	return suo
}

// SetRecordID sets the "record_id" field.
func (suo *SESSIONSUpdateOne) SetRecordID(i int) *SESSIONSUpdateOne {
	suo.mutation.ResetRecordID()
	suo.mutation.SetRecordID(i)
	return suo
}

// SetNillableRecordID sets the "record_id" field if the given value is not nil.
func (suo *SESSIONSUpdateOne) SetNillableRecordID(i *int) *SESSIONSUpdateOne {
	if i != nil {
		suo.SetRecordID(*i)
	}
	return suo
}

// AddRecordID adds i to the "record_id" field.
func (suo *SESSIONSUpdateOne) AddRecordID(i int) *SESSIONSUpdateOne {
	suo.mutation.AddRecordID(i)
	return suo
}

// SetMatchedAt sets the "matched_at" field.
func (suo *SESSIONSUpdateOne) SetMatchedAt(t time.Time) *SESSIONSUpdateOne {
	suo.mutation.SetMatchedAt(t)
	return suo
}

// SetNillableMatchedAt sets the "matched_at" field if the given value is not nil.
func (suo *SESSIONSUpdateOne) SetNillableMatchedAt(t *time.Time) *SESSIONSUpdateOne {
	if t != nil {
		suo.SetMatchedAt(*t)
	}
	return suo
}

// SetStatus sets the "status" field.
func (suo *SESSIONSUpdateOne) SetStatus(s sessions.Status) *SESSIONSUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (suo *SESSIONSUpdateOne) SetNillableStatus(s *sessions.Status) *SESSIONSUpdateOne {
	if s != nil {
		suo.SetStatus(*s)
	}
	return suo
}

// SetHadID sets the "had" edge to the EVENT_RECORDS entity by ID.
func (suo *SESSIONSUpdateOne) SetHadID(id int) *SESSIONSUpdateOne {
	suo.mutation.SetHadID(id)
	return suo
}

// SetNillableHadID sets the "had" edge to the EVENT_RECORDS entity by ID if the given value is not nil.
func (suo *SESSIONSUpdateOne) SetNillableHadID(id *int) *SESSIONSUpdateOne {
	if id != nil {
		suo = suo.SetHadID(*id)
	}
	return suo
}

// SetHad sets the "had" edge to the EVENT_RECORDS entity.
func (suo *SESSIONSUpdateOne) SetHad(e *EVENT_RECORDS) *SESSIONSUpdateOne {
	return suo.SetHadID(e.ID)
}

// SetMakesID sets the "makes" edge to the CALLS entity by ID.
func (suo *SESSIONSUpdateOne) SetMakesID(id int) *SESSIONSUpdateOne {
	suo.mutation.SetMakesID(id)
	return suo
}

// SetNillableMakesID sets the "makes" edge to the CALLS entity by ID if the given value is not nil.
func (suo *SESSIONSUpdateOne) SetNillableMakesID(id *int) *SESSIONSUpdateOne {
	if id != nil {
		suo = suo.SetMakesID(*id)
	}
	return suo
}

// SetMakes sets the "makes" edge to the CALLS entity.
func (suo *SESSIONSUpdateOne) SetMakes(c *CALLS) *SESSIONSUpdateOne {
	return suo.SetMakesID(c.ID)
}

// Mutation returns the SESSIONSMutation object of the builder.
func (suo *SESSIONSUpdateOne) Mutation() *SESSIONSMutation {
	return suo.mutation
}

// ClearHad clears the "had" edge to the EVENT_RECORDS entity.
func (suo *SESSIONSUpdateOne) ClearHad() *SESSIONSUpdateOne {
	suo.mutation.ClearHad()
	return suo
}

// ClearMakes clears the "makes" edge to the CALLS entity.
func (suo *SESSIONSUpdateOne) ClearMakes() *SESSIONSUpdateOne {
	suo.mutation.ClearMakes()
	return suo
}

// Where appends a list predicates to the SESSIONSUpdate builder.
func (suo *SESSIONSUpdateOne) Where(ps ...predicate.SESSIONS) *SESSIONSUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SESSIONSUpdateOne) Select(field string, fields ...string) *SESSIONSUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated SESSIONS entity.
func (suo *SESSIONSUpdateOne) Save(ctx context.Context) (*SESSIONS, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SESSIONSUpdateOne) SaveX(ctx context.Context) *SESSIONS {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SESSIONSUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SESSIONSUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SESSIONSUpdateOne) check() error {
	if v, ok := suo.mutation.Status(); ok {
		if err := sessions.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "SESSIONS.status": %w`, err)}
		}
	}
	return nil
}

func (suo *SESSIONSUpdateOne) sqlSave(ctx context.Context) (_node *SESSIONS, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(sessions.Table, sessions.Columns, sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SESSIONS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sessions.FieldID)
		for _, f := range fields {
			if !sessions.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sessions.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UserID(); ok {
		_spec.SetField(sessions.FieldUserID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedUserID(); ok {
		_spec.AddField(sessions.FieldUserID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.MatchedUserID(); ok {
		_spec.SetField(sessions.FieldMatchedUserID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedMatchedUserID(); ok {
		_spec.AddField(sessions.FieldMatchedUserID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.RecordID(); ok {
		_spec.SetField(sessions.FieldRecordID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedRecordID(); ok {
		_spec.AddField(sessions.FieldRecordID, field.TypeInt, value)
	}
	if value, ok := suo.mutation.MatchedAt(); ok {
		_spec.SetField(sessions.FieldMatchedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(sessions.FieldStatus, field.TypeEnum, value)
	}
	if suo.mutation.HadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sessions.HadTable,
			Columns: []string{sessions.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.HadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   sessions.HadTable,
			Columns: []string{sessions.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.MakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   sessions.MakesTable,
			Columns: []string{sessions.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(calls.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.MakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   sessions.MakesTable,
			Columns: []string{sessions.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(calls.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SESSIONS{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sessions.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
