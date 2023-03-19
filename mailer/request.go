package mailer

import (
    "fmt"
    "net/smtp"
    "os"
    // "strings"
)

func Send(message string) bool {
	from := os.Getenv("MAIL_FROM")
	recipients := []string{os.Getenv("MAIL_RECIPIENTS")}
	password := os.Getenv("MAIL_PASSWORD")
	smtpHost := os.Getenv("MAIL_SMTP_HOST")
	smtpPort := os.Getenv("MAIL_SMTP_PORT")
	fmt.Println(password)
	fmt.Println(smtpHost)
	fmt.Println(smtpPort)
	auth := smtp.CRAMMD5Auth(from, password)
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, []byte(message)); err != nil {
	  return false
	}

	return true
}