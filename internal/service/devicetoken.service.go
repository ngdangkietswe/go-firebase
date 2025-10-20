/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package service

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/repository"
	"go-firebase/internal/request"
	"go-firebase/internal/response"

	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/zap"
)

type deviceTokenSvc struct {
	cli             *ent.Client
	logger          *logger.Logger
	deviceTokenRepo repository.DeviceTokenRepository
	userRepo        repository.UserRepository
}

func (s *deviceTokenSvc) RegisterDeviceToken(ctx context.Context, request *request.RegisterDeviceRequest) (*response.RegisterDeviceResponse, error) {
	user, err := s.userRepo.FindByEmailOrID(ctx, request.Identifier)
	if err != nil {
		s.logger.Error("Failed to find user by identifier", zap.String("identifier", request.Identifier), zap.Error(err))
		return nil, err
	}

	deviceToken, err := repository.WithTxResult(ctx, s.cli, s.logger, func(tx *ent.Tx) (*ent.DeviceToken, error) {
		return s.deviceTokenRepo.Save(ctx, tx, request, user.ID)
	})

	if err != nil {
		s.logger.Error("Failed to save device token", zap.String("user_id", user.ID.String()), zap.String("device_token", request.DeviceToken), zap.Error(err))
	}

	return &response.RegisterDeviceResponse{
		UserID:   user.ID.String(),
		DeviceID: deviceToken.ID.String(),
	}, nil
}

func NewDeviceTokenService(
	cli *ent.Client,
	logger *logger.Logger,
	deviceTokenRepo repository.DeviceTokenRepository,
	userRepo repository.UserRepository,
) DeviceTokenService {
	return &deviceTokenSvc{
		cli:             cli,
		logger:          logger,
		deviceTokenRepo: deviceTokenRepo,
		userRepo:        userRepo,
	}
}
