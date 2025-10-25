/**
 * Author : ngdangkietswe
 * Since  : 10/25/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/permission"
	"go-firebase/internal/data/ent/rolepermission"
	"go-firebase/internal/data/ent/userpermission"

	"github.com/google/uuid"
)

type permissionRepo struct {
	cli *ent.Client
}

func (r *permissionRepo) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*ent.Permission, error) {
	return r.cli.Permission.Query().
		Where(permission.HasUserPermissionsWith(userpermission.UserID(userID))).
		All(ctx)
}

func (r *permissionRepo) FindAllByRoleIDIn(ctx context.Context, roleIDs []uuid.UUID) ([]*ent.Permission, error) {
	return r.cli.Permission.Query().
		Where(permission.HasRolePermissionsWith(rolepermission.RoleIDIn(roleIDs...))).
		All(ctx)
}

func NewPermissionRepository(cli *ent.Client) PermissionRepository {
	return &permissionRepo{cli: cli}
}
