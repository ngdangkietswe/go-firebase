/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package request

type SubscribeNotificationTopicRequest struct {
	TopicIDs []string `json:"topic_ids" binding:"required,min=1,dive,required"`
}
