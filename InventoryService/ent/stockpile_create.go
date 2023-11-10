// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"userService/m/v2/ent/ingridient"
	"userService/m/v2/ent/stockpile"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// StockpileCreate is the builder for creating a Stockpile entity.
type StockpileCreate struct {
	config
	mutation *StockpileMutation
	hooks    []Hook
}

// SetIngridientId sets the "ingridientId" field.
func (sc *StockpileCreate) SetIngridientId(u uuid.UUID) *StockpileCreate {
	sc.mutation.SetIngridientId(u)
	return sc
}

// SetOwnerId sets the "ownerId" field.
func (sc *StockpileCreate) SetOwnerId(u uuid.UUID) *StockpileCreate {
	sc.mutation.SetOwnerId(u)
	return sc
}

// SetQuantity sets the "quantity" field.
func (sc *StockpileCreate) SetQuantity(i int64) *StockpileCreate {
	sc.mutation.SetQuantity(i)
	return sc
}

// SetIngridientID sets the "ingridient" edge to the Ingridient entity by ID.
func (sc *StockpileCreate) SetIngridientID(id uuid.UUID) *StockpileCreate {
	sc.mutation.SetIngridientID(id)
	return sc
}

// SetNillableIngridientID sets the "ingridient" edge to the Ingridient entity by ID if the given value is not nil.
func (sc *StockpileCreate) SetNillableIngridientID(id *uuid.UUID) *StockpileCreate {
	if id != nil {
		sc = sc.SetIngridientID(*id)
	}
	return sc
}

// SetIngridient sets the "ingridient" edge to the Ingridient entity.
func (sc *StockpileCreate) SetIngridient(i *Ingridient) *StockpileCreate {
	return sc.SetIngridientID(i.ID)
}

// Mutation returns the StockpileMutation object of the builder.
func (sc *StockpileCreate) Mutation() *StockpileMutation {
	return sc.mutation
}

// Save creates the Stockpile in the database.
func (sc *StockpileCreate) Save(ctx context.Context) (*Stockpile, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StockpileCreate) SaveX(ctx context.Context) *Stockpile {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StockpileCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StockpileCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StockpileCreate) check() error {
	if _, ok := sc.mutation.IngridientId(); !ok {
		return &ValidationError{Name: "ingridientId", err: errors.New(`ent: missing required field "Stockpile.ingridientId"`)}
	}
	if _, ok := sc.mutation.OwnerId(); !ok {
		return &ValidationError{Name: "ownerId", err: errors.New(`ent: missing required field "Stockpile.ownerId"`)}
	}
	if _, ok := sc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Stockpile.quantity"`)}
	}
	return nil
}

func (sc *StockpileCreate) sqlSave(ctx context.Context) (*Stockpile, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StockpileCreate) createSpec() (*Stockpile, *sqlgraph.CreateSpec) {
	var (
		_node = &Stockpile{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(stockpile.Table, sqlgraph.NewFieldSpec(stockpile.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.IngridientId(); ok {
		_spec.SetField(stockpile.FieldIngridientId, field.TypeUUID, value)
		_node.IngridientId = value
	}
	if value, ok := sc.mutation.OwnerId(); ok {
		_spec.SetField(stockpile.FieldOwnerId, field.TypeUUID, value)
		_node.OwnerId = value
	}
	if value, ok := sc.mutation.Quantity(); ok {
		_spec.SetField(stockpile.FieldQuantity, field.TypeInt64, value)
		_node.Quantity = value
	}
	if nodes := sc.mutation.IngridientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   stockpile.IngridientTable,
			Columns: []string{stockpile.IngridientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ingridient.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ingridient_stock = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StockpileCreateBulk is the builder for creating many Stockpile entities in bulk.
type StockpileCreateBulk struct {
	config
	err      error
	builders []*StockpileCreate
}

// Save creates the Stockpile entities in the database.
func (scb *StockpileCreateBulk) Save(ctx context.Context) ([]*Stockpile, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stockpile, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StockpileMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StockpileCreateBulk) SaveX(ctx context.Context) []*Stockpile {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StockpileCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StockpileCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
