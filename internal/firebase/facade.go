/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package firebase

import "go-firebase/pkg/request"

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
		CustomToken(claims map[string]interface{}) (string /*customToken*/, error)
	}
)
