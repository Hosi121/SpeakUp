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
	"github.com/Hosi121/SpeakUp/ent/chats"
	"github.com/Hosi121/SpeakUp/ent/friends"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// CHATSUpdate is the builder for updating CHATS entities.
type CHATSUpdate struct {
	config
	hooks    []Hook
	mutation *CHATSMutation
}

// Where appends a list predicates to the CHATSUpdate builder.
func (cu *CHATSUpdate) Where(ps ...predicate.CHATS) *CHATSUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetFriendID sets the "friend_id" field.
func (cu *CHATSUpdate) SetFriendID(i int) *CHATSUpdate {
	cu.mutation.ResetFriendID()
	cu.mutation.SetFriendID(i)
	return cu
}

// SetNillableFriendID sets the "friend_id" field if the given value is not nil.
func (cu *CHATSUpdate) SetNillableFriendID(i *int) *CHATSUpdate {
	if i != nil {
		cu.SetFriendID(*i)
	}
	return cu
}

// AddFriendID adds i to the "friend_id" field.
func (cu *CHATSUpdate) AddFriendID(i int) *CHATSUpdate {
	cu.mutation.AddFriendID(i)
	return cu
}

// SetMessage sets the "message" field.
func (cu *CHATSUpdate) SetMessage(s string) *CHATSUpdate {
	cu.mutation.SetMessage(s)
	return cu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cu *CHATSUpdate) SetNillableMessage(s *string) *CHATSUpdate {
	if s != nil {
		cu.SetMessage(*s)
	}
	return cu
}

// SetIsRecieved sets the "is_recieved" field.
func (cu *CHATSUpdate) SetIsRecieved(b bool) *CHATSUpdate {
	cu.mutation.SetIsRecieved(b)
	return cu
}

// SetNillableIsRecieved sets the "is_recieved" field if the given value is not nil.
func (cu *CHATSUpdate) SetNillableIsRecieved(b *bool) *CHATSUpdate {
	if b != nil {
		cu.SetIsRecieved(*b)
	}
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CHATSUpdate) SetCreatedAt(t time.Time) *CHATSUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CHATSUpdate) SetNillableCreatedAt(t *time.Time) *CHATSUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetHadID sets the "had" edge to the FRIENDS entity by ID.
func (cu *CHATSUpdate) SetHadID(id int) *CHATSUpdate {
	cu.mutation.SetHadID(id)
	return cu
}

// SetNillableHadID sets the "had" edge to the FRIENDS entity by ID if the given value is not nil.
func (cu *CHATSUpdate) SetNillableHadID(id *int) *CHATSUpdate {
	if id != nil {
		cu = cu.SetHadID(*id)
	}
	return cu
}

// SetHad sets the "had" edge to the FRIENDS entity.
func (cu *CHATSUpdate) SetHad(f *FRIENDS) *CHATSUpdate {
	return cu.SetHadID(f.ID)
}

// Mutation returns the CHATSMutation object of the builder.
func (cu *CHATSUpdate) Mutation() *CHATSMutation {
	return cu.mutation
}

// ClearHad clears the "had" edge to the FRIENDS entity.
func (cu *CHATSUpdate) ClearHad() *CHATSUpdate {
	cu.mutation.ClearHad()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CHATSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CHATSUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CHATSUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CHATSUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CHATSUpdate) check() error {
	if v, ok := cu.mutation.Message(); ok {
		if err := chats.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "CHATS.message": %w`, err)}
		}
	}
	return nil
}

func (cu *CHATSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(chats.Table, chats.Columns, sqlgraph.NewFieldSpec(chats.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.FriendID(); ok {
		_spec.SetField(chats.FieldFriendID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedFriendID(); ok {
		_spec.AddField(chats.FieldFriendID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Message(); ok {
		_spec.SetField(chats.FieldMessage, field.TypeString, value)
	}
	if value, ok := cu.mutation.IsRecieved(); ok {
		_spec.SetField(chats.FieldIsRecieved, field.TypeBool, value)
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.SetField(chats.FieldCreatedAt, field.TypeTime, value)
	}
	if cu.mutation.HadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chats.HadTable,
			Columns: []string{chats.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friends.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.HadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chats.HadTable,
			Columns: []string{chats.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friends.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CHATSUpdateOne is the builder for updating a single CHATS entity.
type CHATSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CHATSMutation
}

// SetFriendID sets the "friend_id" field.
func (cuo *CHATSUpdateOne) SetFriendID(i int) *CHATSUpdateOne {
	cuo.mutation.ResetFriendID()
	cuo.mutation.SetFriendID(i)
	return cuo
}

// SetNillableFriendID sets the "friend_id" field if the given value is not nil.
func (cuo *CHATSUpdateOne) SetNillableFriendID(i *int) *CHATSUpdateOne {
	if i != nil {
		cuo.SetFriendID(*i)
	}
	return cuo
}

// AddFriendID adds i to the "friend_id" field.
func (cuo *CHATSUpdateOne) AddFriendID(i int) *CHATSUpdateOne {
	cuo.mutation.AddFriendID(i)
	return cuo
}

// SetMessage sets the "message" field.
func (cuo *CHATSUpdateOne) SetMessage(s string) *CHATSUpdateOne {
	cuo.mutation.SetMessage(s)
	return cuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cuo *CHATSUpdateOne) SetNillableMessage(s *string) *CHATSUpdateOne {
	if s != nil {
		cuo.SetMessage(*s)
	}
	return cuo
}

// SetIsRecieved sets the "is_recieved" field.
func (cuo *CHATSUpdateOne) SetIsRecieved(b bool) *CHATSUpdateOne {
	cuo.mutation.SetIsRecieved(b)
	return cuo
}

// SetNillableIsRecieved sets the "is_recieved" field if the given value is not nil.
func (cuo *CHATSUpdateOne) SetNillableIsRecieved(b *bool) *CHATSUpdateOne {
	if b != nil {
		cuo.SetIsRecieved(*b)
	}
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CHATSUpdateOne) SetCreatedAt(t time.Time) *CHATSUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CHATSUpdateOne) SetNillableCreatedAt(t *time.Time) *CHATSUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetHadID sets the "had" edge to the FRIENDS entity by ID.
func (cuo *CHATSUpdateOne) SetHadID(id int) *CHATSUpdateOne {
	cuo.mutation.SetHadID(id)
	return cuo
}

// SetNillableHadID sets the "had" edge to the FRIENDS entity by ID if the given value is not nil.
func (cuo *CHATSUpdateOne) SetNillableHadID(id *int) *CHATSUpdateOne {
	if id != nil {
		cuo = cuo.SetHadID(*id)
	}
	return cuo
}

// SetHad sets the "had" edge to the FRIENDS entity.
func (cuo *CHATSUpdateOne) SetHad(f *FRIENDS) *CHATSUpdateOne {
	return cuo.SetHadID(f.ID)
}

// Mutation returns the CHATSMutation object of the builder.
func (cuo *CHATSUpdateOne) Mutation() *CHATSMutation {
	return cuo.mutation
}

// ClearHad clears the "had" edge to the FRIENDS entity.
func (cuo *CHATSUpdateOne) ClearHad() *CHATSUpdateOne {
	cuo.mutation.ClearHad()
	return cuo
}

// Where appends a list predicates to the CHATSUpdate builder.
func (cuo *CHATSUpdateOne) Where(ps ...predicate.CHATS) *CHATSUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CHATSUpdateOne) Select(field string, fields ...string) *CHATSUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated CHATS entity.
func (cuo *CHATSUpdateOne) Save(ctx context.Context) (*CHATS, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CHATSUpdateOne) SaveX(ctx context.Context) *CHATS {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CHATSUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CHATSUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CHATSUpdateOne) check() error {
	if v, ok := cuo.mutation.Message(); ok {
		if err := chats.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "CHATS.message": %w`, err)}
		}
	}
	return nil
}

func (cuo *CHATSUpdateOne) sqlSave(ctx context.Context) (_node *CHATS, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(chats.Table, chats.Columns, sqlgraph.NewFieldSpec(chats.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CHATS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chats.FieldID)
		for _, f := range fields {
			if !chats.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chats.FieldID {
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
	if value, ok := cuo.mutation.FriendID(); ok {
		_spec.SetField(chats.FieldFriendID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedFriendID(); ok {
		_spec.AddField(chats.FieldFriendID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Message(); ok {
		_spec.SetField(chats.FieldMessage, field.TypeString, value)
	}
	if value, ok := cuo.mutation.IsRecieved(); ok {
		_spec.SetField(chats.FieldIsRecieved, field.TypeBool, value)
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.SetField(chats.FieldCreatedAt, field.TypeTime, value)
	}
	if cuo.mutation.HadCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chats.HadTable,
			Columns: []string{chats.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friends.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.HadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chats.HadTable,
			Columns: []string{chats.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friends.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CHATS{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}