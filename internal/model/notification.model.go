/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package model

type Notification struct {
	NotificationID string `json:"notification_id"`
	Title          string `json:"title"`
	Body           string `json:"body"`
	UserID         string `json:"user_id"`
	SentAt         string `json:"sent_at"`
	IsRead         bool   `json:"is_read"`
}
