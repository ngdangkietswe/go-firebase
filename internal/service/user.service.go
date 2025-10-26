/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package service

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/data/repository"
	"go-firebase/internal/firebase"
	"go-firebase/internal/helper"
	"go-firebase/internal/mapper"
	"go-firebase/pkg/model"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"go-firebase/pkg/util"

	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/zap"
)

type userSvc struct {
	cli        *ent.Client
	logger     *logger.Logger
	fAuthCli   firebase.FAuthClient
	userMapper mapper.UserMapper
	userRepo   repository.UserRepository
	userHelper helper.UserHelper
}

func (s *userSvc) CreateUser(ctx context.Context, request *request.CreateUserRequest) (*response.CreateUserResponse, error) {
	firebaseUID, err := s.fAuthCli.Signup(request)
	if err != nil {
		return nil, err
	}

	user, err := repository.WithTxResult(ctx, s.cli, s.logger, func(tx *ent.Tx) (*ent.User, error) {
		return s.userRepo.Save(ctx, tx, request, firebaseUID)
	})

	if err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return nil, err
	}

	return &response.CreateUserResponse{
		UserID:      user.ID.String(),
		FirebaseUID: firebaseUID,
	}, nil
}

func (s *userSvc) GetUsers(ctx context.Context, request *request.ListUserRequest) (*response.ListResponse, error) {
	util.NormalizePaginationRequest(request.Paginate)

	items, totalItems, err := s.userRepo.FindAll(ctx, request)
	if err != nil {
		s.logger.Error("Failed to get users", zap.Error(err))
		return nil, err
	}

	mUsers := s.userMapper.AsList(items)

	s.userHelper.Preload(ctx, mUsers, request.Preload)

	return &response.ListResponse{
		Items: mUsers,
		Meta:  util.AsPageMeta(request.Paginate, totalItems),
	}, nil
}

func (s *userSvc) GetUser(ctx context.Context, request *request.GetUserRequest) (*model.User, error) {
	user, err := s.userRepo.FindByEmailOrID(ctx, request.Identifier)
	if err != nil {
		s.logger.Error("Failed to get user", zap.String("identifier", request.Identifier), zap.Error(err))
		return nil, err
	}

	mUser := s.userMapper.AsMono(user)

	s.userHelper.Preload(ctx, []*model.User{mUser}, request.Preload)

	return mUser, nil
}

func NewUserService(
	cli *ent.Client,
	logger *logger.Logger,
	fAuthCli firebase.FAuthClient,
	userMapper mapper.UserMapper,
	userRepo repository.UserRepository,
	userHelper helper.UserHelper,
) UserService {
	return &userSvc{
		cli:        cli,
		userMapper: userMapper,
		fAuthCli:   fAuthCli,
		logger:     logger,
		userRepo:   userRepo,
		userHelper: userHelper,
	}
}
