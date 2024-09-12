// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/aithemes"
	"github.com/Hosi121/SpeakUp/ent/matchings"
	"github.com/Hosi121/SpeakUp/ent/sessions"
)

// SESSIONSCreate is the builder for creating a SESSIONS entity.
type SESSIONSCreate struct {
	config
	mutation *SESSIONSMutation
	hooks    []Hook
}

// SetSessionStart sets the "session_start" field.
func (sc *SESSIONSCreate) SetSessionStart(t time.Time) *SESSIONSCreate {
	sc.mutation.SetSessionStart(t)
	return sc
}

// SetSessionEnd sets the "session_end" field.
func (sc *SESSIONSCreate) SetSessionEnd(t time.Time) *SESSIONSCreate {
	sc.mutation.SetSessionEnd(t)
	return sc
}

// SetThemeID sets the "theme_id" field.
func (sc *SESSIONSCreate) SetThemeID(i int) *SESSIONSCreate {
	sc.mutation.SetThemeID(i)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SESSIONSCreate) SetCreatedAt(t time.Time) *SESSIONSCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SESSIONSCreate) SetNillableCreatedAt(t *time.Time) *SESSIONSCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// AddHaIDs adds the "has" edge to the MATCHINGS entity by IDs.
func (sc *SESSIONSCreate) AddHaIDs(ids ...int) *SESSIONSCreate {
	sc.mutation.AddHaIDs(ids...)
	return sc
}

// AddHas adds the "has" edges to the MATCHINGS entity.
func (sc *SESSIONSCreate) AddHas(m ...*MATCHINGS) *SESSIONSCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return sc.AddHaIDs(ids...)
}

// SetUsesID sets the "uses" edge to the AITHEMES entity by ID.
func (sc *SESSIONSCreate) SetUsesID(id int) *SESSIONSCreate {
	sc.mutation.SetUsesID(id)
	return sc
}

// SetNillableUsesID sets the "uses" edge to the AITHEMES entity by ID if the given value is not nil.
func (sc *SESSIONSCreate) SetNillableUsesID(id *int) *SESSIONSCreate {
	if id != nil {
		sc = sc.SetUsesID(*id)
	}
	return sc
}

// SetUses sets the "uses" edge to the AITHEMES entity.
func (sc *SESSIONSCreate) SetUses(a *AITHEMES) *SESSIONSCreate {
	return sc.SetUsesID(a.ID)
}

// Mutation returns the SESSIONSMutation object of the builder.
func (sc *SESSIONSCreate) Mutation() *SESSIONSMutation {
	return sc.mutation
}

// Save creates the SESSIONS in the database.
func (sc *SESSIONSCreate) Save(ctx context.Context) (*SESSIONS, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SESSIONSCreate) SaveX(ctx context.Context) *SESSIONS {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SESSIONSCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SESSIONSCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SESSIONSCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := sessions.DefaultCreatedAt
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SESSIONSCreate) check() error {
	if _, ok := sc.mutation.SessionStart(); !ok {
		return &ValidationError{Name: "session_start", err: errors.New(`ent: missing required field "SESSIONS.session_start"`)}
	}
	if _, ok := sc.mutation.SessionEnd(); !ok {
		return &ValidationError{Name: "session_end", err: errors.New(`ent: missing required field "SESSIONS.session_end"`)}
	}
	if _, ok := sc.mutation.ThemeID(); !ok {
		return &ValidationError{Name: "theme_id", err: errors.New(`ent: missing required field "SESSIONS.theme_id"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SESSIONS.created_at"`)}
	}
	return nil
}

func (sc *SESSIONSCreate) sqlSave(ctx context.Context) (*SESSIONS, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SESSIONSCreate) createSpec() (*SESSIONS, *sqlgraph.CreateSpec) {
	var (
		_node = &SESSIONS{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(sessions.Table, sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.SessionStart(); ok {
		_spec.SetField(sessions.FieldSessionStart, field.TypeTime, value)
		_node.SessionStart = value
	}
	if value, ok := sc.mutation.SessionEnd(); ok {
		_spec.SetField(sessions.FieldSessionEnd, field.TypeTime, value)
		_node.SessionEnd = value
	}
	if value, ok := sc.mutation.ThemeID(); ok {
		_spec.SetField(sessions.FieldThemeID, field.TypeInt, value)
		_node.ThemeID = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(sessions.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := sc.mutation.HasIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sessions.HasTable,
			Columns: []string{sessions.HasColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(matchings.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.UsesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sessions.UsesTable,
			Columns: []string{sessions.UsesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(aithemes.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sessions_uses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SESSIONSCreateBulk is the builder for creating many SESSIONS entities in bulk.
type SESSIONSCreateBulk struct {
	config
	err      error
	builders []*SESSIONSCreate
}

// Save creates the SESSIONS entities in the database.
func (scb *SESSIONSCreateBulk) Save(ctx context.Context) ([]*SESSIONS, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*SESSIONS, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SESSIONSMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SESSIONSCreateBulk) SaveX(ctx context.Context) []*SESSIONS {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SESSIONSCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SESSIONSCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
