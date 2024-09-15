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
	"github.com/Hosi121/SpeakUp/ent/sessions"
)

// AITHEMESCreate is the builder for creating a AITHEMES entity.
type AITHEMESCreate struct {
	config
	mutation *AITHEMESMutation
	hooks    []Hook
}

// SetThemeText sets the "theme_text" field.
func (ac *AITHEMESCreate) SetThemeText(s string) *AITHEMESCreate {
	ac.mutation.SetThemeText(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AITHEMESCreate) SetCreatedAt(t time.Time) *AITHEMESCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AITHEMESCreate) SetNillableCreatedAt(t *time.Time) *AITHEMESCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// AddUsedIDs adds the "used" edge to the SESSIONS entity by IDs.
func (ac *AITHEMESCreate) AddUsedIDs(ids ...int) *AITHEMESCreate {
	ac.mutation.AddUsedIDs(ids...)
	return ac
}

// AddUsed adds the "used" edges to the SESSIONS entity.
func (ac *AITHEMESCreate) AddUsed(s ...*SESSIONS) *AITHEMESCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ac.AddUsedIDs(ids...)
}

// Mutation returns the AITHEMESMutation object of the builder.
func (ac *AITHEMESCreate) Mutation() *AITHEMESMutation {
	return ac.mutation
}

// Save creates the AITHEMES in the database.
func (ac *AITHEMESCreate) Save(ctx context.Context) (*AITHEMES, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AITHEMESCreate) SaveX(ctx context.Context) *AITHEMES {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AITHEMESCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AITHEMESCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AITHEMESCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := aithemes.DefaultCreatedAt
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AITHEMESCreate) check() error {
	if _, ok := ac.mutation.ThemeText(); !ok {
		return &ValidationError{Name: "theme_text", err: errors.New(`ent: missing required field "AITHEMES.theme_text"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AITHEMES.created_at"`)}
	}
	return nil
}

func (ac *AITHEMESCreate) sqlSave(ctx context.Context) (*AITHEMES, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AITHEMESCreate) createSpec() (*AITHEMES, *sqlgraph.CreateSpec) {
	var (
		_node = &AITHEMES{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(aithemes.Table, sqlgraph.NewFieldSpec(aithemes.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.ThemeText(); ok {
		_spec.SetField(aithemes.FieldThemeText, field.TypeString, value)
		_node.ThemeText = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(aithemes.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.UsedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   aithemes.UsedTable,
			Columns: []string{aithemes.UsedColumn},
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

// AITHEMESCreateBulk is the builder for creating many AITHEMES entities in bulk.
type AITHEMESCreateBulk struct {
	config
	err      error
	builders []*AITHEMESCreate
}

// Save creates the AITHEMES entities in the database.
func (acb *AITHEMESCreateBulk) Save(ctx context.Context) ([]*AITHEMES, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*AITHEMES, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AITHEMESMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AITHEMESCreateBulk) SaveX(ctx context.Context) []*AITHEMES {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AITHEMESCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AITHEMESCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
