package fcm

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"

	"firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type FCM struct {
    logger *slog.Logger
    client *firebase.App
}

func getCredentials() map[string]string {
    return map[string]string{
        "type": os.Getenv("FIREBASE_TYPE"),
        "project_id": os.Getenv("FIREBASE_PROJECT_ID"),
        "private_key_id": os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
        "private_key": os.Getenv("FIREBASE_PRIVATE_KEY"),
        "client_email": os.Getenv("FIREBASE_CLIENT_EMAIL"),
        "client_id": os.Getenv("FIREBASE_CLIENT_ID"),
        "auth_uri": os.Getenv("FIREBASE_AUTH_URI"),
        "token_uri": os.Getenv("FIREBASE_TOKEN_URI"),
        "auth_provider_x509_cert_url": os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
        "client_x509_cert_url": os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
        "universe_domain": os.Getenv("FIREBASE_UNIVERSE_DOMAIN"),
    }
}

func Initialize(logger *slog.Logger) *FCM {
    key := os.Getenv("GOOGLE_FIREBASE_KEY")
    projectId := os.Getenv("GOOGLE_FIREBASE_PROJECT_ID")
    options := option.WithAPIKey(key)
    if len(key) == 0 {
        credentials := getCredentials()
        projectId = credentials["project_id"]
        bytes, _ := json.Marshal(credentials)
        options = option.WithCredentialsJSON(bytes)
    }
    config := firebase.Config{ ProjectID: projectId }
    app, err := firebase.NewApp(context.Background(), &config, options)
    if err != nil {
        logger.Error("creating app", "service", "firebase", "error", err.Error())
        return nil
    }
    logger.Info("initialized", "service", "fcm")
    return &FCM{ client: app, logger: logger }
}
