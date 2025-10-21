/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package request

type PaginateRequest struct {
	Page     int    `json:"page" binding:"required,min=1"`
	PageSize int    `json:"page_size" binding:"required,min=1,max=100"`
	Sort     string `json:"sort,omitempty"`  // optional sort field
	Order    string `json:"order,omitempty"` // optional sort order: asc or desc
}
