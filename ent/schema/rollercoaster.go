package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// RollerCoaster holds the schema definition for the RollerCoaster entity.
type RollerCoaster struct {
	ent.Schema
}

// Fields of the RollerCoaster.
func (RollerCoaster) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("roller-coaster"),
		field.String("location").Default("Lisbon, Portugal"),
		field.Float("height").Default(0.00),
		field.Float("speed").Default(0.00),
	}
}

// Edges of the RollerCoaster.
func (RollerCoaster) Edges() []ent.Edge {
	return nil
}
