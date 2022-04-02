// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ArifulProtik/arifulprotik-api/ent/auth"
	"github.com/google/uuid"
)

// Auth is the model entity for the Auth schema.
type Auth struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Userid holds the value of the "userid" field.
	Userid uuid.UUID `json:"userid,omitempty"`
	// RfUUID holds the value of the "rf_uuid" field.
	RfUUID uuid.UUID `json:"rf_uuid,omitempty"`
	// IsBlocked holds the value of the "is_blocked" field.
	IsBlocked bool `json:"is_blocked,omitempty"`
	// RfToken holds the value of the "rf_token" field.
	RfToken string `json:"rf_token,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ExpireAt holds the value of the "expire_at" field.
	ExpireAt int64 `json:"expire_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Auth) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case auth.FieldIsBlocked:
			values[i] = new(sql.NullBool)
		case auth.FieldExpireAt:
			values[i] = new(sql.NullInt64)
		case auth.FieldRfToken:
			values[i] = new(sql.NullString)
		case auth.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case auth.FieldID, auth.FieldUserid, auth.FieldRfUUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Auth", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Auth fields.
func (a *Auth) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case auth.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case auth.FieldUserid:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field userid", values[i])
			} else if value != nil {
				a.Userid = *value
			}
		case auth.FieldRfUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field rf_uuid", values[i])
			} else if value != nil {
				a.RfUUID = *value
			}
		case auth.FieldIsBlocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_blocked", values[i])
			} else if value.Valid {
				a.IsBlocked = value.Bool
			}
		case auth.FieldRfToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rf_token", values[i])
			} else if value.Valid {
				a.RfToken = value.String
			}
		case auth.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case auth.FieldExpireAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field expire_at", values[i])
			} else if value.Valid {
				a.ExpireAt = value.Int64
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Auth.
// Note that you need to call Auth.Unwrap() before calling this method if this Auth
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Auth) Update() *AuthUpdateOne {
	return (&AuthClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Auth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Auth) Unwrap() *Auth {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Auth is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Auth) String() string {
	var builder strings.Builder
	builder.WriteString("Auth(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", userid=")
	builder.WriteString(fmt.Sprintf("%v", a.Userid))
	builder.WriteString(", rf_uuid=")
	builder.WriteString(fmt.Sprintf("%v", a.RfUUID))
	builder.WriteString(", is_blocked=")
	builder.WriteString(fmt.Sprintf("%v", a.IsBlocked))
	builder.WriteString(", rf_token=")
	builder.WriteString(a.RfToken)
	builder.WriteString(", created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", expire_at=")
	builder.WriteString(fmt.Sprintf("%v", a.ExpireAt))
	builder.WriteByte(')')
	return builder.String()
}

// Auths is a parsable slice of Auth.
type Auths []*Auth

func (a Auths) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}