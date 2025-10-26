/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package permission

import (
	"go-firebase/pkg/constant"
	"go-firebase/pkg/model"
)

const ResourceNotificationTopic = "notification_topic"

func CreateNotificationTopicPerm() *model.Permission {
	return &model.Permission{
		Action:   constant.ActionCreate,
		Resource: ResourceNotificationTopic,
	}
}
