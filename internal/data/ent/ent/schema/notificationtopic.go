package schema

import "entgo.io/ent"

// NotificationTopic holds the schema definition for the NotificationTopic entity.
type NotificationTopic struct {
	ent.Schema
}

// Fields of the NotificationTopic.
func (NotificationTopic) Fields() []ent.Field {
	return nil
}

// Edges of the NotificationTopic.
func (NotificationTopic) Edges() []ent.Edge {
	return nil
}
