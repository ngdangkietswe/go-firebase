/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/pkg/request"

	"github.com/google/uuid"
)

type (
	UserRepository interface {
		Save(ctx context.Context, tx *ent.Tx, request *request.CreateUserRequest, firebaseUID string) (*ent.User, error)
		FindByEmail(ctx context.Context, email string) (*ent.User, error)
		FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
		FindByEmailOrID(ctx context.Context, identifier string) (*ent.User, error)
		FindByFirebaseUID(ctx context.Context, firebaseUID string) (*ent.User, error)
		ExistsByID(ctx context.Context, id uuid.UUID) (bool, error)
	}

	DeviceTokenRepository interface {
		Save(ctx context.Context, tx *ent.Tx, request *request.RegisterDeviceRequest, userID uuid.UUID) (*ent.DeviceToken, error)
		FindByUserIDAndDeviceToken(ctx context.Context, userID uuid.UUID, deviceToken string) (*ent.DeviceToken, error)
		FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*ent.DeviceToken, error)
		FindAllByUserIDIn(ctx context.Context, userIDs []uuid.UUID) ([]*ent.DeviceToken, error)
	}

	NotificationRepository interface {
		Save(ctx context.Context, tx *ent.Tx, request *request.SendNotificationRequest) (*ent.Notification, error)
		FindAll(ctx context.Context, request *request.ListNotificationRequest) ([]*ent.Notification, int, error)
		MarkAsRead(ctx context.Context, tx *ent.Tx, notificationID uuid.UUID) error
		MarkAllAsRead(ctx context.Context, tx *ent.Tx) error
		ExistsByID(ctx context.Context, notificationID uuid.UUID) (bool, error)
	}

	NotificationTopicRepository interface {
		Save(ctx context.Context, tx *ent.Tx, topicName string) (*ent.NotificationTopic, error)
		FindAll(ctx context.Context, request *request.ListNotificationTopicRequest) ([]*ent.NotificationTopic, int, error)
		FindByID(ctx context.Context, topicID uuid.UUID) (*ent.NotificationTopic, error)
		FindByName(ctx context.Context, topicName string) (*ent.NotificationTopic, error)
		FindAllByIDIn(ctx context.Context, topicIDs []uuid.UUID) ([]*ent.NotificationTopic, error)
	}

	UserNotificationTopicRepository interface {
		Save(ctx context.Context, tx *ent.Tx, userID uuid.UUID, topicID uuid.UUID) (*ent.UserNotificationTopic, error)
		SaveAll(ctx context.Context, tx *ent.Tx, userID uuid.UUID, topicIDs []uuid.UUID) error
		ExistsByUserIDAndTopicID(ctx context.Context, userID uuid.UUID, topicID uuid.UUID) (bool, error)
		DeleteByUserIDAndTopicID(ctx context.Context, tx *ent.Tx, userID uuid.UUID, topicID uuid.UUID) error
		DeleteByUserIDAndTopicIDIn(ctx context.Context, tx *ent.Tx, userID uuid.UUID, topicIDs []uuid.UUID) error
	}

	RoleRepository interface {
		FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*ent.Role, error)
	}

	PermissionRepository interface {
		FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*ent.Permission, error)
		FindAllByRoleIDIn(ctx context.Context, roleIDs []uuid.UUID) ([]*ent.Permission, error)
	}
)
