/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package request

type GetUserRequest struct {
	Identifier string   `json:"identifier" binding:"required"` // userID or email
	Preload    []string `json:"preload,omitempty"`             // optional preload fields
}
