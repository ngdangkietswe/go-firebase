/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package helper

import (
	"context"
	"go-firebase/internal/model"
)

type (
	UserHelper interface {
		Preload(ctx context.Context, users []*model.User, preload []string)
	}
)
