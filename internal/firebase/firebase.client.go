/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package firebase

import (
	"context"
	"encoding/base64"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"firebase.google.com/go/v4/storage"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"go.uber.org/zap"
	"google.golang.org/api/option"

	"github.com/ngdangkietswe/swe-go-common-shared/logger"
)

type FirebaseApp struct {
	authCli      *auth.Client
	messagingCli *messaging.Client
	storageCli   *storage.Client
}

func NewFirebaseClient(logger *logger.Logger) *FirebaseApp {
	credEncoded := config.GetString("FIREBASE_CREDENTIALS_BASE64", "")
	if credEncoded == "" {
		return nil
	}

	credDecoded, err := base64.StdEncoding.DecodeString(credEncoded)
	if err != nil {
		logger.Warn("Failed to decode Firebase credentials from base64", zap.Error(err))
		return nil
	}

	opt := option.WithCredentialsJSON(credDecoded)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Warn("Failed to initialize Firebase app", zap.Error(err))
		return nil
	}

	authCli, err := app.Auth(context.Background())
	if err != nil {
		logger.Warn("Failed to initialize Firebase auth client", zap.Error(err))
		return nil
	}

	messagingCli, err := app.Messaging(context.Background())
	if err != nil {
		logger.Warn("Failed to initialize Firebase messaging client", zap.Error(err))
		return nil
	}

	storageCli, err := app.Storage(context.Background())
	if err != nil {
		logger.Warn("Failed to initialize Firebase storage client", zap.Error(err))
		return nil
	}

	return &FirebaseApp{
		authCli:      authCli,
		messagingCli: messagingCli,
		storageCli:   storageCli,
	}
}
