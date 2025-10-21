/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package mapper

import (
	"go-firebase/internal/data/ent"
	"go-firebase/pkg/model"

	"github.com/samber/lo"
)

type userMapper struct{}

func (m *userMapper) AsMono(user *ent.User) *model.User {
	builder := model.User{
		UserID: user.ID.String(),
		Email:  user.Email,
	}

	if user.FirstName != "" {
		builder.FirstName = user.FirstName
	}

	if user.LastName != "" {
		builder.LastName = user.LastName
	}

	return &builder
}

func (m *userMapper) AsList(users []*ent.User) []*model.User {
	if len(users) == 0 {
		return []*model.User{}
	}

	builders := make([]*model.User, 0, len(users))

	lo.ForEach(users, func(user *ent.User, _ int) {
		builders = append(builders, m.AsMono(user))
	})

	return builders
}

func NewUserMapper() UserMapper {
	return &userMapper{}
}
