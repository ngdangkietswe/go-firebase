/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package service

import (
	"context"
	"go-firebase/internal/data/ent"
	"go-firebase/internal/firebase"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"
	"time"

	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type fileStorageSvc struct {
	cli         *ent.Client
	logger      *logger.Logger
	fStorageCli firebase.FStorageClient
}

func (s *fileStorageSvc) GenerateUploadURL(ctx context.Context, request *request.GetPresignURLRequest) (*response.PresignURLResponse, error) {
	expireDuration := constant.DefaultExpirePresignURLDuration
	if request.ExpireTime > 0 {
		expireDuration = time.Duration(request.ExpireTime) * time.Second
	}

	url, err := s.fStorageCli.UploadFile(request.BucketName, request.ObjectName, request.ContentType, expireDuration)
	if err != nil {
		return nil, err
	}

	return &response.PresignURLResponse{
		URL: url,
	}, nil
}

func (s *fileStorageSvc) GenerateDownloadURL(ctx context.Context, request *request.GetPresignURLRequest) (*response.PresignURLResponse, error) {
	expireDuration := constant.DefaultExpirePresignURLDuration
	if request.ExpireTime > 0 {
		expireDuration = time.Duration(request.ExpireTime) * time.Second
	}

	url, err := s.fStorageCli.DownloadFile(request.BucketName, request.ObjectName, expireDuration)
	if err != nil {
		return nil, err
	}

	return &response.PresignURLResponse{
		URL: url,
	}, nil
}

func NewFileStorageService(
	cli *ent.Client,
	logger *logger.Logger,
	fStorageCli firebase.FStorageClient,
) FileStorageService {
	return &fileStorageSvc{
		cli:         cli,
		logger:      logger,
		fStorageCli: fStorageCli,
	}
}
