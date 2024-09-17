// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/progress"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// PROGRESSCreate is the builder for creating a PROGRESS entity.
type PROGRESSCreate struct {
	config
	mutation *PROGRESSMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (pc *PROGRESSCreate) SetUserID(i int) *PROGRESSCreate {
	pc.mutation.SetUserID(i)
	return pc
}

// SetLoginDays sets the "login_days" field.
func (pc *PROGRESSCreate) SetLoginDays(i int) *PROGRESSCreate {
	pc.mutation.SetLoginDays(i)
	return pc
}

// SetNillableLoginDays sets the "login_days" field if the given value is not nil.
func (pc *PROGRESSCreate) SetNillableLoginDays(i *int) *PROGRESSCreate {
	if i != nil {
		pc.SetLoginDays(*i)
	}
	return pc
}

// SetConsecutiveParticipants sets the "consecutive_participants" field.
func (pc *PROGRESSCreate) SetConsecutiveParticipants(i int) *PROGRESSCreate {
	pc.mutation.SetConsecutiveParticipants(i)
	return pc
}

// SetNillableConsecutiveParticipants sets the "consecutive_participants" field if the given value is not nil.
func (pc *PROGRESSCreate) SetNillableConsecutiveParticipants(i *int) *PROGRESSCreate {
	if i != nil {
		pc.SetConsecutiveParticipants(*i)
	}
	return pc
}

// SetConsecutiveLoginDays sets the "consecutive_login_days" field.
func (pc *PROGRESSCreate) SetConsecutiveLoginDays(i int) *PROGRESSCreate {
	pc.mutation.SetConsecutiveLoginDays(i)
	return pc
}

// SetNillableConsecutiveLoginDays sets the "consecutive_login_days" field if the given value is not nil.
func (pc *PROGRESSCreate) SetNillableConsecutiveLoginDays(i *int) *PROGRESSCreate {
	if i != nil {
		pc.SetConsecutiveLoginDays(*i)
	}
	return pc
}

// SetRecordedID sets the "recorded" edge to the USERS entity by ID.
func (pc *PROGRESSCreate) SetRecordedID(id int) *PROGRESSCreate {
	pc.mutation.SetRecordedID(id)
	return pc
}

// SetNillableRecordedID sets the "recorded" edge to the USERS entity by ID if the given value is not nil.
func (pc *PROGRESSCreate) SetNillableRecordedID(id *int) *PROGRESSCreate {
	if id != nil {
		pc = pc.SetRecordedID(*id)
	}
	return pc
}

// SetRecorded sets the "recorded" edge to the USERS entity.
func (pc *PROGRESSCreate) SetRecorded(u *USERS) *PROGRESSCreate {
	return pc.SetRecordedID(u.ID)
}

// Mutation returns the PROGRESSMutation object of the builder.
func (pc *PROGRESSCreate) Mutation() *PROGRESSMutation {
	return pc.mutation
}

// Save creates the PROGRESS in the database.
func (pc *PROGRESSCreate) Save(ctx context.Context) (*PROGRESS, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PROGRESSCreate) SaveX(ctx context.Context) *PROGRESS {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PROGRESSCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PROGRESSCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PROGRESSCreate) defaults() {
	if _, ok := pc.mutation.LoginDays(); !ok {
		v := progress.DefaultLoginDays
		pc.mutation.SetLoginDays(v)
	}
	if _, ok := pc.mutation.ConsecutiveParticipants(); !ok {
		v := progress.DefaultConsecutiveParticipants
		pc.mutation.SetConsecutiveParticipants(v)
	}
	if _, ok := pc.mutation.ConsecutiveLoginDays(); !ok {
		v := progress.DefaultConsecutiveLoginDays
		pc.mutation.SetConsecutiveLoginDays(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PROGRESSCreate) check() error {
	if _, ok := pc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "PROGRESS.user_id"`)}
	}
	if _, ok := pc.mutation.LoginDays(); !ok {
		return &ValidationError{Name: "login_days", err: errors.New(`ent: missing required field "PROGRESS.login_days"`)}
	}
	if _, ok := pc.mutation.ConsecutiveParticipants(); !ok {
		return &ValidationError{Name: "consecutive_participants", err: errors.New(`ent: missing required field "PROGRESS.consecutive_participants"`)}
	}
	if _, ok := pc.mutation.ConsecutiveLoginDays(); !ok {
		return &ValidationError{Name: "consecutive_login_days", err: errors.New(`ent: missing required field "PROGRESS.consecutive_login_days"`)}
	}
	return nil
}

func (pc *PROGRESSCreate) sqlSave(ctx context.Context) (*PROGRESS, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PROGRESSCreate) createSpec() (*PROGRESS, *sqlgraph.CreateSpec) {
	var (
		_node = &PROGRESS{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(progress.Table, sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.UserID(); ok {
		_spec.SetField(progress.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := pc.mutation.LoginDays(); ok {
		_spec.SetField(progress.FieldLoginDays, field.TypeInt, value)
		_node.LoginDays = value
	}
	if value, ok := pc.mutation.ConsecutiveParticipants(); ok {
		_spec.SetField(progress.FieldConsecutiveParticipants, field.TypeInt, value)
		_node.ConsecutiveParticipants = value
	}
	if value, ok := pc.mutation.ConsecutiveLoginDays(); ok {
		_spec.SetField(progress.FieldConsecutiveLoginDays, field.TypeInt, value)
		_node.ConsecutiveLoginDays = value
	}
	if nodes := pc.mutation.RecordedIDs(); len(nodes) > 0 {
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
		_node.users_records = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PROGRESSCreateBulk is the builder for creating many PROGRESS entities in bulk.
type PROGRESSCreateBulk struct {
	config
	err      error
	builders []*PROGRESSCreate
}

// Save creates the PROGRESS entities in the database.
func (pcb *PROGRESSCreateBulk) Save(ctx context.Context) ([]*PROGRESS, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*PROGRESS, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PROGRESSMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PROGRESSCreateBulk) SaveX(ctx context.Context) []*PROGRESS {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PROGRESSCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PROGRESSCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
