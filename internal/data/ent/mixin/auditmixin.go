/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package mixin

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type AuditMixin struct {
	mixin.Schema
}

func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional().Nillable(),
		field.Time("deleted_at").Optional().Nillable(),
		field.UUID("created_by", uuid.UUID{}).Optional().Nillable(),
		field.UUID("updated_by", uuid.UUID{}).Optional().Nillable(),
		field.UUID("deleted_by", uuid.UUID{}).Optional().Nillable(),
		field.Bool("deleted").Default(false),
	}
}
