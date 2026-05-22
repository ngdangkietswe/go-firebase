/**
 * Author : ngdangkietswe
 * Since  : 5/21/2026
 */

package request

type ResetPasswordRequest struct {
	OobCode     string `json:"oob_code" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
