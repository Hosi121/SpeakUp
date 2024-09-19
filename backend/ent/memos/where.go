// Code generated by ent, DO NOT EDIT.

package memos

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEQ(FieldUserID, v))
}

// Memo1 applies equality check predicate on the "memo1" field. It's identical to Memo1EQ.
func Memo1(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEQ(FieldMemo1, v))
}

// Memo2 applies equality check predicate on the "memo2" field. It's identical to Memo2EQ.
func Memo2(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEQ(FieldMemo2, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldLTE(FieldUserID, v))
}

// Memo1EQ applies the EQ predicate on the "memo1" field.
func Memo1EQ(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEQ(FieldMemo1, v))
}

// Memo1NEQ applies the NEQ predicate on the "memo1" field.
func Memo1NEQ(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldNEQ(FieldMemo1, v))
}

// Memo1In applies the In predicate on the "memo1" field.
func Memo1In(vs ...string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldIn(FieldMemo1, vs...))
}

// Memo1NotIn applies the NotIn predicate on the "memo1" field.
func Memo1NotIn(vs ...string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldNotIn(FieldMemo1, vs...))
}

// Memo1GT applies the GT predicate on the "memo1" field.
func Memo1GT(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldGT(FieldMemo1, v))
}

// Memo1GTE applies the GTE predicate on the "memo1" field.
func Memo1GTE(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldGTE(FieldMemo1, v))
}

// Memo1LT applies the LT predicate on the "memo1" field.
func Memo1LT(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldLT(FieldMemo1, v))
}

// Memo1LTE applies the LTE predicate on the "memo1" field.
func Memo1LTE(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldLTE(FieldMemo1, v))
}

// Memo1Contains applies the Contains predicate on the "memo1" field.
func Memo1Contains(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldContains(FieldMemo1, v))
}

// Memo1HasPrefix applies the HasPrefix predicate on the "memo1" field.
func Memo1HasPrefix(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldHasPrefix(FieldMemo1, v))
}

// Memo1HasSuffix applies the HasSuffix predicate on the "memo1" field.
func Memo1HasSuffix(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldHasSuffix(FieldMemo1, v))
}

// Memo1EqualFold applies the EqualFold predicate on the "memo1" field.
func Memo1EqualFold(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEqualFold(FieldMemo1, v))
}

// Memo1ContainsFold applies the ContainsFold predicate on the "memo1" field.
func Memo1ContainsFold(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldContainsFold(FieldMemo1, v))
}

// Memo2EQ applies the EQ predicate on the "memo2" field.
func Memo2EQ(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEQ(FieldMemo2, v))
}

// Memo2NEQ applies the NEQ predicate on the "memo2" field.
func Memo2NEQ(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldNEQ(FieldMemo2, v))
}

// Memo2In applies the In predicate on the "memo2" field.
func Memo2In(vs ...string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldIn(FieldMemo2, vs...))
}

// Memo2NotIn applies the NotIn predicate on the "memo2" field.
func Memo2NotIn(vs ...string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldNotIn(FieldMemo2, vs...))
}

// Memo2GT applies the GT predicate on the "memo2" field.
func Memo2GT(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldGT(FieldMemo2, v))
}

// Memo2GTE applies the GTE predicate on the "memo2" field.
func Memo2GTE(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldGTE(FieldMemo2, v))
}

// Memo2LT applies the LT predicate on the "memo2" field.
func Memo2LT(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldLT(FieldMemo2, v))
}

// Memo2LTE applies the LTE predicate on the "memo2" field.
func Memo2LTE(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldLTE(FieldMemo2, v))
}

// Memo2Contains applies the Contains predicate on the "memo2" field.
func Memo2Contains(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldContains(FieldMemo2, v))
}

// Memo2HasPrefix applies the HasPrefix predicate on the "memo2" field.
func Memo2HasPrefix(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldHasPrefix(FieldMemo2, v))
}

// Memo2HasSuffix applies the HasSuffix predicate on the "memo2" field.
func Memo2HasSuffix(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldHasSuffix(FieldMemo2, v))
}

// Memo2EqualFold applies the EqualFold predicate on the "memo2" field.
func Memo2EqualFold(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldEqualFold(FieldMemo2, v))
}

// Memo2ContainsFold applies the ContainsFold predicate on the "memo2" field.
func Memo2ContainsFold(v string) predicate.MEMOS {
	return predicate.MEMOS(sql.FieldContainsFold(FieldMemo2, v))
}

// HasPrepared applies the HasEdge predicate on the "prepared" edge.
func HasPrepared() predicate.MEMOS {
	return predicate.MEMOS(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PreparedTable, PreparedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPreparedWith applies the HasEdge predicate on the "prepared" edge with a given conditions (other predicates).
func HasPreparedWith(preds ...predicate.USERS) predicate.MEMOS {
	return predicate.MEMOS(func(s *sql.Selector) {
		step := newPreparedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MEMOS) predicate.MEMOS {
	return predicate.MEMOS(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MEMOS) predicate.MEMOS {
	return predicate.MEMOS(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MEMOS) predicate.MEMOS {
	return predicate.MEMOS(sql.NotPredicates(p))
}
