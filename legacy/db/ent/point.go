// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"todopoint/common/db/ent/point"
	"todopoint/common/db/ent/pointinfo"
	"todopoint/common/db/ent/subtask"
	"todopoint/common/db/ent/task"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Point is the model entity for the Point schema.
type Point struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Point holds the value of the "point" field.
	Point int `json:"point,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PointQuery when eager-loading is set.
	Edges             PointEdges `json:"edges"`
	point_info_points *int
	sub_task_point    *int
	task_point        *int
	selectValues      sql.SelectValues
}

// PointEdges holds the relations/edges for other nodes in the graph.
type PointEdges struct {
	// Subtask holds the value of the subtask edge.
	Subtask *SubTask `json:"subtask,omitempty"`
	// Task holds the value of the task edge.
	Task *Task `json:"task,omitempty"`
	// PointInfo holds the value of the point_info edge.
	PointInfo *PointInfo `json:"point_info,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// SubtaskOrErr returns the Subtask value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PointEdges) SubtaskOrErr() (*SubTask, error) {
	if e.Subtask != nil {
		return e.Subtask, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: subtask.Label}
	}
	return nil, &NotLoadedError{edge: "subtask"}
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PointEdges) TaskOrErr() (*Task, error) {
	if e.Task != nil {
		return e.Task, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: task.Label}
	}
	return nil, &NotLoadedError{edge: "task"}
}

// PointInfoOrErr returns the PointInfo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PointEdges) PointInfoOrErr() (*PointInfo, error) {
	if e.PointInfo != nil {
		return e.PointInfo, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: pointinfo.Label}
	}
	return nil, &NotLoadedError{edge: "point_info"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Point) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case point.FieldID, point.FieldPoint:
			values[i] = new(sql.NullInt64)
		case point.FieldType:
			values[i] = new(sql.NullString)
		case point.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case point.ForeignKeys[0]: // point_info_points
			values[i] = new(sql.NullInt64)
		case point.ForeignKeys[1]: // sub_task_point
			values[i] = new(sql.NullInt64)
		case point.ForeignKeys[2]: // task_point
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Point fields.
func (po *Point) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case point.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case point.FieldPoint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field point", values[i])
			} else if value.Valid {
				po.Point = int(value.Int64)
			}
		case point.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				po.Type = value.String
			}
		case point.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case point.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field point_info_points", value)
			} else if value.Valid {
				po.point_info_points = new(int)
				*po.point_info_points = int(value.Int64)
			}
		case point.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field sub_task_point", value)
			} else if value.Valid {
				po.sub_task_point = new(int)
				*po.sub_task_point = int(value.Int64)
			}
		case point.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_point", value)
			} else if value.Valid {
				po.task_point = new(int)
				*po.task_point = int(value.Int64)
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Point.
// This includes values selected through modifiers, order, etc.
func (po *Point) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QuerySubtask queries the "subtask" edge of the Point entity.
func (po *Point) QuerySubtask() *SubTaskQuery {
	return NewPointClient(po.config).QuerySubtask(po)
}

// QueryTask queries the "task" edge of the Point entity.
func (po *Point) QueryTask() *TaskQuery {
	return NewPointClient(po.config).QueryTask(po)
}

// QueryPointInfo queries the "point_info" edge of the Point entity.
func (po *Point) QueryPointInfo() *PointInfoQuery {
	return NewPointClient(po.config).QueryPointInfo(po)
}

// Update returns a builder for updating this Point.
// Note that you need to call Point.Unwrap() before calling this method if this Point
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Point) Update() *PointUpdateOne {
	return NewPointClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Point entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the config which created the transaction.
func (po *Point) Unwrap() *Point {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Point is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Point) String() string {
	var builder strings.Builder
	builder.WriteString("Point(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("point=")
	builder.WriteString(fmt.Sprintf("%v", po.Point))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(po.Type)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Points is a parsable slice of Point.
type Points []*Point
