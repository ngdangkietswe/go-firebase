/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package job

import (
	"context"
	"encoding/json"
	"fmt"
	"go-firebase/internal/service"
	"go-firebase/pkg/constant"
	"go-firebase/pkg/request"
	"net/http"
	"time"

	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type TopicTechJobProps struct {
	fx.In
	Logger              *logger.Logger
	Cron                *cron.Cron
	NotificationService service.NotificationService
}

const NewsTechAPIURL = "https://newsapi.org/v2/top-headlines?country=us&category=technology&apiKey="

func RegisterTopicTechJob(props TopicTechJobProps) {
	_, err := props.Cron.AddFunc("0 9 * * *", func() {
		props.Logger.Info("Running send notification to topic tech job")

		httpReq, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			fmt.Sprintf("%s%s", NewsTechAPIURL, config.GetString("NEWS_API_KEY", "")),
			nil,
		)
		if err != nil {
			props.Logger.Error("Failed to create request for news API", zap.Error(err))
		}

		client := &http.Client{Timeout: 10 * time.Second}

		httpResp, err := client.Do(httpReq)
		if err != nil {
			props.Logger.Error("Failed to fetch news from API", zap.Error(err))
			return
		}
		defer httpResp.Body.Close()

		if httpResp.StatusCode != http.StatusOK {
			props.Logger.Error("Non-OK HTTP status from news API", zap.Int("status_code", httpResp.StatusCode))
			return
		}

		var respData map[string]interface{}

		err = json.NewDecoder(httpResp.Body).Decode(&respData)
		if err != nil {
			props.Logger.Error("Failed to decode news API response", zap.Error(err))
			return
		}

		articles, ok := respData["articles"].([]interface{})
		if !ok || len(articles) == 0 {
			props.Logger.Info("No articles found in news API response")
			return
		}

		firstArticle, ok := articles[0].(map[string]interface{})
		if !ok {
			props.Logger.Info("Invalid article format in news API response")
			return
		}

		title, _ := firstArticle["title"].(string)
		description, _ := firstArticle["description"].(string)

		sendNotificationRequest := &request.SendNotificationRequest{
			TopicName: constant.NotificationTopicTech,
			Title:     title,
			Body:      description,
			Payload: map[string]string{
				"source": "NewsAPI",
				"url":    firstArticle["url"].(string),
				"image":  firstArticle["urlToImage"].(string),
			},
		}

		_, err = props.NotificationService.SendNotification(context.Background(), sendNotificationRequest)
		if err != nil {
			props.Logger.Error("Failed to send notification to topic tech", zap.Error(err))
			return
		}

		props.Logger.Info("Successfully sent notification to topic tech")
	})
	if err != nil {
		return
	}
}
