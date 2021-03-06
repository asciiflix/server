package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/asciiflix/server/config"
)

func SendWelcomeMail(to string, username string, code string) error {
	//Return if no server is configured
	if config.SMTP.Host == "" {
		return nil
	}
	//Authenticate with server using login and password
	auth := smtp.PlainAuth("", config.SMTP.User, config.SMTP.Password, config.SMTP.Host)

	//Parse template
	mail, _ := template.ParseFiles("./templates/welcome-mail.html")

	var body bytes.Buffer

	//Set MIME Header
	mimeHeaders := "MIME-version: 1.0; \nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Welcome to ASCIIflix\r\n%s\n\n", config.SMTP.User, to, mimeHeaders)))

	mail.Execute(&body, struct {
		Name       string
		Message    string
		URL        string
		URL_Verify string
	}{
		Name:       username,
		Message:    "Your verification Code is: " + code + " it expires in 2 days.",
		URL:        config.ApiConfig.FrontendURL + "/login",
		URL_Verify: config.ApiConfig.FrontendURL + "/verify",
	})

	//Send Mail
	err := smtp.SendMail(config.SMTP.Host+":"+config.SMTP.Port, auth, config.SMTP.User, []string{to}, body.Bytes())
	return err
}
