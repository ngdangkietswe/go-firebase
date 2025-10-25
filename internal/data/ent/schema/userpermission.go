package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	"github.com/ngdangkietswe/swe-go-common-shared/util"
)

// UserPermission holds the schema definition for the UserPermission entity.
type UserPermission struct {
	ent.Schema
}

// Fields of the UserPermission.
func (UserPermission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("permission_id", uuid.UUID{}),
	}
}

// Edges of the UserPermission.
func (UserPermission) Edges() []ent.Edge {
	return []ent.Edge{
		util.One2ManyInverseRequired("user", User.Type, "user_permissions", "user_id"),
		util.One2ManyInverseRequired("permission", Permission.Type, "user_permissions", "permission_id"),
	}
}

func (UserPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "permission_id"),
	}
}

// Annotations of the UserPermission.
func (UserPermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_permissions"},
	}
}
