package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Collections holds the schema definition for the Collections entity.
type Collections struct {
	ent.Schema
}

// Fields of the Collections.
func (Collections) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64(`id`).Comment(`PK for collection`),
		field.Time(`updated_at`).Comment(`last modified timestamp`).Default(time.Now),
		field.String(`name`).MaxRuneLen(32).Unique().Comment(`unique name for collection`),
	}
}

// Edges of the Collections.
func (Collections) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("collections", Collections.Type),
	}
}

func (Collections) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields(`name`).Unique(),
	}
}
