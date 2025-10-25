/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package service

import (
	"context"
	"go-firebase/pkg/model"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
)

type (
	UserService interface {
		CreateUser(ctx context.Context, request *request.CreateUserRequest) (*response.CreateUserResponse, error)
		GetUser(ctx context.Context, request *request.GetUserRequest) (*model.User, error)
	}

	AuthService interface {
		Login(ctx context.Context, request *request.LoginRequest) (*response.LoginResponse, error)
		VerifyToken(ctx context.Context, request *request.VerifyTokenRequest) (*response.EmptyResponse, error)
		RefreshToken(ctx context.Context, request *request.RefreshTokenRequest) (*response.RefreshTokenResponse, error)
		CurrentUser(ctx context.Context) (*model.User, error)
	}

	DeviceTokenService interface {
		RegisterDeviceToken(ctx context.Context, request *request.RegisterDeviceRequest) (*response.RegisterDeviceResponse, error)
	}

	NotificationService interface {
		SendNotification(ctx context.Context, request *request.SendNotificationRequest) (*response.SendNotificationResponse, error)
		GetNotifications(ctx context.Context, request *request.ListNotificationRequest) (*response.ListResponse, error)
		MarkNotificationAsRead(ctx context.Context, notificationID string) (*response.EmptyResponse, error)
		MarkAllNotificationsAsRead(ctx context.Context) (*response.EmptyResponse, error)
	}

	NotificationTopicService interface {
		GetNotificationTopics(ctx context.Context, request *request.ListNotificationTopicRequest) (*response.ListResponse, error)
		CreateNotificationTopic(ctx context.Context, request *request.CreateNotificationTopicRequest) (*response.IdResponse, error)
		SubscribeNotificationTopic(ctx context.Context, request *request.SubscribeNotificationTopicRequest) (*response.EmptyResponse, error)
		UnsubscribeNotificationTopic(ctx context.Context, request *request.SubscribeNotificationTopicRequest) (*response.EmptyResponse, error)
	}
)
