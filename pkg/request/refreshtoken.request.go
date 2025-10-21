/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package request

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
