/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package response

type RefreshTokenResponse struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
	LocalID      string `json:"local_id"`
}
