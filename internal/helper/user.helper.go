/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package helper

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/repository"
	"go-firebase/internal/mapper"
	"go-firebase/internal/model"
	"go-firebase/internal/util"
	"sync"

	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type userHelper struct {
	logger            *logger.Logger
	deviceTokenRepo   repository.DeviceTokenRepository
	deviceTokenMapper mapper.DeviceTokenMapper
}

const PreloadDeviceTokens = "device_tokens"

func (h *userHelper) Preload(ctx context.Context, users []*model.User, preload []string) {
	var wg sync.WaitGroup

	if lo.Contains(preload, PreloadDeviceTokens) {
		wg.Add(1)
		go h.preloadDeviceTokens(ctx, users, &wg)
	}

	wg.Wait()
}

func (h *userHelper) preloadDeviceTokens(ctx context.Context, users []*model.User, wg *sync.WaitGroup) {
	defer wg.Done()
	defer util.RecoverPanic()

	userIDs := lo.Map(users, func(item *model.User, index int) uuid.UUID {
		return uuid.MustParse(item.UserID)
	})

	deviceTokens, err := h.deviceTokenRepo.FindAllByUserIDIn(ctx, userIDs)
	if err != nil {
		h.logger.Warn("Failed to preload device tokens", zap.Error(err))
		return
	}

	deviceTokensGroup := lo.GroupBy(deviceTokens, func(item *ent.DeviceToken) uuid.UUID {
		return item.UserID
	})

	lo.ForEach(users, func(user *model.User, _ int) {
		if tokens, ok := deviceTokensGroup[uuid.MustParse(user.UserID)]; ok {
			user.Devices = h.deviceTokenMapper.AsList(tokens)
		}
	})
}

func NewUserHelper(
	logger *logger.Logger,
	deviceTokenRepo repository.DeviceTokenRepository,
	deviceTokenMapper mapper.DeviceTokenMapper,
) UserHelper {
	return &userHelper{
		logger:            logger,
		deviceTokenRepo:   deviceTokenRepo,
		deviceTokenMapper: deviceTokenMapper,
	}
}
