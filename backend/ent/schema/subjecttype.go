package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// SubjectType holds the schema definition for the SubjectType entity.
type SubjectType struct {
	ent.Schema
}

// Fields of the SubjectType.
func (SubjectType) Fields() []ent.Field {
	return []ent.Field{
		field.String("type_name").
			NotEmpty().
			Unique(),
	}
}

// Edges of the SubjectType.
func (SubjectType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("course_items", CourseItem.Type).
			StorageKey(edge.Column("type_id")),
	}
}
