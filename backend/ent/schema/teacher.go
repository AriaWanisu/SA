package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Teacher holds the schema definition for the Teacher entity.
type Teacher struct {
	ent.Schema
}

// Fields of the Teacher.
func (Teacher) Fields() []ent.Field {
	return []ent.Field{
		field.String("teacher_email").
			NotEmpty().
			Unique(),
		field.String("teacher_name").
			NotEmpty(),
	}
}

// Edges of the Teacher.
func (Teacher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subjects", Subject.Type).
			StorageKey(edge.Column("owner_id")),
		edge.To("courses", Course.Type).
			StorageKey(edge.Column("owner_id")),
	}
}
