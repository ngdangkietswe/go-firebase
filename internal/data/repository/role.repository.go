/**
 * Author : ngdangkietswe
 * Since  : 10/25/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/role"
	"go-firebase/internal/data/ent/userrole"

	"github.com/google/uuid"
)

type roleRepo struct {
	cli *ent.Client
}

func (r *roleRepo) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*ent.Role, error) {
	return r.cli.Role.Query().Where(role.HasUserRolesWith(userrole.UserID(userID))).All(ctx)
}

func NewRoleRepository(cli *ent.Client) RoleRepository {
	return &roleRepo{cli: cli}
}
