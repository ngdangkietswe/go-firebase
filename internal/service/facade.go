/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package service

import (
	"context"
	"go-firebase/internal/model"
	"go-firebase/internal/request"
	"go-firebase/internal/response"
)

type (
	UserService interface {
		CreateUser(ctx context.Context, request *request.CreateUserRequest) (*response.CreateUserResponse, error)
		GetUser(ctx context.Context, request *request.GetUserRequest) (*model.User, error)
	}

	AuthService interface {
		Login(ctx context.Context, request *request.LoginRequest) (*response.LoginResponse, error)
		VerifyToken(ctx context.Context, request *request.VerifyTokenRequest) (*response.EmptyResponse, error)
	}

	DeviceTokenService interface {
		RegisterDeviceToken(ctx context.Context, request *request.RegisterDeviceRequest) (*response.RegisterDeviceResponse, error)
	}

	NotificationService interface {
		SendNotification(ctx context.Context, request *request.SendNotificationRequest) (*response.SendNotificationResponse, error)
		GetNotifications(ctx context.Context, userID string) ([]*model.Notification, error)
		MarkNotificationAsRead(ctx context.Context, notificationID string) (*response.EmptyResponse, error)
		MarkAllNotificationsAsRead(ctx context.Context, userID string) (*response.EmptyResponse, error)
	}
)
