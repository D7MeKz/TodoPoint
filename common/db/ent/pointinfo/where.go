// Code generated by ent, DO NOT EDIT.

package pointinfo

import (
	"time"
	"todopoint/common/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldLTE(FieldID, id))
}

// TotalPoints applies equality check predicate on the "total_points" field. It's identical to TotalPointsEQ.
func TotalPoints(v int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldEQ(FieldTotalPoints, v))
}

// ModifiedAt applies equality check predicate on the "modified_at" field. It's identical to ModifiedAtEQ.
func ModifiedAt(v time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldEQ(FieldModifiedAt, v))
}

// TotalPointsEQ applies the EQ predicate on the "total_points" field.
func TotalPointsEQ(v int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldEQ(FieldTotalPoints, v))
}

// TotalPointsNEQ applies the NEQ predicate on the "total_points" field.
func TotalPointsNEQ(v int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldNEQ(FieldTotalPoints, v))
}

// TotalPointsIn applies the In predicate on the "total_points" field.
func TotalPointsIn(vs ...int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldIn(FieldTotalPoints, vs...))
}

// TotalPointsNotIn applies the NotIn predicate on the "total_points" field.
func TotalPointsNotIn(vs ...int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldNotIn(FieldTotalPoints, vs...))
}

// TotalPointsGT applies the GT predicate on the "total_points" field.
func TotalPointsGT(v int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldGT(FieldTotalPoints, v))
}

// TotalPointsGTE applies the GTE predicate on the "total_points" field.
func TotalPointsGTE(v int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldGTE(FieldTotalPoints, v))
}

// TotalPointsLT applies the LT predicate on the "total_points" field.
func TotalPointsLT(v int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldLT(FieldTotalPoints, v))
}

// TotalPointsLTE applies the LTE predicate on the "total_points" field.
func TotalPointsLTE(v int64) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldLTE(FieldTotalPoints, v))
}

// ModifiedAtEQ applies the EQ predicate on the "modified_at" field.
func ModifiedAtEQ(v time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldEQ(FieldModifiedAt, v))
}

// ModifiedAtNEQ applies the NEQ predicate on the "modified_at" field.
func ModifiedAtNEQ(v time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldNEQ(FieldModifiedAt, v))
}

// ModifiedAtIn applies the In predicate on the "modified_at" field.
func ModifiedAtIn(vs ...time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldIn(FieldModifiedAt, vs...))
}

// ModifiedAtNotIn applies the NotIn predicate on the "modified_at" field.
func ModifiedAtNotIn(vs ...time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldNotIn(FieldModifiedAt, vs...))
}

// ModifiedAtGT applies the GT predicate on the "modified_at" field.
func ModifiedAtGT(v time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldGT(FieldModifiedAt, v))
}

// ModifiedAtGTE applies the GTE predicate on the "modified_at" field.
func ModifiedAtGTE(v time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldGTE(FieldModifiedAt, v))
}

// ModifiedAtLT applies the LT predicate on the "modified_at" field.
func ModifiedAtLT(v time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldLT(FieldModifiedAt, v))
}

// ModifiedAtLTE applies the LTE predicate on the "modified_at" field.
func ModifiedAtLTE(v time.Time) predicate.PointInfo {
	return predicate.PointInfo(sql.FieldLTE(FieldModifiedAt, v))
}

// HasPoints applies the HasEdge predicate on the "points" edge.
func HasPoints() predicate.PointInfo {
	return predicate.PointInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PointsTable, PointsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPointsWith applies the HasEdge predicate on the "points" edge with a given conditions (other predicates).
func HasPointsWith(preds ...predicate.Point) predicate.PointInfo {
	return predicate.PointInfo(func(s *sql.Selector) {
		step := newPointsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserID applies the HasEdge predicate on the "user_id" edge.
func HasUserID() predicate.PointInfo {
	return predicate.PointInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserIDTable, UserIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserIDWith applies the HasEdge predicate on the "user_id" edge with a given conditions (other predicates).
func HasUserIDWith(preds ...predicate.Member) predicate.PointInfo {
	return predicate.PointInfo(func(s *sql.Selector) {
		step := newUserIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PointInfo) predicate.PointInfo {
	return predicate.PointInfo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PointInfo) predicate.PointInfo {
	return predicate.PointInfo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PointInfo) predicate.PointInfo {
	return predicate.PointInfo(sql.NotPredicates(p))
}