/**
 * Author : ngdangkietswe
 * Since  : 11/1/2025
 */

package request

type EnDisableUserRequest struct {
	UserID  string `json:"user_id"`
	Disable bool   `json:"disable"`
}
