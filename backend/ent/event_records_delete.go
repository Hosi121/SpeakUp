// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// EVENTRECORDSDelete is the builder for deleting a EVENT_RECORDS entity.
type EVENTRECORDSDelete struct {
	config
	hooks    []Hook
	mutation *EVENTRECORDSMutation
}

// Where appends a list predicates to the EVENTRECORDSDelete builder.
func (ed *EVENTRECORDSDelete) Where(ps ...predicate.EVENT_RECORDS) *EVENTRECORDSDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EVENTRECORDSDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ed.sqlExec, ed.mutation, ed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EVENTRECORDSDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EVENTRECORDSDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(event_records.Table, sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt))
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ed.mutation.done = true
	return affected, err
}

// EVENTRECORDSDeleteOne is the builder for deleting a single EVENT_RECORDS entity.
type EVENTRECORDSDeleteOne struct {
	ed *EVENTRECORDSDelete
}

// Where appends a list predicates to the EVENTRECORDSDelete builder.
func (edo *EVENTRECORDSDeleteOne) Where(ps ...predicate.EVENT_RECORDS) *EVENTRECORDSDeleteOne {
	edo.ed.mutation.Where(ps...)
	return edo
}

// Exec executes the deletion query.
func (edo *EVENTRECORDSDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{event_records.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EVENTRECORDSDeleteOne) ExecX(ctx context.Context) {
	if err := edo.Exec(ctx); err != nil {
		panic(err)
	}
}
