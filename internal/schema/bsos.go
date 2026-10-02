package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Bsos holds the schema definition for the Bsos entity.
type Bsos struct {
	ent.Schema
}

// Fields of the Bsos.
func (Bsos) Fields() []ent.Field {
	return []ent.Field{
		field.Int64(`user_id`),
		field.Int64(`collection_id`),
		field.String(`bso_id`),
		field.Int(`sortindex`),
		field.Text(`payload `),
		field.Time(`modified`).Default(time.Now),
		field.Time(`expiry`),
	}
}

// Edges of the Bsos.
func (Bsos) Edges() []ent.Edge {
	return []ent.Edge{

	}
}
