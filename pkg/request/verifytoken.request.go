/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package request

type VerifyTokenRequest struct {
	IDToken string `json:"id_token" binding:"required"`
}
