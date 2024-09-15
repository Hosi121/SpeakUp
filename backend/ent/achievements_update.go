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
	"github.com/Hosi121/SpeakUp/ent/achievements"
	"github.com/Hosi121/SpeakUp/ent/predicate"
	"github.com/Hosi121/SpeakUp/ent/trophies"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// ACHIEVEMENTSUpdate is the builder for updating ACHIEVEMENTS entities.
type ACHIEVEMENTSUpdate struct {
	config
	hooks    []Hook
	mutation *ACHIEVEMENTSMutation
}

// Where appends a list predicates to the ACHIEVEMENTSUpdate builder.
func (au *ACHIEVEMENTSUpdate) Where(ps ...predicate.ACHIEVEMENTS) *ACHIEVEMENTSUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUserID sets the "user_id" field.
func (au *ACHIEVEMENTSUpdate) SetUserID(i int) *ACHIEVEMENTSUpdate {
	au.mutation.ResetUserID()
	au.mutation.SetUserID(i)
	return au
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (au *ACHIEVEMENTSUpdate) SetNillableUserID(i *int) *ACHIEVEMENTSUpdate {
	if i != nil {
		au.SetUserID(*i)
	}
	return au
}

// AddUserID adds i to the "user_id" field.
func (au *ACHIEVEMENTSUpdate) AddUserID(i int) *ACHIEVEMENTSUpdate {
	au.mutation.AddUserID(i)
	return au
}

// SetTrophyID sets the "trophy_id" field.
func (au *ACHIEVEMENTSUpdate) SetTrophyID(i int) *ACHIEVEMENTSUpdate {
	au.mutation.ResetTrophyID()
	au.mutation.SetTrophyID(i)
	return au
}

// SetNillableTrophyID sets the "trophy_id" field if the given value is not nil.
func (au *ACHIEVEMENTSUpdate) SetNillableTrophyID(i *int) *ACHIEVEMENTSUpdate {
	if i != nil {
		au.SetTrophyID(*i)
	}
	return au
}

// AddTrophyID adds i to the "trophy_id" field.
func (au *ACHIEVEMENTSUpdate) AddTrophyID(i int) *ACHIEVEMENTSUpdate {
	au.mutation.AddTrophyID(i)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *ACHIEVEMENTSUpdate) SetCreatedAt(t time.Time) *ACHIEVEMENTSUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *ACHIEVEMENTSUpdate) SetNillableCreatedAt(t *time.Time) *ACHIEVEMENTSUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetGrantedID sets the "granted" edge to the USERS entity by ID.
func (au *ACHIEVEMENTSUpdate) SetGrantedID(id int) *ACHIEVEMENTSUpdate {
	au.mutation.SetGrantedID(id)
	return au
}

// SetNillableGrantedID sets the "granted" edge to the USERS entity by ID if the given value is not nil.
func (au *ACHIEVEMENTSUpdate) SetNillableGrantedID(id *int) *ACHIEVEMENTSUpdate {
	if id != nil {
		au = au.SetGrantedID(*id)
	}
	return au
}

// SetGranted sets the "granted" edge to the USERS entity.
func (au *ACHIEVEMENTSUpdate) SetGranted(u *USERS) *ACHIEVEMENTSUpdate {
	return au.SetGrantedID(u.ID)
}

// SetRefersID sets the "refers" edge to the TROPHIES entity by ID.
func (au *ACHIEVEMENTSUpdate) SetRefersID(id int) *ACHIEVEMENTSUpdate {
	au.mutation.SetRefersID(id)
	return au
}

// SetNillableRefersID sets the "refers" edge to the TROPHIES entity by ID if the given value is not nil.
func (au *ACHIEVEMENTSUpdate) SetNillableRefersID(id *int) *ACHIEVEMENTSUpdate {
	if id != nil {
		au = au.SetRefersID(*id)
	}
	return au
}

// SetRefers sets the "refers" edge to the TROPHIES entity.
func (au *ACHIEVEMENTSUpdate) SetRefers(t *TROPHIES) *ACHIEVEMENTSUpdate {
	return au.SetRefersID(t.ID)
}

// Mutation returns the ACHIEVEMENTSMutation object of the builder.
func (au *ACHIEVEMENTSUpdate) Mutation() *ACHIEVEMENTSMutation {
	return au.mutation
}

// ClearGranted clears the "granted" edge to the USERS entity.
func (au *ACHIEVEMENTSUpdate) ClearGranted() *ACHIEVEMENTSUpdate {
	au.mutation.ClearGranted()
	return au
}

// ClearRefers clears the "refers" edge to the TROPHIES entity.
func (au *ACHIEVEMENTSUpdate) ClearRefers() *ACHIEVEMENTSUpdate {
	au.mutation.ClearRefers()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ACHIEVEMENTSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ACHIEVEMENTSUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ACHIEVEMENTSUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ACHIEVEMENTSUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ACHIEVEMENTSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(achievements.Table, achievements.Columns, sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UserID(); ok {
		_spec.SetField(achievements.FieldUserID, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedUserID(); ok {
		_spec.AddField(achievements.FieldUserID, field.TypeInt, value)
	}
	if value, ok := au.mutation.TrophyID(); ok {
		_spec.SetField(achievements.FieldTrophyID, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedTrophyID(); ok {
		_spec.AddField(achievements.FieldTrophyID, field.TypeInt, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(achievements.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.GrantedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   achievements.GrantedTable,
			Columns: []string{achievements.GrantedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.GrantedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   achievements.GrantedTable,
			Columns: []string{achievements.GrantedColumn},
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
	if au.mutation.RefersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   achievements.RefersTable,
			Columns: []string{achievements.RefersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trophies.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RefersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   achievements.RefersTable,
			Columns: []string{achievements.RefersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trophies.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievements.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ACHIEVEMENTSUpdateOne is the builder for updating a single ACHIEVEMENTS entity.
type ACHIEVEMENTSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ACHIEVEMENTSMutation
}

// SetUserID sets the "user_id" field.
func (auo *ACHIEVEMENTSUpdateOne) SetUserID(i int) *ACHIEVEMENTSUpdateOne {
	auo.mutation.ResetUserID()
	auo.mutation.SetUserID(i)
	return auo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (auo *ACHIEVEMENTSUpdateOne) SetNillableUserID(i *int) *ACHIEVEMENTSUpdateOne {
	if i != nil {
		auo.SetUserID(*i)
	}
	return auo
}

// AddUserID adds i to the "user_id" field.
func (auo *ACHIEVEMENTSUpdateOne) AddUserID(i int) *ACHIEVEMENTSUpdateOne {
	auo.mutation.AddUserID(i)
	return auo
}

// SetTrophyID sets the "trophy_id" field.
func (auo *ACHIEVEMENTSUpdateOne) SetTrophyID(i int) *ACHIEVEMENTSUpdateOne {
	auo.mutation.ResetTrophyID()
	auo.mutation.SetTrophyID(i)
	return auo
}

// SetNillableTrophyID sets the "trophy_id" field if the given value is not nil.
func (auo *ACHIEVEMENTSUpdateOne) SetNillableTrophyID(i *int) *ACHIEVEMENTSUpdateOne {
	if i != nil {
		auo.SetTrophyID(*i)
	}
	return auo
}

// AddTrophyID adds i to the "trophy_id" field.
func (auo *ACHIEVEMENTSUpdateOne) AddTrophyID(i int) *ACHIEVEMENTSUpdateOne {
	auo.mutation.AddTrophyID(i)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *ACHIEVEMENTSUpdateOne) SetCreatedAt(t time.Time) *ACHIEVEMENTSUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *ACHIEVEMENTSUpdateOne) SetNillableCreatedAt(t *time.Time) *ACHIEVEMENTSUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetGrantedID sets the "granted" edge to the USERS entity by ID.
func (auo *ACHIEVEMENTSUpdateOne) SetGrantedID(id int) *ACHIEVEMENTSUpdateOne {
	auo.mutation.SetGrantedID(id)
	return auo
}

// SetNillableGrantedID sets the "granted" edge to the USERS entity by ID if the given value is not nil.
func (auo *ACHIEVEMENTSUpdateOne) SetNillableGrantedID(id *int) *ACHIEVEMENTSUpdateOne {
	if id != nil {
		auo = auo.SetGrantedID(*id)
	}
	return auo
}

// SetGranted sets the "granted" edge to the USERS entity.
func (auo *ACHIEVEMENTSUpdateOne) SetGranted(u *USERS) *ACHIEVEMENTSUpdateOne {
	return auo.SetGrantedID(u.ID)
}

// SetRefersID sets the "refers" edge to the TROPHIES entity by ID.
func (auo *ACHIEVEMENTSUpdateOne) SetRefersID(id int) *ACHIEVEMENTSUpdateOne {
	auo.mutation.SetRefersID(id)
	return auo
}

// SetNillableRefersID sets the "refers" edge to the TROPHIES entity by ID if the given value is not nil.
func (auo *ACHIEVEMENTSUpdateOne) SetNillableRefersID(id *int) *ACHIEVEMENTSUpdateOne {
	if id != nil {
		auo = auo.SetRefersID(*id)
	}
	return auo
}

// SetRefers sets the "refers" edge to the TROPHIES entity.
func (auo *ACHIEVEMENTSUpdateOne) SetRefers(t *TROPHIES) *ACHIEVEMENTSUpdateOne {
	return auo.SetRefersID(t.ID)
}

// Mutation returns the ACHIEVEMENTSMutation object of the builder.
func (auo *ACHIEVEMENTSUpdateOne) Mutation() *ACHIEVEMENTSMutation {
	return auo.mutation
}

// ClearGranted clears the "granted" edge to the USERS entity.
func (auo *ACHIEVEMENTSUpdateOne) ClearGranted() *ACHIEVEMENTSUpdateOne {
	auo.mutation.ClearGranted()
	return auo
}

// ClearRefers clears the "refers" edge to the TROPHIES entity.
func (auo *ACHIEVEMENTSUpdateOne) ClearRefers() *ACHIEVEMENTSUpdateOne {
	auo.mutation.ClearRefers()
	return auo
}

// Where appends a list predicates to the ACHIEVEMENTSUpdate builder.
func (auo *ACHIEVEMENTSUpdateOne) Where(ps ...predicate.ACHIEVEMENTS) *ACHIEVEMENTSUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ACHIEVEMENTSUpdateOne) Select(field string, fields ...string) *ACHIEVEMENTSUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated ACHIEVEMENTS entity.
func (auo *ACHIEVEMENTSUpdateOne) Save(ctx context.Context) (*ACHIEVEMENTS, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ACHIEVEMENTSUpdateOne) SaveX(ctx context.Context) *ACHIEVEMENTS {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ACHIEVEMENTSUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ACHIEVEMENTSUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ACHIEVEMENTSUpdateOne) sqlSave(ctx context.Context) (_node *ACHIEVEMENTS, err error) {
	_spec := sqlgraph.NewUpdateSpec(achievements.Table, achievements.Columns, sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ACHIEVEMENTS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, achievements.FieldID)
		for _, f := range fields {
			if !achievements.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != achievements.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UserID(); ok {
		_spec.SetField(achievements.FieldUserID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedUserID(); ok {
		_spec.AddField(achievements.FieldUserID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.TrophyID(); ok {
		_spec.SetField(achievements.FieldTrophyID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedTrophyID(); ok {
		_spec.AddField(achievements.FieldTrophyID, field.TypeInt, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(achievements.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.GrantedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   achievements.GrantedTable,
			Columns: []string{achievements.GrantedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.GrantedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   achievements.GrantedTable,
			Columns: []string{achievements.GrantedColumn},
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
	if auo.mutation.RefersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   achievements.RefersTable,
			Columns: []string{achievements.RefersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trophies.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RefersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   achievements.RefersTable,
			Columns: []string{achievements.RefersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trophies.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ACHIEVEMENTS{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievements.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
