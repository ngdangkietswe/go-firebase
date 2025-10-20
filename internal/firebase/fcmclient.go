/**
 * Author : ngdangkietswe
 * Since  : 10/20/2025
 */

package firebase

import (
	"context"

	"firebase.google.com/go/v4/messaging"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"go.uber.org/zap"
)

type fcmClient struct {
	logger       *logger.Logger
	messagingCli *messaging.Client
}

func (c *fcmClient) SendToToken(token string, title string, body string, data map[string]string) error {
	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data:  data,
		Token: token,
	}

	_, err := c.messagingCli.Send(context.Background(), msg)
	if err != nil {
		c.logger.Error("Failed to send firebase notification", zap.String("error", err.Error()))
		return err
	}

	c.logger.Info("Successfully sent firebase notification", zap.String("token", token))
	return nil
}

func (c *fcmClient) SendToTokens(tokens []string, title string, body string, data map[string]string) error {
	msg := &messaging.MulticastMessage{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data:   data,
		Tokens: tokens,
	}

	res, err := c.messagingCli.SendEachForMulticast(context.Background(), msg)
	if err != nil {
		c.logger.Error("Failed to send firebase notification to multiple tokens", zap.String("error", err.Error()))
		return err
	}

	c.logger.Info("Successfully sent firebase notification to multiple tokens", zap.Int("successCount", res.SuccessCount), zap.Int("failureCount", res.FailureCount))
	return nil
}

func NewFCMClient(
	logger *logger.Logger,
	firebaseApp *FirebaseApp,
) FCMClient {
	return &fcmClient{
		logger:       logger,
		messagingCli: firebaseApp.messagingCli,
	}
}
