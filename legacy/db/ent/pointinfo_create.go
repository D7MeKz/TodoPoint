// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todopoint/common/db/ent/member"
	"todopoint/common/db/ent/point"
	"todopoint/common/db/ent/pointinfo"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PointInfoCreate is the builder for creating a PointInfo entity.
type PointInfoCreate struct {
	config
	mutation *PointInfoMutation
	hooks    []Hook
}

// SetTotalPoints sets the "total_points" field.
func (pic *PointInfoCreate) SetTotalPoints(i int64) *PointInfoCreate {
	pic.mutation.SetTotalPoints(i)
	return pic
}

// SetModifiedAt sets the "modified_at" field.
func (pic *PointInfoCreate) SetModifiedAt(t time.Time) *PointInfoCreate {
	pic.mutation.SetModifiedAt(t)
	return pic
}

// SetNillableModifiedAt sets the "modified_at" field if the given value is not nil.
func (pic *PointInfoCreate) SetNillableModifiedAt(t *time.Time) *PointInfoCreate {
	if t != nil {
		pic.SetModifiedAt(*t)
	}
	return pic
}

// SetID sets the "id" field.
func (pic *PointInfoCreate) SetID(i int) *PointInfoCreate {
	pic.mutation.SetID(i)
	return pic
}

// AddPointIDs adds the "points" edge to the Point entity by IDs.
func (pic *PointInfoCreate) AddPointIDs(ids ...int) *PointInfoCreate {
	pic.mutation.AddPointIDs(ids...)
	return pic
}

// AddPoints adds the "points" edges to the Point entity.
func (pic *PointInfoCreate) AddPoints(p ...*Point) *PointInfoCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pic.AddPointIDs(ids...)
}

// SetUserIDID sets the "user_id" edge to the Member entity by ID.
func (pic *PointInfoCreate) SetUserIDID(id int) *PointInfoCreate {
	pic.mutation.SetUserIDID(id)
	return pic
}

// SetNillableUserIDID sets the "user_id" edge to the Member entity by ID if the given value is not nil.
func (pic *PointInfoCreate) SetNillableUserIDID(id *int) *PointInfoCreate {
	if id != nil {
		pic = pic.SetUserIDID(*id)
	}
	return pic
}

// SetUserID sets the "user_id" edge to the Member entity.
func (pic *PointInfoCreate) SetUserID(m *Member) *PointInfoCreate {
	return pic.SetUserIDID(m.ID)
}

// Mutation returns the PointInfoMutation object of the builder.
func (pic *PointInfoCreate) Mutation() *PointInfoMutation {
	return pic.mutation
}

// Save creates the PointInfo in the database.
func (pic *PointInfoCreate) Save(ctx context.Context) (*PointInfo, error) {
	pic.defaults()
	return withHooks(ctx, pic.sqlSave, pic.mutation, pic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pic *PointInfoCreate) SaveX(ctx context.Context) *PointInfo {
	v, err := pic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pic *PointInfoCreate) Exec(ctx context.Context) error {
	_, err := pic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pic *PointInfoCreate) ExecX(ctx context.Context) {
	if err := pic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pic *PointInfoCreate) defaults() {
	if _, ok := pic.mutation.ModifiedAt(); !ok {
		v := pointinfo.DefaultModifiedAt
		pic.mutation.SetModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pic *PointInfoCreate) check() error {
	if _, ok := pic.mutation.TotalPoints(); !ok {
		return &ValidationError{Name: "total_points", err: errors.New(`ent: missing required field "PointInfo.total_points"`)}
	}
	if _, ok := pic.mutation.ModifiedAt(); !ok {
		return &ValidationError{Name: "modified_at", err: errors.New(`ent: missing required field "PointInfo.modified_at"`)}
	}
	return nil
}

func (pic *PointInfoCreate) sqlSave(ctx context.Context) (*PointInfo, error) {
	if err := pic.check(); err != nil {
		return nil, err
	}
	_node, _spec := pic.createSpec()
	if err := sqlgraph.CreateNode(ctx, pic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pic.mutation.id = &_node.ID
	pic.mutation.done = true
	return _node, nil
}

func (pic *PointInfoCreate) createSpec() (*PointInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &PointInfo{config: pic.config}
		_spec = sqlgraph.NewCreateSpec(pointinfo.Table, sqlgraph.NewFieldSpec(pointinfo.FieldID, field.TypeInt))
	)
	if id, ok := pic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pic.mutation.TotalPoints(); ok {
		_spec.SetField(pointinfo.FieldTotalPoints, field.TypeInt64, value)
		_node.TotalPoints = value
	}
	if value, ok := pic.mutation.ModifiedAt(); ok {
		_spec.SetField(pointinfo.FieldModifiedAt, field.TypeTime, value)
		_node.ModifiedAt = value
	}
	if nodes := pic.mutation.PointsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pointinfo.PointsTable,
			Columns: []string{pointinfo.PointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(point.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pic.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pointinfo.UserIDTable,
			Columns: []string{pointinfo.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.member_point_info = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PointInfoCreateBulk is the builder for creating many PointInfo entities in bulk.
type PointInfoCreateBulk struct {
	config
	err      error
	builders []*PointInfoCreate
}

// Save creates the PointInfo entities in the database.
func (picb *PointInfoCreateBulk) Save(ctx context.Context) ([]*PointInfo, error) {
	if picb.err != nil {
		return nil, picb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(picb.builders))
	nodes := make([]*PointInfo, len(picb.builders))
	mutators := make([]Mutator, len(picb.builders))
	for i := range picb.builders {
		func(i int, root context.Context) {
			builder := picb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PointInfoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, picb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, picb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, picb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (picb *PointInfoCreateBulk) SaveX(ctx context.Context) []*PointInfo {
	v, err := picb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (picb *PointInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := picb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (picb *PointInfoCreateBulk) ExecX(ctx context.Context) {
	if err := picb.Exec(ctx); err != nil {
		panic(err)
	}
}