package utils

import (
	"github.com/om13rajpal/dbgpt/config"
	"gopkg.in/gomail.v2"
	"fmt"
)

var d = gomail.NewDialer("smtp.gmail.com", 587, config.EMAIL, config.EMAIL_PASSWORD)

func SendMail(to string, subject string, body string) error {

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

func SendOTP(to string, otp int) error {

	m := gomail.NewMessage()
	m.SetHeader("From", config.EMAIL)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your OTP for DbGPT")
	m.SetBody("text/html", "<h1>Your OTP is: "+fmt.Sprint(otp)+"</h1>")

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}

	return nil
}
