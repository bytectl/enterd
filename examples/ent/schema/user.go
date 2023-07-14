package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("user name"),
		field.String("email").Comment("user email"),
		field.Int8("role").Comment("user role"),
		field.Time("created").
			Default(time.Now),
		field.Int("age").
			Range(0, 1000).
			Optional().
			Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
		edge.To("posts", Post.Type),
		edge.To("parent", User.Type).
			Unique(),
		edge.To("cars", Car.Type),
	}
}
