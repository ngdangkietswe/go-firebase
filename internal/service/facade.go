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
		GetUsers(ctx context.Context, request *request.ListUserRequest) (*response.ListResponse, error)
		EnDisableUser(ctx context.Context, request *request.EnDisableUserRequest) (*response.EmptyResponse, error)
		DeleteUser(ctx context.Context, request *request.IDRequest) (*response.EmptyResponse, error)
	}

	AuthService interface {
		Login(ctx context.Context, request *request.LoginRequest) (*response.LoginResponse, error)
		VerifyToken(ctx context.Context, request *request.VerifyTokenRequest) (*response.EmptyResponse, error)
		RefreshToken(ctx context.Context, request *request.RefreshTokenRequest) (*response.RefreshTokenResponse, error)
		RevokeToken(ctx context.Context, request *request.RevokeTokenRequest) (*response.EmptyResponse, error)
		CurrentUser(ctx context.Context) (*model.User, error)
		ForgotPassword(ctx context.Context, request *request.SendPasswordResetMailRequest) (*response.EmptyResponse, error)
		ConfirmResetPassword(ctx context.Context, request *request.ResetPasswordRequest) (*response.EmptyResponse, error)
		AdminChangePassword(ctx context.Context, request *request.ChangePasswordRequest) (*response.EmptyResponse, error)
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

	FileStorageService interface {
		GenerateUploadURL(ctx context.Context, request *request.GetPresignURLRequest) (*response.PresignURLResponse, error)
		GenerateDownloadURL(ctx context.Context, request *request.GetPresignURLRequest) (*response.PresignURLResponse, error)
	}
)
