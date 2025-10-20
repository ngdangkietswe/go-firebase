/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package repository

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/ent/devicetoken"
	"go-firebase/internal/request"

	"github.com/google/uuid"
)

type deviceTokenRepo struct {
	cli *ent.Client
}

func (r *deviceTokenRepo) Save(ctx context.Context, tx *ent.Tx, request *request.RegisterDeviceRequest, userID uuid.UUID) (*ent.DeviceToken, error) {
	builder := tx.DeviceToken.Create().
		SetUserID(userID).
		SetToken(request.DeviceToken).
		SetPlatform(request.Platform)
	return builder.Save(ctx)
}

func (r *deviceTokenRepo) FindByUserIDAndDeviceToken(ctx context.Context, userID uuid.UUID, deviceToken string) (*ent.DeviceToken, error) {
	return r.cli.DeviceToken.Query().
		Where(
			devicetoken.UserID(userID),
			devicetoken.Token(deviceToken),
		).
		Only(ctx)
}

func (r *deviceTokenRepo) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*ent.DeviceToken, error) {
	return r.cli.DeviceToken.Query().Where(devicetoken.UserID(userID)).All(ctx)
}

func (r *deviceTokenRepo) FindAllByUserIDIn(ctx context.Context, userIDs []uuid.UUID) ([]*ent.DeviceToken, error) {
	return r.cli.DeviceToken.Query().Where(devicetoken.UserIDIn(userIDs...)).All(ctx)
}

func NewDeviceTokenRepository(cli *ent.Client) DeviceTokenRepository {
	return &deviceTokenRepo{cli: cli}
}
