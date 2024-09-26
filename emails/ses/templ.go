package ses

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
)

func (service *SES) SendEmailTempl(ctx context.Context, recipient string, subject string, template templ.Component) error {
    buf := new(bytes.Buffer)
    err := template.Render(ctx, buf)
    if err != nil {
        return err
    }
    body := buf.String()
    return service.SendEmail(ctx, recipient, subject, body, body)
}
