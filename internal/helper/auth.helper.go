/**
 * Author : ngdangkietswe
 * Since  : 10/25/2025
 */

package helper

import (
	"context"
	"errors"
	"fmt"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/repository"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/model"
	"time"

	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/cache"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type authHelper struct {
	logger         *logger.Logger
	cache          *cache.RedisCache
	roleRepo       repository.RoleRepository
	permissionRepo repository.PermissionRepository
}

func (h *authHelper) BuildPrincipal(claims map[string]interface{}) (*model.Principal, error) {
	var principal model.Principal

	if firebaseUID, ok := claims["firebase_uid"].(string); ok {
		principal.FirebaseUID = firebaseUID
	} else {
		return nil, errors.New("firebase UID not found in claims")
	}

	if systemUID, ok := claims["system_uid"].(string); ok {
		principal.SystemUID = systemUID
	} else {
		return nil, errors.New("system UID not found in claims")
	}

	if err := h.cacheUserRoles(&principal); err != nil {
		return nil, err
	}

	if err := h.cacheUserPermissions(&principal); err != nil {
		return nil, err
	}

	return &principal, nil
}

func (h *authHelper) cacheUserRoles(principal *model.Principal) error {
	var roles []*model.Role
	userRoleCacheKey := fmt.Sprintf("%s%s", constant.UserRoleCacheKeyPrefix, principal.SystemUID)

	if err := h.cache.Get(userRoleCacheKey, &roles); err != nil {
		h.logger.Warn("Failed to get user roles from cache", zap.Error(err))
		eRoles, err := h.roleRepo.FindAllByUserID(context.Background(), uuid.MustParse(principal.SystemUID))
		if err != nil {
			h.logger.Error("Failed to get user roles from db", zap.Error(err))
			return err
		}

		if len(eRoles) > 0 {
			roles = lo.Map(eRoles, func(r *ent.Role, _ int) *model.Role {
				return &model.Role{
					ID:   r.ID.String(),
					Name: r.Name,
				}
			})

			if err := h.cache.Set(userRoleCacheKey, roles, time.Duration(30)*time.Minute); err != nil {
				h.logger.Warn("Failed to set user roles to cache", zap.Error(err))
			}
		}
	}

	if len(roles) > 0 {
		principal.Roles = roles
	}

	return nil
}

func (h *authHelper) cacheUserPermissions(principal *model.Principal) error {
	var permissions []*model.Permission
	userPermissionCacheKey := fmt.Sprintf("%s%s", constant.UserPermissionCacheKeyPrefix, principal.SystemUID)

	if err := h.cache.Get(userPermissionCacheKey, &permissions); err != nil {
		permissionMap := make(map[uuid.UUID]*ent.Permission)

		eUPermissions, err := h.permissionRepo.FindAllByUserID(context.Background(), uuid.MustParse(principal.SystemUID))
		if err != nil {
			h.logger.Error("Failed to get user permissions from db", zap.Error(err))
			return err
		}

		if len(eUPermissions) > 0 {
			for _, p := range eUPermissions {
				permissionMap[p.ID] = p
			}
		}

		if len(principal.Roles) > 0 {
			roleIDs := lo.Map(principal.Roles, func(r *model.Role, _ int) uuid.UUID {
				return uuid.MustParse(r.ID)
			})

			eRPermissions, err := h.permissionRepo.FindAllByRoleIDIn(context.Background(), roleIDs)
			if err != nil {
				h.logger.Error("Failed to get role permissions from db", zap.Error(err))
				return err
			}

			if len(eRPermissions) > 0 {
				for _, p := range eRPermissions {
					if _, exists := permissionMap[p.ID]; !exists {
						permissionMap[p.ID] = p
					}
				}
			}
		}

		if len(permissionMap) > 0 {
			for _, p := range permissionMap {
				permissions = append(permissions, &model.Permission{
					Action:   p.Action,
					Resource: p.Resource,
				})
			}
			if err := h.cache.Set(userPermissionCacheKey, permissions, time.Duration(30)*time.Minute); err != nil {
				h.logger.Warn("Failed to set user permissions to cache", zap.Error(err))
			}
		}
	}

	if len(permissions) > 0 {
		principal.Permissions = permissions
	}

	return nil
}

func NewAuthHelper(
	logger *logger.Logger,
	cache *cache.RedisCache,
	roleRepo repository.RoleRepository,
	permissionRepo repository.PermissionRepository,
) AuthHelper {
	return &authHelper{
		logger:         logger,
		cache:          cache,
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
	}
}
