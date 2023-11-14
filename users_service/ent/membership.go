// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"users_service/ent/membership"
	"users_service/ent/org"
	"users_service/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Membership is the model entity for the Membership schema.
type Membership struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Role holds the value of the "role" field.
	Role string `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MembershipQuery when eager-loading is set.
	Edges             MembershipEdges `json:"edges"`
	membership_member *uuid.UUID
	membership_org    *uuid.UUID
	selectValues      sql.SelectValues
}

// MembershipEdges holds the relations/edges for other nodes in the graph.
type MembershipEdges struct {
	// Member holds the value of the member edge.
	Member *User `json:"member,omitempty"`
	// Org holds the value of the org edge.
	Org *Org `json:"org,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MemberOrErr returns the Member value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MembershipEdges) MemberOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Member == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Member, nil
	}
	return nil, &NotLoadedError{edge: "member"}
}

// OrgOrErr returns the Org value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MembershipEdges) OrgOrErr() (*Org, error) {
	if e.loadedTypes[1] {
		if e.Org == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: org.Label}
		}
		return e.Org, nil
	}
	return nil, &NotLoadedError{edge: "org"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Membership) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case membership.FieldRole:
			values[i] = new(sql.NullString)
		case membership.FieldID:
			values[i] = new(uuid.UUID)
		case membership.ForeignKeys[0]: // membership_member
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case membership.ForeignKeys[1]: // membership_org
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Membership fields.
func (m *Membership) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case membership.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				m.ID = *value
			}
		case membership.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				m.Role = value.String
			}
		case membership.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field membership_member", values[i])
			} else if value.Valid {
				m.membership_member = new(uuid.UUID)
				*m.membership_member = *value.S.(*uuid.UUID)
			}
		case membership.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field membership_org", values[i])
			} else if value.Valid {
				m.membership_org = new(uuid.UUID)
				*m.membership_org = *value.S.(*uuid.UUID)
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Membership.
// This includes values selected through modifiers, order, etc.
func (m *Membership) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryMember queries the "member" edge of the Membership entity.
func (m *Membership) QueryMember() *UserQuery {
	return NewMembershipClient(m.config).QueryMember(m)
}

// QueryOrg queries the "org" edge of the Membership entity.
func (m *Membership) QueryOrg() *OrgQuery {
	return NewMembershipClient(m.config).QueryOrg(m)
}

// Update returns a builder for updating this Membership.
// Note that you need to call Membership.Unwrap() before calling this method if this Membership
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Membership) Update() *MembershipUpdateOne {
	return NewMembershipClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Membership entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Membership) Unwrap() *Membership {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Membership is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Membership) String() string {
	var builder strings.Builder
	builder.WriteString("Membership(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("role=")
	builder.WriteString(m.Role)
	builder.WriteByte(')')
	return builder.String()
}

// Memberships is a parsable slice of Membership.
type Memberships []*Membership