package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"go-email-verifier-tool/config"
	"io"

	"gopkg.in/gomail.v2"
)

var(
	ErrFailedGenerateToken = errors.New("failed to generate token")
)

func SendMail(toEmail string, subject string, body string) error {
	emailConfig, err := config.NewEmailConfig()
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailConfig.Mail)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(emailConfig.Host, emailConfig.Port, emailConfig.Mail, emailConfig.Password)

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}

func MakeVerificationEmail(receiverEmail string, mail []byte) (map[string]string, error) {
	token := EncodeToString(6)
	if token == "" {
		return nil, ErrFailedGenerateToken
	}

	str := string(mail)

	draftEmail := map[string]string{}
	draftEmail["subject"] = "Template - Email Verification"
	draftEmail["token"] = token 
	draftEmail["body"] = fmt.Sprintf(str, token, receiverEmail)

	return draftEmail, nil
}

func EncodeToString(max int) string {
	var data = [...] byte {'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	b := make([]byte, max)
	n, _ := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return ""
	}

	for i := 0; i < len(b); i++ {
		b[i] = data[int(b[i]) % len(data)]
	}

	return string(b)
}