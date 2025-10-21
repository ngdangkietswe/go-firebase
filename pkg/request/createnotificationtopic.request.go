/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package request

type CreateNotificationTopicRequest struct {
	TopicName string `json:"topic_name" binding:"required"`
}
