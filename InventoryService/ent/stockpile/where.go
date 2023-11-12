// Code generated by ent, DO NOT EDIT.

package stockpile

import (
	"inventoryService/m/v2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldLTE(FieldID, id))
}

// IngridientId applies equality check predicate on the "ingridientId" field. It's identical to IngridientIdEQ.
func IngridientId(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldEQ(FieldIngridientId, v))
}

// OwnerId applies equality check predicate on the "ownerId" field. It's identical to OwnerIdEQ.
func OwnerId(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldEQ(FieldOwnerId, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldEQ(FieldQuantity, v))
}

// IngridientIdEQ applies the EQ predicate on the "ingridientId" field.
func IngridientIdEQ(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldEQ(FieldIngridientId, v))
}

// IngridientIdNEQ applies the NEQ predicate on the "ingridientId" field.
func IngridientIdNEQ(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldNEQ(FieldIngridientId, v))
}

// IngridientIdIn applies the In predicate on the "ingridientId" field.
func IngridientIdIn(vs ...uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldIn(FieldIngridientId, vs...))
}

// IngridientIdNotIn applies the NotIn predicate on the "ingridientId" field.
func IngridientIdNotIn(vs ...uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldNotIn(FieldIngridientId, vs...))
}

// IngridientIdGT applies the GT predicate on the "ingridientId" field.
func IngridientIdGT(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldGT(FieldIngridientId, v))
}

// IngridientIdGTE applies the GTE predicate on the "ingridientId" field.
func IngridientIdGTE(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldGTE(FieldIngridientId, v))
}

// IngridientIdLT applies the LT predicate on the "ingridientId" field.
func IngridientIdLT(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldLT(FieldIngridientId, v))
}

// IngridientIdLTE applies the LTE predicate on the "ingridientId" field.
func IngridientIdLTE(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldLTE(FieldIngridientId, v))
}

// OwnerIdEQ applies the EQ predicate on the "ownerId" field.
func OwnerIdEQ(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldEQ(FieldOwnerId, v))
}

// OwnerIdNEQ applies the NEQ predicate on the "ownerId" field.
func OwnerIdNEQ(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldNEQ(FieldOwnerId, v))
}

// OwnerIdIn applies the In predicate on the "ownerId" field.
func OwnerIdIn(vs ...uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldIn(FieldOwnerId, vs...))
}

// OwnerIdNotIn applies the NotIn predicate on the "ownerId" field.
func OwnerIdNotIn(vs ...uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldNotIn(FieldOwnerId, vs...))
}

// OwnerIdGT applies the GT predicate on the "ownerId" field.
func OwnerIdGT(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldGT(FieldOwnerId, v))
}

// OwnerIdGTE applies the GTE predicate on the "ownerId" field.
func OwnerIdGTE(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldGTE(FieldOwnerId, v))
}

// OwnerIdLT applies the LT predicate on the "ownerId" field.
func OwnerIdLT(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldLT(FieldOwnerId, v))
}

// OwnerIdLTE applies the LTE predicate on the "ownerId" field.
func OwnerIdLTE(v uuid.UUID) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldLTE(FieldOwnerId, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int64) predicate.Stockpile {
	return predicate.Stockpile(sql.FieldLTE(FieldQuantity, v))
}

// HasIngridient applies the HasEdge predicate on the "ingridient" edge.
func HasIngridient() predicate.Stockpile {
	return predicate.Stockpile(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, IngridientTable, IngridientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIngridientWith applies the HasEdge predicate on the "ingridient" edge with a given conditions (other predicates).
func HasIngridientWith(preds ...predicate.Ingridient) predicate.Stockpile {
	return predicate.Stockpile(func(s *sql.Selector) {
		step := newIngridientStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Stockpile) predicate.Stockpile {
	return predicate.Stockpile(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Stockpile) predicate.Stockpile {
	return predicate.Stockpile(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Stockpile) predicate.Stockpile {
	return predicate.Stockpile(sql.NotPredicates(p))
}