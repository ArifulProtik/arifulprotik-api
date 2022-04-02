package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("userid", uuid.UUID{}),
		field.UUID("rf_uuid", uuid.UUID{}),
		field.Bool("is_blocked").Default(false),
		field.String("rf_token"),
		field.Time("created_at").Default(time.Now),
		field.Int64("expire_at"),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return nil
}
