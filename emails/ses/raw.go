package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func (service *SES) SendEmail(ctx context.Context, recipient string, subject string, htmlBody string, textBody string) error {
    params := &ses.SendEmailInput{
        Destination: &types.Destination{
            CcAddresses: []string{ },
            ToAddresses: []string{ recipient },
        },
        Message: &types.Message{
            Body: &types.Body{
                Html: &types.Content{
                    Charset: aws.String(charSet),
                    Data:    aws.String(htmlBody),
                },
                Text: &types.Content{
                    Charset: aws.String(charSet),
                    Data:    aws.String(textBody),
                },
            },
            Subject: &types.Content{
                Charset: aws.String(charSet),
                Data:    aws.String(subject),
            },
        },
        Source: aws.String(service.sender),
    }
    output, err := service.client.SendEmail(ctx, params)
    if err != nil {
        service.logger.Error("sending email", "service", "ses", "error", err.Error())
    } else {
        service.logger.Info("email sent", "service", "ses", "output", output)
    }
    return err
}
