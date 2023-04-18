package mailer

import (
	// "crypto/tls"
	// "encoding/base64"
	"context"
    "fmt"
    "net/smtp"
    "os"
    // "strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

func Send(message string) bool {
	// user := os.Getenv("MAIL_USER")
	from := os.Getenv("MAIL_FROM")
	recipients := []string{os.Getenv("MAIL_RECIPIENTS")}
	password := os.Getenv("MAIL_PASSWORD")
	smtpHost := os.Getenv("MAIL_SMTP_HOST")
	smtpPort := os.Getenv("MAIL_SMTP_PORT")
	region := "ap-northeast-1"
	subject := "You've gotten an message from a user!"

	if os.Getenv("ENVIRONMENT") == "production" {
	  ctx := context.Background()
	  cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	  if err != nil {
		fmt.Println(err)
	  }
	  client := sesv2.NewFromConfig(cfg)
	  input := &sesv2.SendEmailInput{
		FromEmailAddress: &from,
		Destination: &types.Destination{
		  ToAddresses: recipients,
		},
		Content: &types.EmailContent{
		  Simple: &types.Message{
			  Body: &types.Body{
				  Text: &types.Content{
					  Data: &message,
				  },
			  },
			  Subject: &types.Content{
				Data: &subject,
			  },
		  },
		},
	  }
	  res, err := client.SendEmail(ctx, input)
	  if err != nil {
		fmt.Println(err)
		return false
	  }
	  fmt.Println(res.MessageId)
	  return true
	} else {
	  auth := smtp.CRAMMD5Auth(from, password)
	  if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, []byte(message)); err != nil {
		fmt.Println(err)
		return false
	  }
	
	  return true
	}
	
}

// aws側はsdkを使うことにする