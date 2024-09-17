// Code generated by ent, DO NOT EDIT.

package calls

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CALLS {
	return predicate.CALLS(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CALLS {
	return predicate.CALLS(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CALLS {
	return predicate.CALLS(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CALLS {
	return predicate.CALLS(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CALLS {
	return predicate.CALLS(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CALLS {
	return predicate.CALLS(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CALLS {
	return predicate.CALLS(sql.FieldLTE(FieldID, id))
}

// SessionID applies equality check predicate on the "session_id" field. It's identical to SessionIDEQ.
func SessionID(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldSessionID, v))
}

// CallStart applies equality check predicate on the "call_start" field. It's identical to CallStartEQ.
func CallStart(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldCallStart, v))
}

// CallEnd applies equality check predicate on the "call_end" field. It's identical to CallEndEQ.
func CallEnd(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldCallEnd, v))
}

// Rating applies equality check predicate on the "rating" field. It's identical to RatingEQ.
func Rating(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldRating, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldCreatedAt, v))
}

// SessionIDEQ applies the EQ predicate on the "session_id" field.
func SessionIDEQ(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldSessionID, v))
}

// SessionIDNEQ applies the NEQ predicate on the "session_id" field.
func SessionIDNEQ(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldNEQ(FieldSessionID, v))
}

// SessionIDIn applies the In predicate on the "session_id" field.
func SessionIDIn(vs ...int) predicate.CALLS {
	return predicate.CALLS(sql.FieldIn(FieldSessionID, vs...))
}

// SessionIDNotIn applies the NotIn predicate on the "session_id" field.
func SessionIDNotIn(vs ...int) predicate.CALLS {
	return predicate.CALLS(sql.FieldNotIn(FieldSessionID, vs...))
}

// SessionIDGT applies the GT predicate on the "session_id" field.
func SessionIDGT(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldGT(FieldSessionID, v))
}

// SessionIDGTE applies the GTE predicate on the "session_id" field.
func SessionIDGTE(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldGTE(FieldSessionID, v))
}

// SessionIDLT applies the LT predicate on the "session_id" field.
func SessionIDLT(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldLT(FieldSessionID, v))
}

// SessionIDLTE applies the LTE predicate on the "session_id" field.
func SessionIDLTE(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldLTE(FieldSessionID, v))
}

// CallStartEQ applies the EQ predicate on the "call_start" field.
func CallStartEQ(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldCallStart, v))
}

// CallStartNEQ applies the NEQ predicate on the "call_start" field.
func CallStartNEQ(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldNEQ(FieldCallStart, v))
}

// CallStartIn applies the In predicate on the "call_start" field.
func CallStartIn(vs ...time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldIn(FieldCallStart, vs...))
}

// CallStartNotIn applies the NotIn predicate on the "call_start" field.
func CallStartNotIn(vs ...time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldNotIn(FieldCallStart, vs...))
}

// CallStartGT applies the GT predicate on the "call_start" field.
func CallStartGT(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldGT(FieldCallStart, v))
}

// CallStartGTE applies the GTE predicate on the "call_start" field.
func CallStartGTE(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldGTE(FieldCallStart, v))
}

// CallStartLT applies the LT predicate on the "call_start" field.
func CallStartLT(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldLT(FieldCallStart, v))
}

// CallStartLTE applies the LTE predicate on the "call_start" field.
func CallStartLTE(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldLTE(FieldCallStart, v))
}

// CallEndEQ applies the EQ predicate on the "call_end" field.
func CallEndEQ(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldCallEnd, v))
}

// CallEndNEQ applies the NEQ predicate on the "call_end" field.
func CallEndNEQ(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldNEQ(FieldCallEnd, v))
}

// CallEndIn applies the In predicate on the "call_end" field.
func CallEndIn(vs ...time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldIn(FieldCallEnd, vs...))
}

// CallEndNotIn applies the NotIn predicate on the "call_end" field.
func CallEndNotIn(vs ...time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldNotIn(FieldCallEnd, vs...))
}

// CallEndGT applies the GT predicate on the "call_end" field.
func CallEndGT(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldGT(FieldCallEnd, v))
}

// CallEndGTE applies the GTE predicate on the "call_end" field.
func CallEndGTE(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldGTE(FieldCallEnd, v))
}

// CallEndLT applies the LT predicate on the "call_end" field.
func CallEndLT(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldLT(FieldCallEnd, v))
}

// CallEndLTE applies the LTE predicate on the "call_end" field.
func CallEndLTE(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldLTE(FieldCallEnd, v))
}

// RatingEQ applies the EQ predicate on the "rating" field.
func RatingEQ(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldRating, v))
}

// RatingNEQ applies the NEQ predicate on the "rating" field.
func RatingNEQ(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldNEQ(FieldRating, v))
}

// RatingIn applies the In predicate on the "rating" field.
func RatingIn(vs ...int) predicate.CALLS {
	return predicate.CALLS(sql.FieldIn(FieldRating, vs...))
}

// RatingNotIn applies the NotIn predicate on the "rating" field.
func RatingNotIn(vs ...int) predicate.CALLS {
	return predicate.CALLS(sql.FieldNotIn(FieldRating, vs...))
}

// RatingGT applies the GT predicate on the "rating" field.
func RatingGT(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldGT(FieldRating, v))
}

// RatingGTE applies the GTE predicate on the "rating" field.
func RatingGTE(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldGTE(FieldRating, v))
}

// RatingLT applies the LT predicate on the "rating" field.
func RatingLT(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldLT(FieldRating, v))
}

// RatingLTE applies the LTE predicate on the "rating" field.
func RatingLTE(v int) predicate.CALLS {
	return predicate.CALLS(sql.FieldLTE(FieldRating, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CALLS {
	return predicate.CALLS(sql.FieldLTE(FieldCreatedAt, v))
}

// HasMade applies the HasEdge predicate on the "made" edge.
func HasMade() predicate.CALLS {
	return predicate.CALLS(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MadeTable, MadeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMadeWith applies the HasEdge predicate on the "made" edge with a given conditions (other predicates).
func HasMadeWith(preds ...predicate.SESSIONS) predicate.CALLS {
	return predicate.CALLS(func(s *sql.Selector) {
		step := newMadeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CALLS) predicate.CALLS {
	return predicate.CALLS(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CALLS) predicate.CALLS {
	return predicate.CALLS(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CALLS) predicate.CALLS {
	return predicate.CALLS(sql.NotPredicates(p))
}
