package resend

// import (
// 	"os"
//
// 	"github.com/resendlabs/resend-go"
// )
//
// var client *resend.Client
//
// func getClient() *resend.Client {
//     if client == nil {
//         client = resend.NewClient(os.Getenv("RESEND_KEY"))
//         return client
//     }
//     return client
// }
//
// func Initialize() { getClient() }
//
// func SendEmail(subject string, to string, emailBody string) error {
//     params := &resend.SendEmailRequest{
//         From:    "myPro <no-reply@buscaenmypro.com>",
//         To:      []string{to},
//         Html:    emailBody,
//         Subject: subject,
//         // Cc:      []string{"cc@example.com"},
//         // Bcc:     []string{"bcc@example.com"},
//         // ReplyTo: "replyto@example.com",
//     }
//     client := getClient()
//     _, err := client.Emails.Send(params)
//     return err
// }
