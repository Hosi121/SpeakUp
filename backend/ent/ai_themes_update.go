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
	"github.com/Hosi121/SpeakUp/ent/ai_themes"
	"github.com/Hosi121/SpeakUp/ent/events"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// AITHEMESUpdate is the builder for updating AI_THEMES entities.
type AITHEMESUpdate struct {
	config
	hooks    []Hook
	mutation *AITHEMESMutation
}

// Where appends a list predicates to the AITHEMESUpdate builder.
func (au *AITHEMESUpdate) Where(ps ...predicate.AI_THEMES) *AITHEMESUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetThemeText sets the "theme_text" field.
func (au *AITHEMESUpdate) SetThemeText(s string) *AITHEMESUpdate {
	au.mutation.SetThemeText(s)
	return au
}

// SetNillableThemeText sets the "theme_text" field if the given value is not nil.
func (au *AITHEMESUpdate) SetNillableThemeText(s *string) *AITHEMESUpdate {
	if s != nil {
		au.SetThemeText(*s)
	}
	return au
}

// SetTopic1 sets the "topic1" field.
func (au *AITHEMESUpdate) SetTopic1(s string) *AITHEMESUpdate {
	au.mutation.SetTopic1(s)
	return au
}

// SetNillableTopic1 sets the "topic1" field if the given value is not nil.
func (au *AITHEMESUpdate) SetNillableTopic1(s *string) *AITHEMESUpdate {
	if s != nil {
		au.SetTopic1(*s)
	}
	return au
}

// SetTopic2 sets the "topic2" field.
func (au *AITHEMESUpdate) SetTopic2(s string) *AITHEMESUpdate {
	au.mutation.SetTopic2(s)
	return au
}

// SetNillableTopic2 sets the "topic2" field if the given value is not nil.
func (au *AITHEMESUpdate) SetNillableTopic2(s *string) *AITHEMESUpdate {
	if s != nil {
		au.SetTopic2(*s)
	}
	return au
}

// SetTopic3 sets the "topic3" field.
func (au *AITHEMESUpdate) SetTopic3(s string) *AITHEMESUpdate {
	au.mutation.SetTopic3(s)
	return au
}

// SetNillableTopic3 sets the "topic3" field if the given value is not nil.
func (au *AITHEMESUpdate) SetNillableTopic3(s *string) *AITHEMESUpdate {
	if s != nil {
		au.SetTopic3(*s)
	}
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AITHEMESUpdate) SetCreatedAt(t time.Time) *AITHEMESUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AITHEMESUpdate) SetNillableCreatedAt(t *time.Time) *AITHEMESUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// AddUsedIDs adds the "used" edge to the EVENTS entity by IDs.
func (au *AITHEMESUpdate) AddUsedIDs(ids ...int) *AITHEMESUpdate {
	au.mutation.AddUsedIDs(ids...)
	return au
}

// AddUsed adds the "used" edges to the EVENTS entity.
func (au *AITHEMESUpdate) AddUsed(e ...*EVENTS) *AITHEMESUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.AddUsedIDs(ids...)
}

// Mutation returns the AITHEMESMutation object of the builder.
func (au *AITHEMESUpdate) Mutation() *AITHEMESMutation {
	return au.mutation
}

// ClearUsed clears all "used" edges to the EVENTS entity.
func (au *AITHEMESUpdate) ClearUsed() *AITHEMESUpdate {
	au.mutation.ClearUsed()
	return au
}

// RemoveUsedIDs removes the "used" edge to EVENTS entities by IDs.
func (au *AITHEMESUpdate) RemoveUsedIDs(ids ...int) *AITHEMESUpdate {
	au.mutation.RemoveUsedIDs(ids...)
	return au
}

// RemoveUsed removes "used" edges to EVENTS entities.
func (au *AITHEMESUpdate) RemoveUsed(e ...*EVENTS) *AITHEMESUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return au.RemoveUsedIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AITHEMESUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AITHEMESUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AITHEMESUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AITHEMESUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AITHEMESUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ai_themes.Table, ai_themes.Columns, sqlgraph.NewFieldSpec(ai_themes.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.ThemeText(); ok {
		_spec.SetField(ai_themes.FieldThemeText, field.TypeString, value)
	}
	if value, ok := au.mutation.Topic1(); ok {
		_spec.SetField(ai_themes.FieldTopic1, field.TypeString, value)
	}
	if value, ok := au.mutation.Topic2(); ok {
		_spec.SetField(ai_themes.FieldTopic2, field.TypeString, value)
	}
	if value, ok := au.mutation.Topic3(); ok {
		_spec.SetField(ai_themes.FieldTopic3, field.TypeString, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(ai_themes.FieldCreatedAt, field.TypeTime, value)
	}
	if au.mutation.UsedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ai_themes.UsedTable,
			Columns: []string{ai_themes.UsedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedUsedIDs(); len(nodes) > 0 && !au.mutation.UsedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ai_themes.UsedTable,
			Columns: []string{ai_themes.UsedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UsedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ai_themes.UsedTable,
			Columns: []string{ai_themes.UsedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ai_themes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AITHEMESUpdateOne is the builder for updating a single AI_THEMES entity.
type AITHEMESUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AITHEMESMutation
}

// SetThemeText sets the "theme_text" field.
func (auo *AITHEMESUpdateOne) SetThemeText(s string) *AITHEMESUpdateOne {
	auo.mutation.SetThemeText(s)
	return auo
}

// SetNillableThemeText sets the "theme_text" field if the given value is not nil.
func (auo *AITHEMESUpdateOne) SetNillableThemeText(s *string) *AITHEMESUpdateOne {
	if s != nil {
		auo.SetThemeText(*s)
	}
	return auo
}

// SetTopic1 sets the "topic1" field.
func (auo *AITHEMESUpdateOne) SetTopic1(s string) *AITHEMESUpdateOne {
	auo.mutation.SetTopic1(s)
	return auo
}

// SetNillableTopic1 sets the "topic1" field if the given value is not nil.
func (auo *AITHEMESUpdateOne) SetNillableTopic1(s *string) *AITHEMESUpdateOne {
	if s != nil {
		auo.SetTopic1(*s)
	}
	return auo
}

// SetTopic2 sets the "topic2" field.
func (auo *AITHEMESUpdateOne) SetTopic2(s string) *AITHEMESUpdateOne {
	auo.mutation.SetTopic2(s)
	return auo
}

// SetNillableTopic2 sets the "topic2" field if the given value is not nil.
func (auo *AITHEMESUpdateOne) SetNillableTopic2(s *string) *AITHEMESUpdateOne {
	if s != nil {
		auo.SetTopic2(*s)
	}
	return auo
}

// SetTopic3 sets the "topic3" field.
func (auo *AITHEMESUpdateOne) SetTopic3(s string) *AITHEMESUpdateOne {
	auo.mutation.SetTopic3(s)
	return auo
}

// SetNillableTopic3 sets the "topic3" field if the given value is not nil.
func (auo *AITHEMESUpdateOne) SetNillableTopic3(s *string) *AITHEMESUpdateOne {
	if s != nil {
		auo.SetTopic3(*s)
	}
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AITHEMESUpdateOne) SetCreatedAt(t time.Time) *AITHEMESUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AITHEMESUpdateOne) SetNillableCreatedAt(t *time.Time) *AITHEMESUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// AddUsedIDs adds the "used" edge to the EVENTS entity by IDs.
func (auo *AITHEMESUpdateOne) AddUsedIDs(ids ...int) *AITHEMESUpdateOne {
	auo.mutation.AddUsedIDs(ids...)
	return auo
}

// AddUsed adds the "used" edges to the EVENTS entity.
func (auo *AITHEMESUpdateOne) AddUsed(e ...*EVENTS) *AITHEMESUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.AddUsedIDs(ids...)
}

// Mutation returns the AITHEMESMutation object of the builder.
func (auo *AITHEMESUpdateOne) Mutation() *AITHEMESMutation {
	return auo.mutation
}

// ClearUsed clears all "used" edges to the EVENTS entity.
func (auo *AITHEMESUpdateOne) ClearUsed() *AITHEMESUpdateOne {
	auo.mutation.ClearUsed()
	return auo
}

// RemoveUsedIDs removes the "used" edge to EVENTS entities by IDs.
func (auo *AITHEMESUpdateOne) RemoveUsedIDs(ids ...int) *AITHEMESUpdateOne {
	auo.mutation.RemoveUsedIDs(ids...)
	return auo
}

// RemoveUsed removes "used" edges to EVENTS entities.
func (auo *AITHEMESUpdateOne) RemoveUsed(e ...*EVENTS) *AITHEMESUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return auo.RemoveUsedIDs(ids...)
}

// Where appends a list predicates to the AITHEMESUpdate builder.
func (auo *AITHEMESUpdateOne) Where(ps ...predicate.AI_THEMES) *AITHEMESUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AITHEMESUpdateOne) Select(field string, fields ...string) *AITHEMESUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated AI_THEMES entity.
func (auo *AITHEMESUpdateOne) Save(ctx context.Context) (*AI_THEMES, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AITHEMESUpdateOne) SaveX(ctx context.Context) *AI_THEMES {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AITHEMESUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AITHEMESUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AITHEMESUpdateOne) sqlSave(ctx context.Context) (_node *AI_THEMES, err error) {
	_spec := sqlgraph.NewUpdateSpec(ai_themes.Table, ai_themes.Columns, sqlgraph.NewFieldSpec(ai_themes.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AI_THEMES.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ai_themes.FieldID)
		for _, f := range fields {
			if !ai_themes.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ai_themes.FieldID {
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
	if value, ok := auo.mutation.ThemeText(); ok {
		_spec.SetField(ai_themes.FieldThemeText, field.TypeString, value)
	}
	if value, ok := auo.mutation.Topic1(); ok {
		_spec.SetField(ai_themes.FieldTopic1, field.TypeString, value)
	}
	if value, ok := auo.mutation.Topic2(); ok {
		_spec.SetField(ai_themes.FieldTopic2, field.TypeString, value)
	}
	if value, ok := auo.mutation.Topic3(); ok {
		_spec.SetField(ai_themes.FieldTopic3, field.TypeString, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(ai_themes.FieldCreatedAt, field.TypeTime, value)
	}
	if auo.mutation.UsedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ai_themes.UsedTable,
			Columns: []string{ai_themes.UsedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedUsedIDs(); len(nodes) > 0 && !auo.mutation.UsedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ai_themes.UsedTable,
			Columns: []string{ai_themes.UsedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UsedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ai_themes.UsedTable,
			Columns: []string{ai_themes.UsedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(events.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AI_THEMES{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ai_themes.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
