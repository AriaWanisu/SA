package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("course_name").
			NotEmpty().
			Unique(),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Teacher.Type).
			Ref("courses").
			Unique(),
		edge.To("course_items", CourseItem.Type).
			StorageKey(edge.Column("course_id")),
	}
}
