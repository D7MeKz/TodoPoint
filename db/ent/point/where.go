// Code generated by ent, DO NOT EDIT.

package point

import (
	"common/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Point {
	return predicate.Point(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Point {
	return predicate.Point(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Point {
	return predicate.Point(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Point {
	return predicate.Point(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Point {
	return predicate.Point(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Point {
	return predicate.Point(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Point {
	return predicate.Point(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Point {
	return predicate.Point(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Point {
	return predicate.Point(sql.FieldLTE(FieldID, id))
}

// Point applies equality check predicate on the "point" field. It's identical to PointEQ.
func Point(v int) predicate.Point {
	return predicate.Point(sql.FieldEQ(FieldPoint, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Point {
	return predicate.Point(sql.FieldEQ(FieldType, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Point {
	return predicate.Point(sql.FieldEQ(FieldCreatedAt, v))
}

// PointEQ applies the EQ predicate on the "point" field.
func PointEQ(v int) predicate.Point {
	return predicate.Point(sql.FieldEQ(FieldPoint, v))
}

// PointNEQ applies the NEQ predicate on the "point" field.
func PointNEQ(v int) predicate.Point {
	return predicate.Point(sql.FieldNEQ(FieldPoint, v))
}

// PointIn applies the In predicate on the "point" field.
func PointIn(vs ...int) predicate.Point {
	return predicate.Point(sql.FieldIn(FieldPoint, vs...))
}

// PointNotIn applies the NotIn predicate on the "point" field.
func PointNotIn(vs ...int) predicate.Point {
	return predicate.Point(sql.FieldNotIn(FieldPoint, vs...))
}

// PointGT applies the GT predicate on the "point" field.
func PointGT(v int) predicate.Point {
	return predicate.Point(sql.FieldGT(FieldPoint, v))
}

// PointGTE applies the GTE predicate on the "point" field.
func PointGTE(v int) predicate.Point {
	return predicate.Point(sql.FieldGTE(FieldPoint, v))
}

// PointLT applies the LT predicate on the "point" field.
func PointLT(v int) predicate.Point {
	return predicate.Point(sql.FieldLT(FieldPoint, v))
}

// PointLTE applies the LTE predicate on the "point" field.
func PointLTE(v int) predicate.Point {
	return predicate.Point(sql.FieldLTE(FieldPoint, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Point {
	return predicate.Point(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Point {
	return predicate.Point(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Point {
	return predicate.Point(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Point {
	return predicate.Point(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Point {
	return predicate.Point(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Point {
	return predicate.Point(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Point {
	return predicate.Point(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Point {
	return predicate.Point(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Point {
	return predicate.Point(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Point {
	return predicate.Point(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Point {
	return predicate.Point(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Point {
	return predicate.Point(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Point {
	return predicate.Point(sql.FieldContainsFold(FieldType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Point {
	return predicate.Point(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Point {
	return predicate.Point(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Point {
	return predicate.Point(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Point {
	return predicate.Point(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Point {
	return predicate.Point(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Point {
	return predicate.Point(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Point {
	return predicate.Point(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Point {
	return predicate.Point(sql.FieldLTE(FieldCreatedAt, v))
}

// HasSubtask applies the HasEdge predicate on the "subtask" edge.
func HasSubtask() predicate.Point {
	return predicate.Point(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, SubtaskTable, SubtaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubtaskWith applies the HasEdge predicate on the "subtask" edge with a given conditions (other predicates).
func HasSubtaskWith(preds ...predicate.SubTask) predicate.Point {
	return predicate.Point(func(s *sql.Selector) {
		step := newSubtaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.Point {
	return predicate.Point(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.Point {
	return predicate.Point(func(s *sql.Selector) {
		step := newTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPointInfo applies the HasEdge predicate on the "point_info" edge.
func HasPointInfo() predicate.Point {
	return predicate.Point(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PointInfoTable, PointInfoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPointInfoWith applies the HasEdge predicate on the "point_info" edge with a given conditions (other predicates).
func HasPointInfoWith(preds ...predicate.PointInfo) predicate.Point {
	return predicate.Point(func(s *sql.Selector) {
		step := newPointInfoStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Point) predicate.Point {
	return predicate.Point(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Point) predicate.Point {
	return predicate.Point(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Point) predicate.Point {
	return predicate.Point(sql.NotPredicates(p))
}
