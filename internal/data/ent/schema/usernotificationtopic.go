package schema

import (
	"time"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"entgo.io/ent"

	"github.com/ngdangkietswe/swe-go-common-shared/util"
)

// UserNotificationTopic holds the schema definition for the UserNotificationTopic entity.
type UserNotificationTopic struct {
	ent.Schema
}

// Fields of the UserNotificationTopic.
func (UserNotificationTopic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("notification_topic_id", uuid.UUID{}),
		field.Time("subscribed_at").Default(time.Now),
	}
}

// Edges of the UserNotificationTopic.
func (UserNotificationTopic) Edges() []ent.Edge {
	return []ent.Edge{
		util.One2ManyInverseRequired("user", User.Type, "user_notification_topics", "user_id"),
		util.One2ManyInverseRequired("notification_topic", NotificationTopic.Type, "user_notification_topics", "notification_topic_id"),
	}
}

// Annotations of the UserNotificationTopic.
func (UserNotificationTopic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_notification_topic"},
	}
}
