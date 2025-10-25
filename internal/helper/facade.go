/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package helper

import (
	"context"
	"go-firebase/pkg/model"

	"github.com/google/uuid"
)

type (
	UserHelper interface {
		Preload(ctx context.Context, users []*model.User, preload []string)
	}

	AuthHelper interface {
		BuildPrincipal(claims map[string]interface{}) (*model.Principal, error)
	}

	NotificationTopicHelper interface {
		FirebaseSubscribeToTopic(ctx context.Context, userID uuid.UUID, topicIDs []uuid.UUID) error
		FirebaseUnsubscribeFromTopic(ctx context.Context, userID uuid.UUID, topicIDs []uuid.UUID) error
	}
)
