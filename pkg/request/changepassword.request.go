/**
 * Author : ngdangkietswe
 * Since  : 5/21/2026
 */

package request

type ChangePasswordRequest struct {
	FirebaseUID string `json:"firebase_uid" validate:"required,firebase_uid"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}
