/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package firebase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go-firebase/internal/request"
	"go-firebase/pkg/constant"
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/zap"
)

type fAuthCli struct {
	logger  *logger.Logger
	authCli *auth.Client
}

const FAuthURL = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key="

func (c *fAuthCli) Login(request *request.LoginRequest) (map[string]interface{}, error) {
	apiKey := config.GetString("FIREBASE_API_KEY", "")

	reqBody := map[string]interface{}{
		"email":             request.Email,
		"password":          request.Password,
		"returnSecureToken": true,
	}

	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		c.logger.Error("Failed to marshal login request body", zap.Error(err))
		return nil, err
	}

	httpCli := http.Client{Timeout: time.Second * 10}

	resp, err := httpCli.Post(FAuthURL+apiKey, "application/json", bytes.NewBuffer(reqBodyByte))
	if err != nil {
		c.logger.Error("Failed to send login request to firebase", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.logger.Error("Firebase login request failed", zap.Int("statusCode", resp.StatusCode))
		return nil, errors.New("firebase login request failed")
	}

	var respData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		c.logger.Error("Failed to decode firebase login response", zap.Error(err))
		return nil, err
	}

	return respData, nil
}

func (c *fAuthCli) Signup(request *request.CreateUserRequest) (string, error) {
	params := c.toFirebaseUserParams(request)

	record, err := c.authCli.CreateUser(context.Background(), params)
	if err != nil {
		c.logger.Error("Failed to create firebase user", zap.Error(err))
		return "", err
	}

	c.logger.Info("Successfully created firebase user", zap.String("firebaseUID", record.UID))
	return record.UID, nil
}

func (c *fAuthCli) toFirebaseUserParams(request *request.CreateUserRequest) *auth.UserToCreate {
	params := (&auth.UserToCreate{}).
		Email(request.Email)

	password := constant.DefaultPassword
	if request.Password != "" {
		password = request.Password
	}

	params.Password(password)

	if request.FirstName != "" && request.LastName != "" {
		params.DisplayName(request.FirstName + " " + request.LastName)
	}

	return params
}

func (c *fAuthCli) VerifyIDToken(idToken string) (string, error) {
	fToken, err := c.authCli.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.logger.Error("Failed to verify firebase ID token", zap.Error(err))
		return "", err
	}
	return fToken.UID, nil
}

func NewFAuthClient(
	logger *logger.Logger,
	firebaseApp *FirebaseApp,
) FAuthClient {
	return &fAuthCli{
		logger:  logger,
		authCli: firebaseApp.authCli,
	}
}
