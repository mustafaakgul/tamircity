package middleware

import (
	"fmt"
	"net/smtp"
)

type EmailManager struct {
}

func SendEmail4Agreement(emailTo string) error {
	from := "YOUR_ACC"
	password := "YOUR_PASS"

	to := []string{
		emailTo,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("Agreement email")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}
