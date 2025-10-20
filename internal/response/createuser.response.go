/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package response

type CreateUserResponse struct {
	UserID      string `json:"user_id"`
	FirebaseUID string `json:"firebase_uid"`
}
