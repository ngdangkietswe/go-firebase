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

// DeviceToken holds the schema definition for the DeviceToken entity.
type DeviceToken struct {
	ent.Schema
}

// Fields of the DeviceToken.
func (DeviceToken) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.String("token").Unique(),
		field.String("platform").Default("android").Comment("android|ios|web"),
		field.Time("last_seen").Default(time.Now),
		field.Bool("is_active").Default(true),
	}
}

// Edges of the DeviceToken.
func (DeviceToken) Edges() []ent.Edge {
	return []ent.Edge{
		util.One2ManyInverseRequired("user", User.Type, "device_tokens", "user_id"),
	}
}

// Mixin of the DeviceToken.
func (DeviceToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}

// Annotations of the DeviceToken.
func (DeviceToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "device_token"},
	}
}
