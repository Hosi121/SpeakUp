// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/calls"
	"github.com/Hosi121/SpeakUp/ent/matchings"
	"github.com/Hosi121/SpeakUp/ent/sessions"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// MATCHINGSCreate is the builder for creating a MATCHINGS entity.
type MATCHINGSCreate struct {
	config
	mutation *MATCHINGSMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (mc *MATCHINGSCreate) SetUserID(i int) *MATCHINGSCreate {
	mc.mutation.SetUserID(i)
	return mc
}

// SetMatchedUserID sets the "matched_user_id" field.
func (mc *MATCHINGSCreate) SetMatchedUserID(i int) *MATCHINGSCreate {
	mc.mutation.SetMatchedUserID(i)
	return mc
}

// SetSessionID sets the "session_id" field.
func (mc *MATCHINGSCreate) SetSessionID(i int) *MATCHINGSCreate {
	mc.mutation.SetSessionID(i)
	return mc
}

// SetMatchedAt sets the "matched_at" field.
func (mc *MATCHINGSCreate) SetMatchedAt(t time.Time) *MATCHINGSCreate {
	mc.mutation.SetMatchedAt(t)
	return mc
}

// SetNillableMatchedAt sets the "matched_at" field if the given value is not nil.
func (mc *MATCHINGSCreate) SetNillableMatchedAt(t *time.Time) *MATCHINGSCreate {
	if t != nil {
		mc.SetMatchedAt(*t)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MATCHINGSCreate) SetStatus(m matchings.Status) *MATCHINGSCreate {
	mc.mutation.SetStatus(m)
	return mc
}

// AddMemberIDs adds the "member" edge to the USERS entity by IDs.
func (mc *MATCHINGSCreate) AddMemberIDs(ids ...int) *MATCHINGSCreate {
	mc.mutation.AddMemberIDs(ids...)
	return mc
}

// AddMember adds the "member" edges to the USERS entity.
func (mc *MATCHINGSCreate) AddMember(u ...*USERS) *MATCHINGSCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return mc.AddMemberIDs(ids...)
}

// SetHadID sets the "had" edge to the SESSIONS entity by ID.
func (mc *MATCHINGSCreate) SetHadID(id int) *MATCHINGSCreate {
	mc.mutation.SetHadID(id)
	return mc
}

// SetNillableHadID sets the "had" edge to the SESSIONS entity by ID if the given value is not nil.
func (mc *MATCHINGSCreate) SetNillableHadID(id *int) *MATCHINGSCreate {
	if id != nil {
		mc = mc.SetHadID(*id)
	}
	return mc
}

// SetHad sets the "had" edge to the SESSIONS entity.
func (mc *MATCHINGSCreate) SetHad(s *SESSIONS) *MATCHINGSCreate {
	return mc.SetHadID(s.ID)
}

// SetMakesID sets the "makes" edge to the CALLS entity by ID.
func (mc *MATCHINGSCreate) SetMakesID(id int) *MATCHINGSCreate {
	mc.mutation.SetMakesID(id)
	return mc
}

// SetNillableMakesID sets the "makes" edge to the CALLS entity by ID if the given value is not nil.
func (mc *MATCHINGSCreate) SetNillableMakesID(id *int) *MATCHINGSCreate {
	if id != nil {
		mc = mc.SetMakesID(*id)
	}
	return mc
}

// SetMakes sets the "makes" edge to the CALLS entity.
func (mc *MATCHINGSCreate) SetMakes(c *CALLS) *MATCHINGSCreate {
	return mc.SetMakesID(c.ID)
}

// Mutation returns the MATCHINGSMutation object of the builder.
func (mc *MATCHINGSCreate) Mutation() *MATCHINGSMutation {
	return mc.mutation
}

// Save creates the MATCHINGS in the database.
func (mc *MATCHINGSCreate) Save(ctx context.Context) (*MATCHINGS, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MATCHINGSCreate) SaveX(ctx context.Context) *MATCHINGS {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MATCHINGSCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MATCHINGSCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MATCHINGSCreate) defaults() {
	if _, ok := mc.mutation.MatchedAt(); !ok {
		v := matchings.DefaultMatchedAt
		mc.mutation.SetMatchedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MATCHINGSCreate) check() error {
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "MATCHINGS.user_id"`)}
	}
	if _, ok := mc.mutation.MatchedUserID(); !ok {
		return &ValidationError{Name: "matched_user_id", err: errors.New(`ent: missing required field "MATCHINGS.matched_user_id"`)}
	}
	if _, ok := mc.mutation.SessionID(); !ok {
		return &ValidationError{Name: "session_id", err: errors.New(`ent: missing required field "MATCHINGS.session_id"`)}
	}
	if _, ok := mc.mutation.MatchedAt(); !ok {
		return &ValidationError{Name: "matched_at", err: errors.New(`ent: missing required field "MATCHINGS.matched_at"`)}
	}
	if _, ok := mc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "MATCHINGS.status"`)}
	}
	if v, ok := mc.mutation.Status(); ok {
		if err := matchings.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "MATCHINGS.status": %w`, err)}
		}
	}
	return nil
}

func (mc *MATCHINGSCreate) sqlSave(ctx context.Context) (*MATCHINGS, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MATCHINGSCreate) createSpec() (*MATCHINGS, *sqlgraph.CreateSpec) {
	var (
		_node = &MATCHINGS{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(matchings.Table, sqlgraph.NewFieldSpec(matchings.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.UserID(); ok {
		_spec.SetField(matchings.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := mc.mutation.MatchedUserID(); ok {
		_spec.SetField(matchings.FieldMatchedUserID, field.TypeInt, value)
		_node.MatchedUserID = value
	}
	if value, ok := mc.mutation.SessionID(); ok {
		_spec.SetField(matchings.FieldSessionID, field.TypeInt, value)
		_node.SessionID = value
	}
	if value, ok := mc.mutation.MatchedAt(); ok {
		_spec.SetField(matchings.FieldMatchedAt, field.TypeTime, value)
		_node.MatchedAt = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(matchings.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := mc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   matchings.MemberTable,
			Columns: matchings.MemberPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.HadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   matchings.HadTable,
			Columns: []string{matchings.HadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessions.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sessions_has = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   matchings.MakesTable,
			Columns: []string{matchings.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(calls.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MATCHINGSCreateBulk is the builder for creating many MATCHINGS entities in bulk.
type MATCHINGSCreateBulk struct {
	config
	err      error
	builders []*MATCHINGSCreate
}

// Save creates the MATCHINGS entities in the database.
func (mcb *MATCHINGSCreateBulk) Save(ctx context.Context) ([]*MATCHINGS, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*MATCHINGS, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MATCHINGSMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MATCHINGSCreateBulk) SaveX(ctx context.Context) []*MATCHINGS {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MATCHINGSCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MATCHINGSCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}