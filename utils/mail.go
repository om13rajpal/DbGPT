package utils

import (
	"github.com/om13rajpal/dbgpt/config"
	"gopkg.in/gomail.v2"
)

func SendMail(to string, subject string, body string) error {
	d := gomail.NewDialer("smtp.gmail.com", 587, config.EMAIL, config.EMAIL_PASSWORD)

	m := gomail.NewMessage()
	m.SetHeader("From", config.EMAIL)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
