package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// CourseItem holds the schema definition for the CourseItem entity.
type CourseItem struct {
	ent.Schema
}

// Fields of the CourseItem.
func (CourseItem) Fields() []ent.Field {
	return nil
}

// Edges of the CourseItem.
func (CourseItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("courses", Course.Type).
			Ref("course_items").
			Unique(),
		edge.From("subjects", Subject.Type).
			Ref("course_items").
			Unique(),
		edge.From("types", SubjectType.Type).
			Ref("course_items").
			Unique(),
	}
}
