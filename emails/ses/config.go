package ses

import (
	"context"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

const charSet = "UTF-8"

type SES struct {
    sender string
    client *ses.Client
    logger *slog.Logger
}

func Initialize(logger *slog.Logger) *SES {
    service := new(SES)
    service.logger = logger
    cfg, err := config.LoadDefaultConfig(context.Background())
    if err != nil {
        service.logger.Error("initializing", "service", "ses", "error", err.Error())
        os.Exit(1)
    } else {
        service.logger.Info("initialized", "service", "ses")
    }
    service.client = ses.NewFromConfig(cfg)
    service.sender = os.Getenv("EMAIL_SENDER")
    return service
}
