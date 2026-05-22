package schema

import (
	"go-firebase/internal/data/ent/mixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/ngdangkietswe/swe-go-common-shared/util"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email").Unique(),
		field.String("first_name").Optional().Nillable(),
		field.String("last_name").Optional().Nillable(),
		field.String("display_name").Optional().Nillable(),
		field.String("firebase_uid").Unique(),
		field.Int32("status").Default(1), // e.g., 1: active, 2: inactive
		field.Time("last_login_at").Optional().Nillable(),
		field.String("last_login_ip").Optional().Nillable(),
		field.String("last_login_user_agent").Optional().Nillable(),
		field.Int32("failed_login_attempts").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		util.One2Many("device_tokens", DeviceToken.Type),
		util.One2Many("notifications", Notification.Type),
		util.One2Many("user_notification_topics", UserNotificationTopic.Type),
		util.One2Many("user_roles", UserRole.Type),
		util.One2Many("user_permissions", UserPermission.Type),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}
