// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"inventoryService/m/v2/ent/predicate"
	"inventoryService/m/v2/ent/stockpile"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockpileDelete is the builder for deleting a Stockpile entity.
type StockpileDelete struct {
	config
	hooks    []Hook
	mutation *StockpileMutation
}

// Where appends a list predicates to the StockpileDelete builder.
func (sd *StockpileDelete) Where(ps ...predicate.Stockpile) *StockpileDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *StockpileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *StockpileDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *StockpileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(stockpile.Table, sqlgraph.NewFieldSpec(stockpile.FieldID, field.TypeInt))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// StockpileDeleteOne is the builder for deleting a single Stockpile entity.
type StockpileDeleteOne struct {
	sd *StockpileDelete
}

// Where appends a list predicates to the StockpileDelete builder.
func (sdo *StockpileDeleteOne) Where(ps ...predicate.Stockpile) *StockpileDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *StockpileDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{stockpile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *StockpileDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}