/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package response

type LoginResponse struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	LocalID      string `json:"local_id"`
	Error        any    `json:"error,omitempty"`
}
