/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package request

type SendNotificationRequest struct {
	UserID    string            `json:"user_id"`
	TopicID   string            `json:"topic_id"`
	TopicName string            `json:"topic_name"`
	Title     string            `json:"title"`
	Body      string            `json:"body"`
	Payload   map[string]string `json:"payload"`
}
