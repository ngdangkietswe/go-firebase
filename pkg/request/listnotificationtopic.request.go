/**
 * Author : ngdangkietswe
 * Since  : 10/25/2025
 */

package request

type ListNotificationTopicRequest struct {
	Paginate *PaginateRequest `json:"paginate"`
	Search   string           `json:"search"`
}
