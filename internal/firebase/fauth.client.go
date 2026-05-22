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
	"go-firebase/pkg/constant"
	"go-firebase/pkg/request"
	"net/http"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/zap"
)

type fAuthCli struct {
	logger     *logger.Logger
	authCli    *auth.Client
	apiKey     string
	fProjectID string
}

const (
	FSignInWithPasswordURL    = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key="
	FSignInWithCustomTokenURL = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key="
	FRefreshTokenURL          = "https://securetoken.googleapis.com/v1/token?key="
	FSendOobCodeURL           = "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key="
	FResetPasswordURL         = "https://identitytoolkit.googleapis.com/v1/accounts:resetPassword?key="
)

const (
	RequestTypePasswordReset = "PASSWORD_RESET"
)

const (
	GrantTypeRefreshToken = "refresh_token"
)

func (c *fAuthCli) LoginWithPassword(request *request.LoginRequest) (map[string]interface{}, error) {
	reqBody := map[string]interface{}{
		"email":             request.Email,
		"password":          request.Password,
		"returnSecureToken": true,
	}

	loginResp, err := postFirebase(FSignInWithPasswordURL+c.apiKey, reqBody)
	if err != nil {
		c.logger.Error("Failed to login with firebase email and password", zap.Error(err))
		return nil, err
	}

	return loginResp, nil
}

func (c *fAuthCli) LoginWithCustomToken(customToken string) (map[string]interface{}, error) {
	reqBody := map[string]interface{}{
		"token":             customToken,
		"returnSecureToken": true,
	}

	loginResp, err := postFirebase(FSignInWithCustomTokenURL+c.apiKey, reqBody)
	if err != nil {
		c.logger.Error("Failed to login with firebase custom token", zap.Error(err))
		return nil, err
	}

	return loginResp, nil
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

func (c *fAuthCli) VerifyIDToken(idToken string) (map[string]interface{}, error) {
	fToken, err := c.authCli.VerifyIDTokenAndCheckRevoked(context.Background(), idToken)
	if err != nil {
		if auth.IsIDTokenRevoked(err) {
			c.logger.Warn("Firebase ID token has been revoked", zap.Error(err))
			return nil, errors.New("firebase ID token has been revoked")
		}
		c.logger.Error("Failed to verify firebase ID token", zap.Error(err))
		return nil, err
	}
	return fToken.Claims, nil
}

func (c *fAuthCli) RefreshToken(request *request.RefreshTokenRequest) (map[string]interface{}, error) {
	reqBody := map[string]interface{}{
		"grant_type":    GrantTypeRefreshToken,
		"refresh_token": request.RefreshToken,
	}

	refreshResp, err := postFirebase(FRefreshTokenURL+c.apiKey, reqBody)
	if err != nil {
		c.logger.Error("Failed to refresh firebase token", zap.Error(err))
		return nil, err
	}

	return refreshResp, nil
}

func (c *fAuthCli) RevokeToken(request *request.RevokeTokenRequest) error {
	if err := c.authCli.RevokeRefreshTokens(context.Background(), request.FirebaseUID); err != nil {
		c.logger.Error("Failed to revoke firebase token", zap.Error(err))
		return err
	}
	return nil
}

func (c *fAuthCli) CustomToken(claims map[string]interface{}) (string, error) {
	return c.authCli.CustomTokenWithClaims(context.Background(), claims["firebase_uid"].(string), claims)
}

func (c *fAuthCli) EnDisableAccount(firebaseUID string, disabled bool) error {
	params := (&auth.UserToUpdate{}).Disabled(disabled)

	if _, err := c.authCli.UpdateUser(context.Background(), firebaseUID, params); err != nil {
		c.logger.Error("Failed to enable/disable firebase account", zap.Error(err))
		return err
	}

	return nil
}

func (c *fAuthCli) DeleteAccount(firebaseUID string) error {
	if err := c.authCli.DeleteUser(context.Background(), firebaseUID); err != nil {
		c.logger.Error("Failed to delete firebase account", zap.Error(err))
		return err
	}
	return nil
}

func (c *fAuthCli) SendPasswordResetEmail(request *request.SendPasswordResetMailRequest) error {
	reqBody := map[string]interface{}{
		"email":       request.Email,
		"requestType": RequestTypePasswordReset,
	}

	_, err := postFirebase(FSendOobCodeURL+c.apiKey, reqBody)
	if err != nil {
		c.logger.Error("Failed to send password reset email", zap.Error(err))
		return err
	}

	c.logger.Info("Successfully sent password reset email", zap.String("email", request.Email))

	return nil
}

func (c *fAuthCli) ResetPassword(request *request.ResetPasswordRequest) error {
	reqBody := map[string]interface{}{
		"oobCode":     request.OobCode,
		"newPassword": request.NewPassword,
	}

	_, err := postFirebase(FResetPasswordURL+c.apiKey, reqBody)
	if err != nil {
		c.logger.Error("Failed to reset password", zap.Error(err))
		return err
	}

	c.logger.Info("Successfully reset password via oobCode", zap.String("oobCode", request.OobCode))

	return nil
}

func (c *fAuthCli) ChangePassword(request *request.ChangePasswordRequest) error {
	params := (&auth.UserToUpdate{}).Password(request.NewPassword)

	if _, err := c.authCli.UpdateUser(context.Background(), request.FirebaseUID, params); err != nil {
		c.logger.Error("Failed to change password", zap.String("firebaseUID", request.FirebaseUID), zap.Error(err))
		return err
	}

	c.logger.Info("Successfully changed password", zap.String("firebaseUID", request.FirebaseUID))

	return nil
}

func postFirebase(url string, body map[string]interface{}) (map[string]interface{}, error) {
	bodyBytes, _ := json.Marshal(body)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp map[string]interface{}
		_ = json.NewDecoder(resp.Body).Decode(&errResp)

		if errMsg, ok := errResp["error"].(map[string]interface{}); ok && errMsg["message"] != nil {
			return nil, errors.New(errMsg["message"].(string))
		}

		return nil, errors.New("firebase request failed")
	}

	var respData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func NewFAuthClient(
	logger *logger.Logger,
	firebaseApp *FirebaseApp,
) FAuthClient {
	return &fAuthCli{
		logger:     logger,
		authCli:    firebaseApp.authCli,
		apiKey:     config.GetString("FIREBASE_API_KEY", ""),
		fProjectID: config.GetString("FIREBASE_PROJECT_ID", ""),
	}
}
