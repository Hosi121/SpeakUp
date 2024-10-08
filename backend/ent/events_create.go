// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/ai_themes"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/events"
)

// EVENTSCreate is the builder for creating a EVENTS entity.
type EVENTSCreate struct {
	config
	mutation *EVENTSMutation
	hooks    []Hook
}

// SetEventStart sets the "event_start" field.
func (ec *EVENTSCreate) SetEventStart(t time.Time) *EVENTSCreate {
	ec.mutation.SetEventStart(t)
	return ec
}

// SetEventEnd sets the "event_end" field.
func (ec *EVENTSCreate) SetEventEnd(t time.Time) *EVENTSCreate {
	ec.mutation.SetEventEnd(t)
	return ec
}

// SetThemeID sets the "theme_id" field.
func (ec *EVENTSCreate) SetThemeID(i int) *EVENTSCreate {
	ec.mutation.SetThemeID(i)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EVENTSCreate) SetCreatedAt(t time.Time) *EVENTSCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EVENTSCreate) SetNillableCreatedAt(t *time.Time) *EVENTSCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// AddParticipatedIDs adds the "participated" edge to the EVENT_RECORDS entity by IDs.
func (ec *EVENTSCreate) AddParticipatedIDs(ids ...int) *EVENTSCreate {
	ec.mutation.AddParticipatedIDs(ids...)
	return ec
}

// AddParticipated adds the "participated" edges to the EVENT_RECORDS entity.
func (ec *EVENTSCreate) AddParticipated(e ...*EVENT_RECORDS) *EVENTSCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddParticipatedIDs(ids...)
}

// SetUsesID sets the "uses" edge to the AI_THEMES entity by ID.
func (ec *EVENTSCreate) SetUsesID(id int) *EVENTSCreate {
	ec.mutation.SetUsesID(id)
	return ec
}

// SetNillableUsesID sets the "uses" edge to the AI_THEMES entity by ID if the given value is not nil.
func (ec *EVENTSCreate) SetNillableUsesID(id *int) *EVENTSCreate {
	if id != nil {
		ec = ec.SetUsesID(*id)
	}
	return ec
}

// SetUses sets the "uses" edge to the AI_THEMES entity.
func (ec *EVENTSCreate) SetUses(a *AI_THEMES) *EVENTSCreate {
	return ec.SetUsesID(a.ID)
}

// Mutation returns the EVENTSMutation object of the builder.
func (ec *EVENTSCreate) Mutation() *EVENTSMutation {
	return ec.mutation
}

// Save creates the EVENTS in the database.
func (ec *EVENTSCreate) Save(ctx context.Context) (*EVENTS, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EVENTSCreate) SaveX(ctx context.Context) *EVENTS {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EVENTSCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EVENTSCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EVENTSCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := events.DefaultCreatedAt
		ec.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EVENTSCreate) check() error {
	if _, ok := ec.mutation.EventStart(); !ok {
		return &ValidationError{Name: "event_start", err: errors.New(`ent: missing required field "EVENTS.event_start"`)}
	}
	if _, ok := ec.mutation.EventEnd(); !ok {
		return &ValidationError{Name: "event_end", err: errors.New(`ent: missing required field "EVENTS.event_end"`)}
	}
	if _, ok := ec.mutation.ThemeID(); !ok {
		return &ValidationError{Name: "theme_id", err: errors.New(`ent: missing required field "EVENTS.theme_id"`)}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EVENTS.created_at"`)}
	}
	return nil
}

func (ec *EVENTSCreate) sqlSave(ctx context.Context) (*EVENTS, error) {
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

func (ec *EVENTSCreate) createSpec() (*EVENTS, *sqlgraph.CreateSpec) {
	var (
		_node = &EVENTS{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(events.Table, sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.EventStart(); ok {
		_spec.SetField(events.FieldEventStart, field.TypeTime, value)
		_node.EventStart = value
	}
	if value, ok := ec.mutation.EventEnd(); ok {
		_spec.SetField(events.FieldEventEnd, field.TypeTime, value)
		_node.EventEnd = value
	}
	if value, ok := ec.mutation.ThemeID(); ok {
		_spec.SetField(events.FieldThemeID, field.TypeInt, value)
		_node.ThemeID = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(events.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ec.mutation.ParticipatedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   events.ParticipatedTable,
			Columns: []string{events.ParticipatedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.UsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   events.UsesTable,
			Columns: []string{events.UsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ai_themes.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.events_uses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EVENTSCreateBulk is the builder for creating many EVENTS entities in bulk.
type EVENTSCreateBulk struct {
	config
	err      error
	builders []*EVENTSCreate
}

// Save creates the EVENTS entities in the database.
func (ecb *EVENTSCreateBulk) Save(ctx context.Context) ([]*EVENTS, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*EVENTS, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EVENTSMutation)
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
func (ecb *EVENTSCreateBulk) SaveX(ctx context.Context) []*EVENTS {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EVENTSCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EVENTSCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
