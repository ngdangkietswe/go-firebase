/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package response

type SendNotificationResponse struct {
	NotificationID string   `json:"notification_id" binding:"required"`
	UserID         string   `json:"user_id" binding:"required"`
	DeviceIDs      []string `json:"device_ids" binding:"required"`
}
