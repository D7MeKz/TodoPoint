// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"todopoint/banking/models/ent/bankaccount"
	"todopoint/banking/models/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BankAccountDelete is the builder for deleting a BankAccount entity.
type BankAccountDelete struct {
	config
	hooks    []Hook
	mutation *BankAccountMutation
}

// Where appends a list predicates to the BankAccountDelete builder.
func (bad *BankAccountDelete) Where(ps ...predicate.BankAccount) *BankAccountDelete {
	bad.mutation.Where(ps...)
	return bad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bad *BankAccountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, bad.sqlExec, bad.mutation, bad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bad *BankAccountDelete) ExecX(ctx context.Context) int {
	n, err := bad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bad *BankAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bankaccount.Table, sqlgraph.NewFieldSpec(bankaccount.FieldID, field.TypeInt))
	if ps := bad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bad.mutation.done = true
	return affected, err
}

// BankAccountDeleteOne is the builder for deleting a single BankAccount entity.
type BankAccountDeleteOne struct {
	bad *BankAccountDelete
}

// Where appends a list predicates to the BankAccountDelete builder.
func (bado *BankAccountDeleteOne) Where(ps ...predicate.BankAccount) *BankAccountDeleteOne {
	bado.bad.mutation.Where(ps...)
	return bado
}

// Exec executes the deletion query.
func (bado *BankAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := bado.bad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bankaccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bado *BankAccountDeleteOne) ExecX(ctx context.Context) {
	if err := bado.Exec(ctx); err != nil {
		panic(err)
	}
}
