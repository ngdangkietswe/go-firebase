/**
 * Author : ngdangkietswe
 * Since  : 5/21/2026
 */

package request

type SendPasswordResetMailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
