package mailer

import (
    "fmt"
    "net/smtp"
    // "os"
    // "strings"
)

func Send(message string) {
	fmt.Println(message)
	from := "xxx"
	recipients := []string{"xxx"}
	password := "xxx"
	smtpHost := "xxx"
	smtpPort := "xxx" //環境変数

	auth := smtp.PlainAuth("", from, password, smtpHost)

	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, recipients, []byte(message)); err != nil {
	  panic(err.Error())
	}
}