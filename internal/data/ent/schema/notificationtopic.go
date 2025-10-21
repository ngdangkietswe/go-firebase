package schema

import (
	"go-firebase/internal/data/ent/mixin"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"entgo.io/ent"

	"github.com/ngdangkietswe/swe-go-common-shared/util"
)

// NotificationTopic holds the schema definition for the NotificationTopic entity.
type NotificationTopic struct {
	ent.Schema
}

// Fields of the NotificationTopic.
func (NotificationTopic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").Unique(),
		field.String("description").Optional(),
	}
}

// Edges of the NotificationTopic.
func (NotificationTopic) Edges() []ent.Edge {
	return []ent.Edge{
		util.One2Many("user_notification_topics", UserNotificationTopic.Type),
		util.One2Many("notifications", Notification.Type),
	}
}

// Mixin of the NotificationTopic.
func (NotificationTopic) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Annotations of the NotificationTopic.
func (NotificationTopic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notification_topic"},
	}
}
