/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	authCli      *auth.Client
	messagingCli *messaging.Client
}

func NewFirebaseClient() *FirebaseApp {
	cred := config.GetString("FIREBASE_CREDENTIALS_PATH", "")
	if cred == "" {
		return nil
	}

	opt := option.WithCredentialsFile(cred)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil
	}

	authCli, err := app.Auth(context.Background())
	if err != nil {
		return nil
	}

	messagingCli, err := app.Messaging(context.Background())
	if err != nil {
		return nil
	}

	return &FirebaseApp{
		authCli:      authCli,
		messagingCli: messagingCli,
	}
}
