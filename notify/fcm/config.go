package fcm

import (
	"context"
	"log/slog"
	"os"

	"firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

type FCM struct {
    logger *slog.Logger
    client *firebase.App
}

func Initialize(logger *slog.Logger) *FCM {
    credentials := os.Getenv("GOOGLE_FIREBASE_CONFIG")
    app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsJSON([]byte(credentials)))
    if err != nil {
        logger.Error("creating app", "service", "firebase", "error", err.Error())
        return nil
    }
    logger.Info("initialized", "service", "firebase")
    return &FCM{ client: app, logger: logger }
}
