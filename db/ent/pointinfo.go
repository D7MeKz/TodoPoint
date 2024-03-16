// Code generated by ent, DO NOT EDIT.

package ent

import (
	"common/db/ent/member"
	"common/db/ent/pointinfo"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PointInfo is the model entity for the PointInfo schema.
type PointInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TotalPoints holds the value of the "total_points" field.
	TotalPoints int64 `json:"total_points,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PointInfoQuery when eager-loading is set.
	Edges             PointInfoEdges `json:"edges"`
	member_point_info *int
	selectValues      sql.SelectValues
}

// PointInfoEdges holds the relations/edges for other nodes in the graph.
type PointInfoEdges struct {
	// Points holds the value of the points edge.
	Points []*Point `json:"points,omitempty"`
	// UserID holds the value of the user_id edge.
	UserID *Member `json:"user_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PointsOrErr returns the Points value or an error if the edge
// was not loaded in eager-loading.
func (e PointInfoEdges) PointsOrErr() ([]*Point, error) {
	if e.loadedTypes[0] {
		return e.Points, nil
	}
	return nil, &NotLoadedError{edge: "points"}
}

// UserIDOrErr returns the UserID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PointInfoEdges) UserIDOrErr() (*Member, error) {
	if e.UserID != nil {
		return e.UserID, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: member.Label}
	}
	return nil, &NotLoadedError{edge: "user_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PointInfo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pointinfo.FieldID, pointinfo.FieldTotalPoints:
			values[i] = new(sql.NullInt64)
		case pointinfo.FieldModifiedAt:
			values[i] = new(sql.NullTime)
		case pointinfo.ForeignKeys[0]: // member_point_info
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PointInfo fields.
func (pi *PointInfo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pointinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = int(value.Int64)
		case pointinfo.FieldTotalPoints:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_points", values[i])
			} else if value.Valid {
				pi.TotalPoints = value.Int64
			}
		case pointinfo.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				pi.ModifiedAt = value.Time
			}
		case pointinfo.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field member_point_info", value)
			} else if value.Valid {
				pi.member_point_info = new(int)
				*pi.member_point_info = int(value.Int64)
			}
		default:
			pi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PointInfo.
// This includes values selected through modifiers, order, etc.
func (pi *PointInfo) Value(name string) (ent.Value, error) {
	return pi.selectValues.Get(name)
}

// QueryPoints queries the "points" edge of the PointInfo entity.
func (pi *PointInfo) QueryPoints() *PointQuery {
	return NewPointInfoClient(pi.config).QueryPoints(pi)
}

// QueryUserID queries the "user_id" edge of the PointInfo entity.
func (pi *PointInfo) QueryUserID() *MemberQuery {
	return NewPointInfoClient(pi.config).QueryUserID(pi)
}

// Update returns a builder for updating this PointInfo.
// Note that you need to call PointInfo.Unwrap() before calling this method if this PointInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *PointInfo) Update() *PointInfoUpdateOne {
	return NewPointInfoClient(pi.config).UpdateOne(pi)
}

// Unwrap unwraps the PointInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *PointInfo) Unwrap() *PointInfo {
	_tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: PointInfo is not a transactional entity")
	}
	pi.config.driver = _tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *PointInfo) String() string {
	var builder strings.Builder
	builder.WriteString("PointInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pi.ID))
	builder.WriteString("total_points=")
	builder.WriteString(fmt.Sprintf("%v", pi.TotalPoints))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(pi.ModifiedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// PointInfos is a parsable slice of PointInfo.
type PointInfos []*PointInfo
