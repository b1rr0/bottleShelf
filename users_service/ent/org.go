// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"users_service/ent/org"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Org is the model entity for the Org schema.
type Org struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Org) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case org.FieldName:
			values[i] = new(sql.NullString)
		case org.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Org fields.
func (o *Org) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case org.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		case org.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				o.Name = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Org.
// This includes values selected through modifiers, order, etc.
func (o *Org) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// Update returns a builder for updating this Org.
// Note that you need to call Org.Unwrap() before calling this method if this Org
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Org) Update() *OrgUpdateOne {
	return NewOrgClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Org entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Org) Unwrap() *Org {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Org is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Org) String() string {
	var builder strings.Builder
	builder.WriteString("Org(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("name=")
	builder.WriteString(o.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Orgs is a parsable slice of Org.
type Orgs []*Org
