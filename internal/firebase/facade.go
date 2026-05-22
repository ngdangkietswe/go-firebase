/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package firebase

import (
	"go-firebase/pkg/request"
	"time"
)

type (
	FCMClient interface {
		SendToToken(token string, title string, body string, data map[string]string) error
		SendToTokens(tokens []string, title string, body string, data map[string]string) error
		SendToTopic(topic string, title string, body string, data map[string]string) error
		SubscribeToTopic(tokens []string, topic string) error
		UnsubscribeFromTopic(tokens []string, topic string) error
	}

	FAuthClient interface {
		LoginWithPassword(request *request.LoginRequest) (map[string]interface{}, error)
		LoginWithCustomToken(customToken string) (map[string]interface{}, error)
		Signup(request *request.CreateUserRequest) (string /*firebaseUID*/, error)
		VerifyIDToken(idToken string) (map[string]interface{}, error)
		RefreshToken(request *request.RefreshTokenRequest) (map[string]interface{}, error)
		RevokeToken(request *request.RevokeTokenRequest) error
		CustomToken(claims map[string]interface{}) (string /*customToken*/, error)
		EnDisableAccount(firebaseUID string, disabled bool) error
		DeleteAccount(firebaseUID string) error
		SendPasswordResetEmail(request *request.SendPasswordResetMailRequest) error
		ResetPassword(request *request.ResetPasswordRequest) error
		ChangePassword(request *request.ChangePasswordRequest) error
	}

	FStorageClient interface {
		UploadFile(bucketName string, objectName string, contentType string, expireDuration time.Duration) (string /*presignUrl*/, error)
		DownloadFile(bucketName string, objectName string, expireDuration time.Duration) (string /*presignUrl*/, error)
	}
)
