// Code generated by entc, DO NOT EDIT.

package auth

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ArifulProtik/arifulprotik-api/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Userid applies equality check predicate on the "userid" field. It's identical to UseridEQ.
func Userid(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserid), v))
	})
}

// RfUUID applies equality check predicate on the "rf_uuid" field. It's identical to RfUUIDEQ.
func RfUUID(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRfUUID), v))
	})
}

// IsBlocked applies equality check predicate on the "is_blocked" field. It's identical to IsBlockedEQ.
func IsBlocked(v bool) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsBlocked), v))
	})
}

// RfToken applies equality check predicate on the "rf_token" field. It's identical to RfTokenEQ.
func RfToken(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRfToken), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// ExpireAt applies equality check predicate on the "expire_at" field. It's identical to ExpireAtEQ.
func ExpireAt(v int64) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpireAt), v))
	})
}

// UseridEQ applies the EQ predicate on the "userid" field.
func UseridEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserid), v))
	})
}

// UseridNEQ applies the NEQ predicate on the "userid" field.
func UseridNEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserid), v))
	})
}

// UseridIn applies the In predicate on the "userid" field.
func UseridIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserid), v...))
	})
}

// UseridNotIn applies the NotIn predicate on the "userid" field.
func UseridNotIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserid), v...))
	})
}

// UseridGT applies the GT predicate on the "userid" field.
func UseridGT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserid), v))
	})
}

// UseridGTE applies the GTE predicate on the "userid" field.
func UseridGTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserid), v))
	})
}

// UseridLT applies the LT predicate on the "userid" field.
func UseridLT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserid), v))
	})
}

// UseridLTE applies the LTE predicate on the "userid" field.
func UseridLTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserid), v))
	})
}

// RfUUIDEQ applies the EQ predicate on the "rf_uuid" field.
func RfUUIDEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRfUUID), v))
	})
}

// RfUUIDNEQ applies the NEQ predicate on the "rf_uuid" field.
func RfUUIDNEQ(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRfUUID), v))
	})
}

// RfUUIDIn applies the In predicate on the "rf_uuid" field.
func RfUUIDIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRfUUID), v...))
	})
}

// RfUUIDNotIn applies the NotIn predicate on the "rf_uuid" field.
func RfUUIDNotIn(vs ...uuid.UUID) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRfUUID), v...))
	})
}

// RfUUIDGT applies the GT predicate on the "rf_uuid" field.
func RfUUIDGT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRfUUID), v))
	})
}

// RfUUIDGTE applies the GTE predicate on the "rf_uuid" field.
func RfUUIDGTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRfUUID), v))
	})
}

// RfUUIDLT applies the LT predicate on the "rf_uuid" field.
func RfUUIDLT(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRfUUID), v))
	})
}

// RfUUIDLTE applies the LTE predicate on the "rf_uuid" field.
func RfUUIDLTE(v uuid.UUID) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRfUUID), v))
	})
}

// IsBlockedEQ applies the EQ predicate on the "is_blocked" field.
func IsBlockedEQ(v bool) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsBlocked), v))
	})
}

// IsBlockedNEQ applies the NEQ predicate on the "is_blocked" field.
func IsBlockedNEQ(v bool) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsBlocked), v))
	})
}

// RfTokenEQ applies the EQ predicate on the "rf_token" field.
func RfTokenEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRfToken), v))
	})
}

// RfTokenNEQ applies the NEQ predicate on the "rf_token" field.
func RfTokenNEQ(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRfToken), v))
	})
}

// RfTokenIn applies the In predicate on the "rf_token" field.
func RfTokenIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRfToken), v...))
	})
}

// RfTokenNotIn applies the NotIn predicate on the "rf_token" field.
func RfTokenNotIn(vs ...string) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRfToken), v...))
	})
}

// RfTokenGT applies the GT predicate on the "rf_token" field.
func RfTokenGT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRfToken), v))
	})
}

// RfTokenGTE applies the GTE predicate on the "rf_token" field.
func RfTokenGTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRfToken), v))
	})
}

// RfTokenLT applies the LT predicate on the "rf_token" field.
func RfTokenLT(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRfToken), v))
	})
}

// RfTokenLTE applies the LTE predicate on the "rf_token" field.
func RfTokenLTE(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRfToken), v))
	})
}

// RfTokenContains applies the Contains predicate on the "rf_token" field.
func RfTokenContains(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRfToken), v))
	})
}

// RfTokenHasPrefix applies the HasPrefix predicate on the "rf_token" field.
func RfTokenHasPrefix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRfToken), v))
	})
}

// RfTokenHasSuffix applies the HasSuffix predicate on the "rf_token" field.
func RfTokenHasSuffix(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRfToken), v))
	})
}

// RfTokenEqualFold applies the EqualFold predicate on the "rf_token" field.
func RfTokenEqualFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRfToken), v))
	})
}

// RfTokenContainsFold applies the ContainsFold predicate on the "rf_token" field.
func RfTokenContainsFold(v string) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRfToken), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// ExpireAtEQ applies the EQ predicate on the "expire_at" field.
func ExpireAtEQ(v int64) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpireAt), v))
	})
}

// ExpireAtNEQ applies the NEQ predicate on the "expire_at" field.
func ExpireAtNEQ(v int64) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpireAt), v))
	})
}

// ExpireAtIn applies the In predicate on the "expire_at" field.
func ExpireAtIn(vs ...int64) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExpireAt), v...))
	})
}

// ExpireAtNotIn applies the NotIn predicate on the "expire_at" field.
func ExpireAtNotIn(vs ...int64) predicate.Auth {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Auth(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExpireAt), v...))
	})
}

// ExpireAtGT applies the GT predicate on the "expire_at" field.
func ExpireAtGT(v int64) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpireAt), v))
	})
}

// ExpireAtGTE applies the GTE predicate on the "expire_at" field.
func ExpireAtGTE(v int64) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpireAt), v))
	})
}

// ExpireAtLT applies the LT predicate on the "expire_at" field.
func ExpireAtLT(v int64) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpireAt), v))
	})
}

// ExpireAtLTE applies the LTE predicate on the "expire_at" field.
func ExpireAtLTE(v int64) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpireAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Auth) predicate.Auth {
	return predicate.Auth(func(s *sql.Selector) {
		p(s.Not())
	})
}