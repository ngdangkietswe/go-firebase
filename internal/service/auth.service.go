/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package service

import (
	"context"
	"go-firebase/internal/firebase"
	"go-firebase/internal/request"
	"go-firebase/internal/response"

	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type authSvc struct {
	logger   *logger.Logger
	fAuthCli firebase.FAuthClient
}

func (s *authSvc) Login(ctx context.Context, request *request.LoginRequest) (*response.LoginResponse, error) {
	fAuthLoginResp, err := s.fAuthCli.Login(request)
	if err != nil {
		return nil, err
	}

	resp := &response.LoginResponse{
		IDToken:      fAuthLoginResp["idToken"].(string),
		RefreshToken: fAuthLoginResp["refreshToken"].(string),
		ExpiresIn:    fAuthLoginResp["expiresIn"].(string),
		LocalID:      fAuthLoginResp["localId"].(string),
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

func NewAuthService(
	logger *logger.Logger,
	fAuthCli firebase.FAuthClient,
) AuthService {
	return &authSvc{
		logger:   logger,
		fAuthCli: fAuthCli,
	}
}
