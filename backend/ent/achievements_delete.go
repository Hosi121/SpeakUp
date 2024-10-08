// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Hosi121/SpeakUp/ent/achievements"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// ACHIEVEMENTSDelete is the builder for deleting a ACHIEVEMENTS entity.
type ACHIEVEMENTSDelete struct {
	config
	hooks    []Hook
	mutation *ACHIEVEMENTSMutation
}

// Where appends a list predicates to the ACHIEVEMENTSDelete builder.
func (ad *ACHIEVEMENTSDelete) Where(ps ...predicate.ACHIEVEMENTS) *ACHIEVEMENTSDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *ACHIEVEMENTSDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *ACHIEVEMENTSDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *ACHIEVEMENTSDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(achievements.Table, sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// ACHIEVEMENTSDeleteOne is the builder for deleting a single ACHIEVEMENTS entity.
type ACHIEVEMENTSDeleteOne struct {
	ad *ACHIEVEMENTSDelete
}

// Where appends a list predicates to the ACHIEVEMENTSDelete builder.
func (ado *ACHIEVEMENTSDeleteOne) Where(ps ...predicate.ACHIEVEMENTS) *ACHIEVEMENTSDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *ACHIEVEMENTSDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{achievements.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *ACHIEVEMENTSDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}
