// Code generated by ent, DO NOT EDIT.

package progress

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Hosi121/SpeakUp/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldUserID, v))
}

// LoginDays applies equality check predicate on the "login_days" field. It's identical to LoginDaysEQ.
func LoginDays(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldLoginDays, v))
}

// ConsecutiveParticipants applies equality check predicate on the "consecutive_participants" field. It's identical to ConsecutiveParticipantsEQ.
func ConsecutiveParticipants(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldConsecutiveParticipants, v))
}

// ConsecutiveLoginDays applies equality check predicate on the "consecutive_login_days" field. It's identical to ConsecutiveLoginDaysEQ.
func ConsecutiveLoginDays(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldConsecutiveLoginDays, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLTE(FieldUserID, v))
}

// LoginDaysEQ applies the EQ predicate on the "login_days" field.
func LoginDaysEQ(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldLoginDays, v))
}

// LoginDaysNEQ applies the NEQ predicate on the "login_days" field.
func LoginDaysNEQ(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNEQ(FieldLoginDays, v))
}

// LoginDaysIn applies the In predicate on the "login_days" field.
func LoginDaysIn(vs ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldIn(FieldLoginDays, vs...))
}

// LoginDaysNotIn applies the NotIn predicate on the "login_days" field.
func LoginDaysNotIn(vs ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNotIn(FieldLoginDays, vs...))
}

// LoginDaysGT applies the GT predicate on the "login_days" field.
func LoginDaysGT(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGT(FieldLoginDays, v))
}

// LoginDaysGTE applies the GTE predicate on the "login_days" field.
func LoginDaysGTE(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGTE(FieldLoginDays, v))
}

// LoginDaysLT applies the LT predicate on the "login_days" field.
func LoginDaysLT(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLT(FieldLoginDays, v))
}

// LoginDaysLTE applies the LTE predicate on the "login_days" field.
func LoginDaysLTE(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLTE(FieldLoginDays, v))
}

// ConsecutiveParticipantsEQ applies the EQ predicate on the "consecutive_participants" field.
func ConsecutiveParticipantsEQ(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldConsecutiveParticipants, v))
}

// ConsecutiveParticipantsNEQ applies the NEQ predicate on the "consecutive_participants" field.
func ConsecutiveParticipantsNEQ(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNEQ(FieldConsecutiveParticipants, v))
}

// ConsecutiveParticipantsIn applies the In predicate on the "consecutive_participants" field.
func ConsecutiveParticipantsIn(vs ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldIn(FieldConsecutiveParticipants, vs...))
}

// ConsecutiveParticipantsNotIn applies the NotIn predicate on the "consecutive_participants" field.
func ConsecutiveParticipantsNotIn(vs ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNotIn(FieldConsecutiveParticipants, vs...))
}

// ConsecutiveParticipantsGT applies the GT predicate on the "consecutive_participants" field.
func ConsecutiveParticipantsGT(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGT(FieldConsecutiveParticipants, v))
}

// ConsecutiveParticipantsGTE applies the GTE predicate on the "consecutive_participants" field.
func ConsecutiveParticipantsGTE(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGTE(FieldConsecutiveParticipants, v))
}

// ConsecutiveParticipantsLT applies the LT predicate on the "consecutive_participants" field.
func ConsecutiveParticipantsLT(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLT(FieldConsecutiveParticipants, v))
}

// ConsecutiveParticipantsLTE applies the LTE predicate on the "consecutive_participants" field.
func ConsecutiveParticipantsLTE(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLTE(FieldConsecutiveParticipants, v))
}

// ConsecutiveLoginDaysEQ applies the EQ predicate on the "consecutive_login_days" field.
func ConsecutiveLoginDaysEQ(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldEQ(FieldConsecutiveLoginDays, v))
}

// ConsecutiveLoginDaysNEQ applies the NEQ predicate on the "consecutive_login_days" field.
func ConsecutiveLoginDaysNEQ(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNEQ(FieldConsecutiveLoginDays, v))
}

// ConsecutiveLoginDaysIn applies the In predicate on the "consecutive_login_days" field.
func ConsecutiveLoginDaysIn(vs ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldIn(FieldConsecutiveLoginDays, vs...))
}

// ConsecutiveLoginDaysNotIn applies the NotIn predicate on the "consecutive_login_days" field.
func ConsecutiveLoginDaysNotIn(vs ...int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldNotIn(FieldConsecutiveLoginDays, vs...))
}

// ConsecutiveLoginDaysGT applies the GT predicate on the "consecutive_login_days" field.
func ConsecutiveLoginDaysGT(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGT(FieldConsecutiveLoginDays, v))
}

// ConsecutiveLoginDaysGTE applies the GTE predicate on the "consecutive_login_days" field.
func ConsecutiveLoginDaysGTE(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldGTE(FieldConsecutiveLoginDays, v))
}

// ConsecutiveLoginDaysLT applies the LT predicate on the "consecutive_login_days" field.
func ConsecutiveLoginDaysLT(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLT(FieldConsecutiveLoginDays, v))
}

// ConsecutiveLoginDaysLTE applies the LTE predicate on the "consecutive_login_days" field.
func ConsecutiveLoginDaysLTE(v int) predicate.PROGRESS {
	return predicate.PROGRESS(sql.FieldLTE(FieldConsecutiveLoginDays, v))
}

// HasRecorded applies the HasEdge predicate on the "recorded" edge.
func HasRecorded() predicate.PROGRESS {
	return predicate.PROGRESS(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, RecordedTable, RecordedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecordedWith applies the HasEdge predicate on the "recorded" edge with a given conditions (other predicates).
func HasRecordedWith(preds ...predicate.USERS) predicate.PROGRESS {
	return predicate.PROGRESS(func(s *sql.Selector) {
		step := newRecordedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PROGRESS) predicate.PROGRESS {
	return predicate.PROGRESS(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PROGRESS) predicate.PROGRESS {
	return predicate.PROGRESS(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PROGRESS) predicate.PROGRESS {
	return predicate.PROGRESS(sql.NotPredicates(p))
}