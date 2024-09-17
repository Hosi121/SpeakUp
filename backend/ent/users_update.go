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
	"github.com/Hosi121/SpeakUp/ent/achievements"
	"github.com/Hosi121/SpeakUp/ent/event_records"
	"github.com/Hosi121/SpeakUp/ent/friends"
	"github.com/Hosi121/SpeakUp/ent/memos"
	"github.com/Hosi121/SpeakUp/ent/predicate"
	"github.com/Hosi121/SpeakUp/ent/progress"
	"github.com/Hosi121/SpeakUp/ent/users"
)

// USERSUpdate is the builder for updating USERS entities.
type USERSUpdate struct {
	config
	hooks    []Hook
	mutation *USERSMutation
}

// Where appends a list predicates to the USERSUpdate builder.
func (uu *USERSUpdate) Where(ps ...predicate.USERS) *USERSUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *USERSUpdate) SetUsername(s string) *USERSUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *USERSUpdate) SetNillableUsername(s *string) *USERSUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// SetAvatarURL sets the "avatar_url" field.
func (uu *USERSUpdate) SetAvatarURL(s string) *USERSUpdate {
	uu.mutation.SetAvatarURL(s)
	return uu
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uu *USERSUpdate) SetNillableAvatarURL(s *string) *USERSUpdate {
	if s != nil {
		uu.SetAvatarURL(*s)
	}
	return uu
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (uu *USERSUpdate) ClearAvatarURL() *USERSUpdate {
	uu.mutation.ClearAvatarURL()
	return uu
}

// SetRole sets the "role" field.
func (uu *USERSUpdate) SetRole(u users.Role) *USERSUpdate {
	uu.mutation.SetRole(u)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *USERSUpdate) SetNillableRole(u *users.Role) *USERSUpdate {
	if u != nil {
		uu.SetRole(*u)
	}
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *USERSUpdate) SetCreatedAt(t time.Time) *USERSUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *USERSUpdate) SetNillableCreatedAt(t *time.Time) *USERSUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetIsDeleted sets the "is_deleted" field.
func (uu *USERSUpdate) SetIsDeleted(b bool) *USERSUpdate {
	uu.mutation.SetIsDeleted(b)
	return uu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (uu *USERSUpdate) SetNillableIsDeleted(b *bool) *USERSUpdate {
	if b != nil {
		uu.SetIsDeleted(*b)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *USERSUpdate) SetUpdatedAt(t time.Time) *USERSUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *USERSUpdate) SetNillableUpdatedAt(t *time.Time) *USERSUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// AddConnectIDs adds the "connects" edge to the FRIENDS entity by IDs.
func (uu *USERSUpdate) AddConnectIDs(ids ...int) *USERSUpdate {
	uu.mutation.AddConnectIDs(ids...)
	return uu
}

// AddConnects adds the "connects" edges to the FRIENDS entity.
func (uu *USERSUpdate) AddConnects(f ...*FRIENDS) *USERSUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.AddConnectIDs(ids...)
}

// AddMakeIDs adds the "makes" edge to the EVENT_RECORDS entity by IDs.
func (uu *USERSUpdate) AddMakeIDs(ids ...int) *USERSUpdate {
	uu.mutation.AddMakeIDs(ids...)
	return uu
}

// AddMakes adds the "makes" edges to the EVENT_RECORDS entity.
func (uu *USERSUpdate) AddMakes(e ...*EVENT_RECORDS) *USERSUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.AddMakeIDs(ids...)
}

// SetPreparesID sets the "prepares" edge to the MEMOS entity by ID.
func (uu *USERSUpdate) SetPreparesID(id int) *USERSUpdate {
	uu.mutation.SetPreparesID(id)
	return uu
}

// SetNillablePreparesID sets the "prepares" edge to the MEMOS entity by ID if the given value is not nil.
func (uu *USERSUpdate) SetNillablePreparesID(id *int) *USERSUpdate {
	if id != nil {
		uu = uu.SetPreparesID(*id)
	}
	return uu
}

// SetPrepares sets the "prepares" edge to the MEMOS entity.
func (uu *USERSUpdate) SetPrepares(m *MEMOS) *USERSUpdate {
	return uu.SetPreparesID(m.ID)
}

// AddAcquireIDs adds the "acquires" edge to the ACHIEVEMENTS entity by IDs.
func (uu *USERSUpdate) AddAcquireIDs(ids ...int) *USERSUpdate {
	uu.mutation.AddAcquireIDs(ids...)
	return uu
}

// AddAcquires adds the "acquires" edges to the ACHIEVEMENTS entity.
func (uu *USERSUpdate) AddAcquires(a ...*ACHIEVEMENTS) *USERSUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAcquireIDs(ids...)
}

// SetRecordsID sets the "records" edge to the PROGRESS entity by ID.
func (uu *USERSUpdate) SetRecordsID(id int) *USERSUpdate {
	uu.mutation.SetRecordsID(id)
	return uu
}

// SetNillableRecordsID sets the "records" edge to the PROGRESS entity by ID if the given value is not nil.
func (uu *USERSUpdate) SetNillableRecordsID(id *int) *USERSUpdate {
	if id != nil {
		uu = uu.SetRecordsID(*id)
	}
	return uu
}

// SetRecords sets the "records" edge to the PROGRESS entity.
func (uu *USERSUpdate) SetRecords(p *PROGRESS) *USERSUpdate {
	return uu.SetRecordsID(p.ID)
}

// Mutation returns the USERSMutation object of the builder.
func (uu *USERSUpdate) Mutation() *USERSMutation {
	return uu.mutation
}

// ClearConnects clears all "connects" edges to the FRIENDS entity.
func (uu *USERSUpdate) ClearConnects() *USERSUpdate {
	uu.mutation.ClearConnects()
	return uu
}

// RemoveConnectIDs removes the "connects" edge to FRIENDS entities by IDs.
func (uu *USERSUpdate) RemoveConnectIDs(ids ...int) *USERSUpdate {
	uu.mutation.RemoveConnectIDs(ids...)
	return uu
}

// RemoveConnects removes "connects" edges to FRIENDS entities.
func (uu *USERSUpdate) RemoveConnects(f ...*FRIENDS) *USERSUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.RemoveConnectIDs(ids...)
}

// ClearMakes clears all "makes" edges to the EVENT_RECORDS entity.
func (uu *USERSUpdate) ClearMakes() *USERSUpdate {
	uu.mutation.ClearMakes()
	return uu
}

// RemoveMakeIDs removes the "makes" edge to EVENT_RECORDS entities by IDs.
func (uu *USERSUpdate) RemoveMakeIDs(ids ...int) *USERSUpdate {
	uu.mutation.RemoveMakeIDs(ids...)
	return uu
}

// RemoveMakes removes "makes" edges to EVENT_RECORDS entities.
func (uu *USERSUpdate) RemoveMakes(e ...*EVENT_RECORDS) *USERSUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uu.RemoveMakeIDs(ids...)
}

// ClearPrepares clears the "prepares" edge to the MEMOS entity.
func (uu *USERSUpdate) ClearPrepares() *USERSUpdate {
	uu.mutation.ClearPrepares()
	return uu
}

// ClearAcquires clears all "acquires" edges to the ACHIEVEMENTS entity.
func (uu *USERSUpdate) ClearAcquires() *USERSUpdate {
	uu.mutation.ClearAcquires()
	return uu
}

// RemoveAcquireIDs removes the "acquires" edge to ACHIEVEMENTS entities by IDs.
func (uu *USERSUpdate) RemoveAcquireIDs(ids ...int) *USERSUpdate {
	uu.mutation.RemoveAcquireIDs(ids...)
	return uu
}

// RemoveAcquires removes "acquires" edges to ACHIEVEMENTS entities.
func (uu *USERSUpdate) RemoveAcquires(a ...*ACHIEVEMENTS) *USERSUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAcquireIDs(ids...)
}

// ClearRecords clears the "records" edge to the PROGRESS entity.
func (uu *USERSUpdate) ClearRecords() *USERSUpdate {
	uu.mutation.ClearRecords()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *USERSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *USERSUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *USERSUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *USERSUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *USERSUpdate) check() error {
	if v, ok := uu.mutation.Username(); ok {
		if err := users.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "USERS.username": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Role(); ok {
		if err := users.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "USERS.role": %w`, err)}
		}
	}
	return nil
}

func (uu *USERSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.AvatarURL(); ok {
		_spec.SetField(users.FieldAvatarURL, field.TypeString, value)
	}
	if uu.mutation.AvatarURLCleared() {
		_spec.ClearField(users.FieldAvatarURL, field.TypeString)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(users.FieldRole, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(users.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.IsDeleted(); ok {
		_spec.SetField(users.FieldIsDeleted, field.TypeBool, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(users.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.ConnectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   users.ConnectsTable,
			Columns: users.ConnectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friends.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedConnectsIDs(); len(nodes) > 0 && !uu.mutation.ConnectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   users.ConnectsTable,
			Columns: users.ConnectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friends.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ConnectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   users.ConnectsTable,
			Columns: users.ConnectsPrimaryKey,
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
	if uu.mutation.MakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.MakesTable,
			Columns: []string{users.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedMakesIDs(); len(nodes) > 0 && !uu.mutation.MakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.MakesTable,
			Columns: []string{users.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.MakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.MakesTable,
			Columns: []string{users.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PreparesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   users.PreparesTable,
			Columns: []string{users.PreparesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memos.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PreparesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   users.PreparesTable,
			Columns: []string{users.PreparesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memos.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AcquiresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AcquiresTable,
			Columns: []string{users.AcquiresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAcquiresIDs(); len(nodes) > 0 && !uu.mutation.AcquiresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AcquiresTable,
			Columns: []string{users.AcquiresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AcquiresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AcquiresTable,
			Columns: []string{users.AcquiresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   users.RecordsTable,
			Columns: []string{users.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   users.RecordsTable,
			Columns: []string{users.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// USERSUpdateOne is the builder for updating a single USERS entity.
type USERSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *USERSMutation
}

// SetUsername sets the "username" field.
func (uuo *USERSUpdateOne) SetUsername(s string) *USERSUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *USERSUpdateOne) SetNillableUsername(s *string) *USERSUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// SetAvatarURL sets the "avatar_url" field.
func (uuo *USERSUpdateOne) SetAvatarURL(s string) *USERSUpdateOne {
	uuo.mutation.SetAvatarURL(s)
	return uuo
}

// SetNillableAvatarURL sets the "avatar_url" field if the given value is not nil.
func (uuo *USERSUpdateOne) SetNillableAvatarURL(s *string) *USERSUpdateOne {
	if s != nil {
		uuo.SetAvatarURL(*s)
	}
	return uuo
}

// ClearAvatarURL clears the value of the "avatar_url" field.
func (uuo *USERSUpdateOne) ClearAvatarURL() *USERSUpdateOne {
	uuo.mutation.ClearAvatarURL()
	return uuo
}

// SetRole sets the "role" field.
func (uuo *USERSUpdateOne) SetRole(u users.Role) *USERSUpdateOne {
	uuo.mutation.SetRole(u)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *USERSUpdateOne) SetNillableRole(u *users.Role) *USERSUpdateOne {
	if u != nil {
		uuo.SetRole(*u)
	}
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *USERSUpdateOne) SetCreatedAt(t time.Time) *USERSUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *USERSUpdateOne) SetNillableCreatedAt(t *time.Time) *USERSUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetIsDeleted sets the "is_deleted" field.
func (uuo *USERSUpdateOne) SetIsDeleted(b bool) *USERSUpdateOne {
	uuo.mutation.SetIsDeleted(b)
	return uuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (uuo *USERSUpdateOne) SetNillableIsDeleted(b *bool) *USERSUpdateOne {
	if b != nil {
		uuo.SetIsDeleted(*b)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *USERSUpdateOne) SetUpdatedAt(t time.Time) *USERSUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *USERSUpdateOne) SetNillableUpdatedAt(t *time.Time) *USERSUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// AddConnectIDs adds the "connects" edge to the FRIENDS entity by IDs.
func (uuo *USERSUpdateOne) AddConnectIDs(ids ...int) *USERSUpdateOne {
	uuo.mutation.AddConnectIDs(ids...)
	return uuo
}

// AddConnects adds the "connects" edges to the FRIENDS entity.
func (uuo *USERSUpdateOne) AddConnects(f ...*FRIENDS) *USERSUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.AddConnectIDs(ids...)
}

// AddMakeIDs adds the "makes" edge to the EVENT_RECORDS entity by IDs.
func (uuo *USERSUpdateOne) AddMakeIDs(ids ...int) *USERSUpdateOne {
	uuo.mutation.AddMakeIDs(ids...)
	return uuo
}

// AddMakes adds the "makes" edges to the EVENT_RECORDS entity.
func (uuo *USERSUpdateOne) AddMakes(e ...*EVENT_RECORDS) *USERSUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.AddMakeIDs(ids...)
}

// SetPreparesID sets the "prepares" edge to the MEMOS entity by ID.
func (uuo *USERSUpdateOne) SetPreparesID(id int) *USERSUpdateOne {
	uuo.mutation.SetPreparesID(id)
	return uuo
}

// SetNillablePreparesID sets the "prepares" edge to the MEMOS entity by ID if the given value is not nil.
func (uuo *USERSUpdateOne) SetNillablePreparesID(id *int) *USERSUpdateOne {
	if id != nil {
		uuo = uuo.SetPreparesID(*id)
	}
	return uuo
}

// SetPrepares sets the "prepares" edge to the MEMOS entity.
func (uuo *USERSUpdateOne) SetPrepares(m *MEMOS) *USERSUpdateOne {
	return uuo.SetPreparesID(m.ID)
}

// AddAcquireIDs adds the "acquires" edge to the ACHIEVEMENTS entity by IDs.
func (uuo *USERSUpdateOne) AddAcquireIDs(ids ...int) *USERSUpdateOne {
	uuo.mutation.AddAcquireIDs(ids...)
	return uuo
}

// AddAcquires adds the "acquires" edges to the ACHIEVEMENTS entity.
func (uuo *USERSUpdateOne) AddAcquires(a ...*ACHIEVEMENTS) *USERSUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAcquireIDs(ids...)
}

// SetRecordsID sets the "records" edge to the PROGRESS entity by ID.
func (uuo *USERSUpdateOne) SetRecordsID(id int) *USERSUpdateOne {
	uuo.mutation.SetRecordsID(id)
	return uuo
}

// SetNillableRecordsID sets the "records" edge to the PROGRESS entity by ID if the given value is not nil.
func (uuo *USERSUpdateOne) SetNillableRecordsID(id *int) *USERSUpdateOne {
	if id != nil {
		uuo = uuo.SetRecordsID(*id)
	}
	return uuo
}

// SetRecords sets the "records" edge to the PROGRESS entity.
func (uuo *USERSUpdateOne) SetRecords(p *PROGRESS) *USERSUpdateOne {
	return uuo.SetRecordsID(p.ID)
}

// Mutation returns the USERSMutation object of the builder.
func (uuo *USERSUpdateOne) Mutation() *USERSMutation {
	return uuo.mutation
}

// ClearConnects clears all "connects" edges to the FRIENDS entity.
func (uuo *USERSUpdateOne) ClearConnects() *USERSUpdateOne {
	uuo.mutation.ClearConnects()
	return uuo
}

// RemoveConnectIDs removes the "connects" edge to FRIENDS entities by IDs.
func (uuo *USERSUpdateOne) RemoveConnectIDs(ids ...int) *USERSUpdateOne {
	uuo.mutation.RemoveConnectIDs(ids...)
	return uuo
}

// RemoveConnects removes "connects" edges to FRIENDS entities.
func (uuo *USERSUpdateOne) RemoveConnects(f ...*FRIENDS) *USERSUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.RemoveConnectIDs(ids...)
}

// ClearMakes clears all "makes" edges to the EVENT_RECORDS entity.
func (uuo *USERSUpdateOne) ClearMakes() *USERSUpdateOne {
	uuo.mutation.ClearMakes()
	return uuo
}

// RemoveMakeIDs removes the "makes" edge to EVENT_RECORDS entities by IDs.
func (uuo *USERSUpdateOne) RemoveMakeIDs(ids ...int) *USERSUpdateOne {
	uuo.mutation.RemoveMakeIDs(ids...)
	return uuo
}

// RemoveMakes removes "makes" edges to EVENT_RECORDS entities.
func (uuo *USERSUpdateOne) RemoveMakes(e ...*EVENT_RECORDS) *USERSUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uuo.RemoveMakeIDs(ids...)
}

// ClearPrepares clears the "prepares" edge to the MEMOS entity.
func (uuo *USERSUpdateOne) ClearPrepares() *USERSUpdateOne {
	uuo.mutation.ClearPrepares()
	return uuo
}

// ClearAcquires clears all "acquires" edges to the ACHIEVEMENTS entity.
func (uuo *USERSUpdateOne) ClearAcquires() *USERSUpdateOne {
	uuo.mutation.ClearAcquires()
	return uuo
}

// RemoveAcquireIDs removes the "acquires" edge to ACHIEVEMENTS entities by IDs.
func (uuo *USERSUpdateOne) RemoveAcquireIDs(ids ...int) *USERSUpdateOne {
	uuo.mutation.RemoveAcquireIDs(ids...)
	return uuo
}

// RemoveAcquires removes "acquires" edges to ACHIEVEMENTS entities.
func (uuo *USERSUpdateOne) RemoveAcquires(a ...*ACHIEVEMENTS) *USERSUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAcquireIDs(ids...)
}

// ClearRecords clears the "records" edge to the PROGRESS entity.
func (uuo *USERSUpdateOne) ClearRecords() *USERSUpdateOne {
	uuo.mutation.ClearRecords()
	return uuo
}

// Where appends a list predicates to the USERSUpdate builder.
func (uuo *USERSUpdateOne) Where(ps ...predicate.USERS) *USERSUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *USERSUpdateOne) Select(field string, fields ...string) *USERSUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated USERS entity.
func (uuo *USERSUpdateOne) Save(ctx context.Context) (*USERS, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *USERSUpdateOne) SaveX(ctx context.Context) *USERS {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *USERSUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *USERSUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *USERSUpdateOne) check() error {
	if v, ok := uuo.mutation.Username(); ok {
		if err := users.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "USERS.username": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Role(); ok {
		if err := users.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "USERS.role": %w`, err)}
		}
	}
	return nil
}

func (uuo *USERSUpdateOne) sqlSave(ctx context.Context) (_node *USERS, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "USERS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for _, f := range fields {
			if !users.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != users.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(users.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.AvatarURL(); ok {
		_spec.SetField(users.FieldAvatarURL, field.TypeString, value)
	}
	if uuo.mutation.AvatarURLCleared() {
		_spec.ClearField(users.FieldAvatarURL, field.TypeString)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(users.FieldRole, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(users.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.IsDeleted(); ok {
		_spec.SetField(users.FieldIsDeleted, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(users.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.ConnectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   users.ConnectsTable,
			Columns: users.ConnectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friends.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedConnectsIDs(); len(nodes) > 0 && !uuo.mutation.ConnectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   users.ConnectsTable,
			Columns: users.ConnectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friends.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ConnectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   users.ConnectsTable,
			Columns: users.ConnectsPrimaryKey,
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
	if uuo.mutation.MakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.MakesTable,
			Columns: []string{users.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedMakesIDs(); len(nodes) > 0 && !uuo.mutation.MakesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.MakesTable,
			Columns: []string{users.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.MakesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.MakesTable,
			Columns: []string{users.MakesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event_records.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PreparesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   users.PreparesTable,
			Columns: []string{users.PreparesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memos.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PreparesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   users.PreparesTable,
			Columns: []string{users.PreparesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memos.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AcquiresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AcquiresTable,
			Columns: []string{users.AcquiresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAcquiresIDs(); len(nodes) > 0 && !uuo.mutation.AcquiresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AcquiresTable,
			Columns: []string{users.AcquiresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AcquiresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   users.AcquiresTable,
			Columns: []string{users.AcquiresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(achievements.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RecordsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   users.RecordsTable,
			Columns: []string{users.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RecordsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   users.RecordsTable,
			Columns: []string{users.RecordsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(progress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &USERS{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
