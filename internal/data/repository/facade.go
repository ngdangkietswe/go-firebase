/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/request"

	"github.com/google/uuid"
)

type (
	UserRepository interface {
		Save(ctx context.Context, tx *ent.Tx, request *request.CreateUserRequest, firebaseUID string) (*ent.User, error)
		FindByEmail(ctx context.Context, email string) (*ent.User, error)
		FindByID(ctx context.Context, id uuid.UUID) (*ent.User, error)
		FindByEmailOrID(ctx context.Context, identifier string) (*ent.User, error)
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
		FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*ent.Notification, error)
		MarkAsRead(ctx context.Context, tx *ent.Tx, notificationID uuid.UUID) error
		MarkAllAsReadByUserID(ctx context.Context, tx *ent.Tx, userID uuid.UUID) error
		ExistsByID(ctx context.Context, notificationID uuid.UUID) (bool, error)
	}
)
