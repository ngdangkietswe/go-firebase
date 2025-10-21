/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package response

type SendNotificationResponse struct {
	NotificationID string   `json:"notification_id" binding:"required"`
	UserID         string   `json:"user_id"`
	TopicID        string   `json:"topic_id"`
	DeviceIDs      []string `json:"device_ids"`
}
