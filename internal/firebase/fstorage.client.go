/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package firebase

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/zap"
)

type fStorageClient struct {
	logger               *logger.Logger
	defaultBucket        string
	googleSvcClientEmail string
	googleSvcPrivateKey  string
}

func (c *fStorageClient) UploadFile(bucketName string, objectName string, contentType string, expireDuration time.Duration) (string, error) {
	rsaKey, err := c.parseRSAPrivateKey(c.googleSvcPrivateKey)
	if err != nil {
		return "", err
	}

	opts := storage.SignedURLOptions{
		GoogleAccessID: c.googleSvcClientEmail,
		PrivateKey:     x509.MarshalPKCS1PrivateKey(rsaKey),
		Method:         "PUT",
		Expires:        time.Now().Add(expireDuration),
		ContentType:    contentType,
	}

	url, err := storage.SignedURL(c.defaultBucket, objectName, &opts)
	if err != nil {
		c.logger.Error("Failed to generate signed URL for upload", zap.Error(err))
		return "", err
	}

	return url, nil
}

func (c *fStorageClient) DownloadFile(bucketName string, objectName string, expireDuration time.Duration) (string, error) {
	rsaKey, err := c.parseRSAPrivateKey(c.googleSvcPrivateKey)
	if err != nil {
		return "", err
	}

	opts := storage.SignedURLOptions{
		GoogleAccessID: c.googleSvcClientEmail,
		PrivateKey:     x509.MarshalPKCS1PrivateKey(rsaKey),
		Method:         "GET",
		Expires:        time.Now().Add(expireDuration),
	}

	url, err := storage.SignedURL(c.defaultBucket, objectName, &opts)
	if err != nil {
		c.logger.Error("Failed to generate signed URL for download", zap.Error(err))
		return "", err
	}

	return url, nil
}

func (c *fStorageClient) parseRSAPrivateKey(pemEncodedKey string) (*rsa.PrivateKey, error) {
	normalizedKey := strings.ReplaceAll(pemEncodedKey, "\\n", "\n")
	block, _ := pem.Decode([]byte(normalizedKey))
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	// PKCS8 -> PKCS1
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err == nil {
		if rsaKey, ok := parsedKey.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func NewFStorageClient(
	logger *logger.Logger,
) FStorageClient {
	return &fStorageClient{
		logger:               logger,
		defaultBucket:        config.GetString("FIREBASE_STORAGE_BUCKET", ""),
		googleSvcClientEmail: config.GetString("GOOGLE_SERVICE_CLIENT_EMAIL", ""),
		googleSvcPrivateKey:  config.GetString("GOOGLE_SERVICE_PRIVATE_KEY", ""),
	}
}
