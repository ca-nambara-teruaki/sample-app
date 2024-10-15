package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Comment("ID"),
		field.String("title").NotEmpty().Comment("タイトル"),
		field.String("description").Optional().Comment("説明"),
		field.Int("created_by").Comment("作成者ID"),
		field.Bool("is_deleted").Default(false).Comment("削除フラグ"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
