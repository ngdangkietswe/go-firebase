/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package mapper

import (
	"go-firebase/internal/data/ent"
	"go-firebase/internal/model"

	"github.com/samber/lo"
)

type deviceTokenMapper struct {
}

func (m *deviceTokenMapper) AsMono(deviceToken *ent.DeviceToken) *model.Device {
	builder := &model.Device{
		DeviceID: deviceToken.ID.String(),
		UserID:   deviceToken.UserID.String(),
		Token:    deviceToken.Token,
		Platform: deviceToken.Platform,
	}
	return builder
}

func (m *deviceTokenMapper) AsList(deviceTokens []*ent.DeviceToken) []*model.Device {
	if len(deviceTokens) == 0 {
		return []*model.Device{}
	}

	builders := make([]*model.Device, 0, len(deviceTokens))

	lo.ForEach(deviceTokens, func(deviceToken *ent.DeviceToken, _ int) {
		builders = append(builders, m.AsMono(deviceToken))
	})

	return builders
}

func NewDeviceTokenMapper() DeviceTokenMapper {
	return &deviceTokenMapper{}
}
