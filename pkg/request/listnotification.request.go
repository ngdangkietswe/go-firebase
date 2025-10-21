/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package request

type ListNotificationRequest struct {
	Paginate *PaginateRequest `json:"paginate"`
	IsRead   *bool            `json:"is_read"`
}
