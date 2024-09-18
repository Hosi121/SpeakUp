// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/events"
	"github.com/Hosi121/SpeakUp/ent/sessions"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// EVENTRECORDSCreate is the builder for creating a EVENT_RECORDS entity.
type EVENTRECORDSCreate struct {
	config
	mutation *EVENTRECORDSMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ec *EVENTRECORDSCreate) SetUserID(i int) *EVENTRECORDSCreate {
	ec.mutation.SetUserID(i)
	return ec
}

// SetEventID sets the "event_id" field.
func (ec *EVENTRECORDSCreate) SetEventID(i int) *EVENTRECORDSCreate {
	ec.mutation.SetEventID(i)
	return ec
}

// SetParticipatesBit sets the "participates_bit" field.
func (ec *EVENTRECORDSCreate) SetParticipatesBit(i int) *EVENTRECORDSCreate {
	ec.mutation.SetParticipatesBit(i)
	return ec
}

// SetNillableParticipatesBit sets the "participates_bit" field if the given value is not nil.
func (ec *EVENTRECORDSCreate) SetNillableParticipatesBit(i *int) *EVENTRECORDSCreate {
	if i != nil {
		ec.SetParticipatesBit(*i)
	}
	return ec
}

// SetRecords sets the "records" field.
func (ec *EVENTRECORDSCreate) SetRecords(s string) *EVENTRECORDSCreate {
	ec.mutation.SetRecords(s)
	return ec
}

// SetNillableRecords sets the "records" field if the given value is not nil.
func (ec *EVENTRECORDSCreate) SetNillableRecords(s *string) *EVENTRECORDSCreate {
	if s != nil {
		ec.SetRecords(*s)
	}
	return ec
}

// SetMadeID sets the "made" edge to the USERS entity by ID.
func (ec *EVENTRECORDSCreate) SetMadeID(id int) *EVENTRECORDSCreate {
	ec.mutation.SetMadeID(id)
	return ec
}

// SetNillableMadeID sets the "made" edge to the USERS entity by ID if the given value is not nil.
func (ec *EVENTRECORDSCreate) SetNillableMadeID(id *int) *EVENTRECORDSCreate {
	if id != nil {
		ec = ec.SetMadeID(*id)
	}
	return ec
}

// SetMade sets the "made" edge to the USERS entity.
func (ec *EVENTRECORDSCreate) SetMade(u *USERS) *EVENTRECORDSCreate {
	return ec.SetMadeID(u.ID)
}

// SetParticipatesID sets the "participates" edge to the EVENTS entity by ID.
func (ec *EVENTRECORDSCreate) SetParticipatesID(id int) *EVENTRECORDSCreate {
	ec.mutation.SetParticipatesID(id)
	return ec
}

// SetNillableParticipatesID sets the "participates" edge to the EVENTS entity by ID if the given value is not nil.
func (ec *EVENTRECORDSCreate) SetNillableParticipatesID(id *int) *EVENTRECORDSCreate {
	if id != nil {
		ec = ec.SetParticipatesID(*id)
	}
	return ec
}

// SetParticipates sets the "participates" edge to the EVENTS entity.
func (ec *EVENTRECORDSCreate) SetParticipates(e *EVENTS) *EVENTRECORDSCreate {
	return ec.SetParticipatesID(e.ID)
}

// AddHaIDs adds the "has" edge to the SESSIONS entity by IDs.
func (ec *EVENTRECORDSCreate) AddHaIDs(ids ...int) *EVENTRECORDSCreate {
	ec.mutation.AddHaIDs(ids...)
	return ec
}

// AddHas adds the "has" edges to the SESSIONS entity.
func (ec *EVENTRECORDSCreate) AddHas(s ...*SESSIONS) *EVENTRECORDSCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ec.AddHaIDs(ids...)
}

// Mutation returns the EVENTRECORDSMutation object of the builder.
func (ec *EVENTRECORDSCreate) Mutation() *EVENTRECORDSMutation {
	return ec.mutation
}

// Save creates the EVENT_RECORDS in the database.
func (ec *EVENTRECORDSCreate) Save(ctx context.Context) (*EVENT_RECORDS, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EVENTRECORDSCreate) SaveX(ctx context.Context) *EVENT_RECORDS {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EVENTRECORDSCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EVENTRECORDSCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EVENTRECORDSCreate) defaults() {
	if _, ok := ec.mutation.ParticipatesBit(); !ok {
		v := event_records.DefaultParticipatesBit
		ec.mutation.SetParticipatesBit(v)
	}
	if _, ok := ec.mutation.Records(); !ok {
		v := event_records.DefaultRecords
		ec.mutation.SetRecords(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EVENTRECORDSCreate) check() error {
	if _, ok := ec.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "EVENT_RECORDS.user_id"`)}
	}
	if _, ok := ec.mutation.EventID(); !ok {
		return &ValidationError{Name: "event_id", err: errors.New(`ent: missing required field "EVENT_RECORDS.event_id"`)}
	}
	if _, ok := ec.mutation.ParticipatesBit(); !ok {
		return &ValidationError{Name: "participates_bit", err: errors.New(`ent: missing required field "EVENT_RECORDS.participates_bit"`)}
	}
	if v, ok := ec.mutation.ParticipatesBit(); ok {
		if err := event_records.ParticipatesBitValidator(v); err != nil {
			return &ValidationError{Name: "participates_bit", err: fmt.Errorf(`ent: validator failed for field "EVENT_RECORDS.participates_bit": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Records(); !ok {
		return &ValidationError{Name: "records", err: errors.New(`ent: missing required field "EVENT_RECORDS.records"`)}
	}
	return nil
}

func (ec *EVENTRECORDSCreate) sqlSave(ctx context.Context) (*EVENT_RECORDS, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EVENTRECORDSCreate) createSpec() (*EVENT_RECORDS, *sqlgraph.CreateSpec) {
	var (
		_node = &EVENT_RECORDS{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(event_records.Table, sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.UserID(); ok {
		_spec.SetField(event_records.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := ec.mutation.EventID(); ok {
		_spec.SetField(event_records.FieldEventID, field.TypeInt, value)
		_node.EventID = value
	}
	if value, ok := ec.mutation.ParticipatesBit(); ok {
		_spec.SetField(event_records.FieldParticipatesBit, field.TypeInt, value)
		_node.ParticipatesBit = value
	}
	if value, ok := ec.mutation.Records(); ok {
		_spec.SetField(event_records.FieldRecords, field.TypeString, value)
		_node.Records = value
	}
	if nodes := ec.mutation.MadeIDs(); len(nodes) > 0 {
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
		_node.users_makes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ParticipatesIDs(); len(nodes) > 0 {
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
		_node.events_participated = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.HasIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EVENTRECORDSCreateBulk is the builder for creating many EVENT_RECORDS entities in bulk.
type EVENTRECORDSCreateBulk struct {
	config
	err      error
	builders []*EVENTRECORDSCreate
}

// Save creates the EVENT_RECORDS entities in the database.
func (ecb *EVENTRECORDSCreateBulk) Save(ctx context.Context) ([]*EVENT_RECORDS, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*EVENT_RECORDS, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EVENTRECORDSMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EVENTRECORDSCreateBulk) SaveX(ctx context.Context) []*EVENT_RECORDS {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EVENTRECORDSCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EVENTRECORDSCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
