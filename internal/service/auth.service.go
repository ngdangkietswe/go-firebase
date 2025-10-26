/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package service

import (
	"context"
	"errors"
	"go-firebase/internal/data/repository"
	"go-firebase/internal/firebase"
	"go-firebase/internal/helper"
	"go-firebase/internal/mapper"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/model"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/zap"
)

type authSvc struct {
	logger     *logger.Logger
	fAuthCli   firebase.FAuthClient
	userRepo   repository.UserRepository
	userMapper mapper.UserMapper
	userHelper helper.UserHelper
}

func (s *authSvc) Login(ctx context.Context, request *request.LoginRequest) (*response.LoginResponse, error) {
	fAuthLoginWithPasswordResp, err := s.fAuthCli.LoginWithPassword(request)
	if err != nil {
		return nil, err
	}

	firebaseUID, ok := fAuthLoginWithPasswordResp["localId"].(string)
	if !ok || firebaseUID == "" {
		return nil, errors.New("firebase UID not found in login response")
	}

	user, err := s.userRepo.FindByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		s.logger.Error("Failed to find user by firebase UID", zap.String("firebase_uid", firebaseUID), zap.Error(err))
		return nil, err
	}

	claims := map[string]interface{}{
		"firebase_uid": firebaseUID,
		"system_uid":   user.ID.String(),
	}

	customToken, err := s.fAuthCli.CustomToken(claims)
	if err != nil {
		s.logger.Error("Failed to create custom token", zap.String("firebase_uid", firebaseUID), zap.Error(err))
		return nil, err
	}

	fAuthLoginWithCustomTokenResp, err := s.fAuthCli.LoginWithCustomToken(customToken)
	if err != nil {
		s.logger.Error("Failed to login with custom token", zap.String("firebase_uid", firebaseUID), zap.Error(err))
		return nil, err
	}

	idToken, ok := fAuthLoginWithCustomTokenResp["idToken"].(string)
	if !ok || idToken == "" {
		return nil, errors.New("ID token not found in custom token login response")
	}

	resp := &response.LoginResponse{
		AccessToken:  idToken,
		ExpiresIn:    fAuthLoginWithCustomTokenResp["expiresIn"].(string),
		RefreshToken: fAuthLoginWithCustomTokenResp["refreshToken"].(string),
		FirebaseUID:  firebaseUID,
		UserID:       user.ID.String(),
	}

	return resp, nil
}

func (s *authSvc) VerifyToken(ctx context.Context, request *request.VerifyTokenRequest) (*response.EmptyResponse, error) {
	_, err := s.fAuthCli.VerifyIDToken(request.IDToken)
	if err != nil {
		return nil, err
	}
	return &response.EmptyResponse{}, nil
}

func (s *authSvc) RefreshToken(ctx context.Context, request *request.RefreshTokenRequest) (*response.RefreshTokenResponse, error) {
	fAuthRefreshResp, err := s.fAuthCli.RefreshToken(request)
	if err != nil {
		return nil, err
	}

	resp := &response.RefreshTokenResponse{
		IDToken:      fAuthRefreshResp["id_token"].(string),
		RefreshToken: fAuthRefreshResp["refresh_token"].(string),
		ExpiresIn:    fAuthRefreshResp["expires_in"].(string),
		LocalID:      fAuthRefreshResp["user_id"].(string),
	}

	return resp, nil
}

func (s *authSvc) CurrentUser(ctx context.Context) (*model.User, error) {
	principal := util.GetPrincipal(ctx)

	user, err := s.userRepo.FindByID(ctx, uuid.MustParse(principal.SystemUID))
	if err != nil {
		s.logger.Error("Failed to get current user", zap.Error(err))
		return nil, err
	}

	mUser := s.userMapper.AsMono(user)

	preload := []string{constant.PreloadDeviceTokens}
	s.userHelper.Preload(ctx, []*model.User{mUser}, preload)

	return mUser, nil
}

func NewAuthService(
	logger *logger.Logger,
	fAuthCli firebase.FAuthClient,
	userRepo repository.UserRepository,
	userMapper mapper.UserMapper,
	userHelper helper.UserHelper,
) AuthService {
	return &authSvc{
		logger:     logger,
		fAuthCli:   fAuthCli,
		userRepo:   userRepo,
		userMapper: userMapper,
		userHelper: userHelper,
	}
}
