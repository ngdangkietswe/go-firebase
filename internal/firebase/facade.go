/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package firebase

import "go-firebase/internal/request"

type (
	FCMClient interface {
		SendToToken(token string, title string, body string, data map[string]string) error
		SendToTokens(tokens []string, title string, body string, data map[string]string) error
	}

	FAuthClient interface {
		Login(request *request.LoginRequest) (map[string]interface{}, error)
		Signup(request *request.CreateUserRequest) (string /*firebaseUID*/, error)
		VerifyIDToken(idToken string) (string /*firebaseUID*/, error)
	}
)
