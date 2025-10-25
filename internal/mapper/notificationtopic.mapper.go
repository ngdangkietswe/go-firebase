/**
 * Author : ngdangkietswe
 * Since  : 10/25/2025
 */

package mapper

import (
	"go-firebase/internal/data/ent"
	"go-firebase/pkg/model"

	"github.com/samber/lo"
)

type notificationTopicMapper struct{}

func (m *notificationTopicMapper) AsMono(topic *ent.NotificationTopic) *model.NotificationTopic {
	builder := &model.NotificationTopic{
		ID:   topic.ID.String(),
		Name: topic.Name,
	}

	if topic.Description != "" {
		builder.Description = topic.Description
	}

	return builder
}

func (m *notificationTopicMapper) AsList(topics []*ent.NotificationTopic) []*model.NotificationTopic {
	if len(topics) == 0 {
		return []*model.NotificationTopic{}
	}

	builders := make([]*model.NotificationTopic, 0, len(topics))

	lo.ForEach(topics, func(topic *ent.NotificationTopic, _ int) {
		builders = append(builders, m.AsMono(topic))
	})

	return builders
}

func NewNotificationTopicMapper() NotificationTopicMapper {
	return &notificationTopicMapper{}
}
