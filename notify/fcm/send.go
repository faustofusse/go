package fcm

import (
	"context"

	"firebase.google.com/go/v4/messaging"
)

func (service *FCM) buildMessage(title, body, image, token string, data map[string]string) *messaging.Message {
    return &messaging.Message{
        Token: token,
        Data: data,
        Android: &messaging.AndroidConfig{
            CollapseKey: title,
            Notification: &messaging.AndroidNotification{
                // ImageURL: "https://avatars.githubusercontent.com/u/15711090?v=4",
            },
        },
        Notification: &messaging.Notification{
            Title: title,
            Body: body,
        },
    }
}

func (service *FCM) SendAll(ctx context.Context, title, body, image string, data map[string]string, tokens []string) error {
    batches := [][]*messaging.Message{}
    for _, token := range tokens {
        message := service.buildMessage(title, body, image, token, data)
        added := false
        for i, batch := range batches {
            if len(batch) < 500 {
                batches[i] = append(batch, message)
                added = true
            }
        }
        if !added {
            batches = append(batches, []*messaging.Message{ message })
        }
    }
    m, err := service.client.Messaging(ctx)
    if err != nil {
        service.logger.Error("initializing messaging service", "service", "firebase", "error", err.Error())
        return err
    }
    for _, batch := range batches {
        service.logger.Info("sending messages batch", "service", "firebase", "amount", len(batch))
        a, err := m.SendEach(ctx, batch)
        if err != nil {
            service.logger.Error("sending messages batch", "service", "firebase", "error", err.Error())
            return err
        } else {
            service.logger.Info("sent messages batch", "service", "firebase", "successful", a.SuccessCount, "failed", a.FailureCount)
            if a.FailureCount > 0 {
                for _, response := range a.Responses {
                    if !response.Success {
                        service.logger.Error("sending individual message", "service", "firebase", "error", response.Error.Error())
                    }
                }
            }
        }
    }
    return nil
}

// TODO: hacer un m.SendAll para los mensajes
func (service *FCM) Send(ctx context.Context, title, body, image, token string, data map[string]string) error {
    message := service.buildMessage(title, body, image, token, data)
    m, err := service.client.Messaging(ctx)
    if err != nil {
        service.logger.Error("initializing messaging service", "service", "firebase", "error", err.Error())
        return err
    }
    a, err := m.Send(ctx, message)
    if err != nil {
        service.logger.Error("sending message", "service", "firebase", "error", err.Error())
    } else {
        service.logger.Info("message sent", "service", "firebase", "response", a)
    }
    return err
}
