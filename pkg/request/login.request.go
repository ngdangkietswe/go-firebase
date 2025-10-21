/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package request

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
