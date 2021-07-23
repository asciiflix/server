package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/asciiflix/server/config"
)

func SendWelcomeMail(to string, username string) error {

	//Authenticate with server using login and password
	auth := smtp.PlainAuth("", config.SMTP.User, config.SMTP.Password, config.SMTP.Host)

	mail, _ := template.ParseFiles("templates/welcome-mail.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Welcome to ASCIIflix\n%s\n\n", mimeHeaders)))

	mail.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    username,
		Message: "Yoo whats up!!",
	})

	//SendMail
	err := smtp.SendMail(config.SMTP.Host+":"+config.SMTP.Port, auth, config.SMTP.User, []string{to}, body.Bytes())
	return err
}
