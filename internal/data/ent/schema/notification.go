package schema

import (
	"go-firebase/internal/data/ent/mixin"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/ngdangkietswe/swe-go-common-shared/util"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.UUID("notification_topic_id", uuid.UUID{}).Optional(),
		field.String("title"),
		field.String("body"),
		field.JSON("data", map[string]string{}).Optional().Comment("Additional data payload"),
		field.Time("sent_at").Default(time.Now),
		field.Bool("is_read").Default(false),
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		util.One2ManyInverse("user", User.Type, "notifications", "user_id"),
		util.One2ManyInverse("notification_topic", NotificationTopic.Type, "notifications", "notification_topic_id"),
	}
}

// Mixin of the Notification.
func (Notification) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Annotations of the Notification.
func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notification"},
	}
}
