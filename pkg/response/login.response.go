/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package response

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	FirebaseUID  string `json:"firebase_uid"`
	UserID       string `json:"user_id"`
	Error        any    `json:"error,omitempty"`
}
