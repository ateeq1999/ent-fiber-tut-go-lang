package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("userId").Nillable(),
		field.String("title").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type),
	}
}

// Indexes of the User.
func (Post) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
