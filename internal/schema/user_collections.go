package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserCollections holds the schema definition for the UserCollections entity.
type UserCollections struct {
	ent.Schema
}

// Fields of the UserCollections.
func (UserCollections) Fields() []ent.Field {
	return []ent.Field{
		field.Int64(`user_id`),
		field.Int64(`collection_id`),
		field.Time(`modified`).Default(time.Now),
		field.Int64(`count`),
		field.Int64(`total_bytes`),
	}
}

// Edges of the UserCollections.
func (UserCollections) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("collections", Collections.Type).Unique(),
	}
}
func (c UserCollections) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields(`user_id`, `collection_id`).Unique(),
	}
}
