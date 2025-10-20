/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package request

type ListUserRequest struct {
	Paginate *PaginateRequest `json:"paginate,omitempty"`
	Search   string           `json:"search,omitempty"`  // optional search keyword
	Preload  []string         `json:"preload,omitempty"` // optional preload fields
}
