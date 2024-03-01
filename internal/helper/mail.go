package helper

import (
	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/XxThunderBlast/thunder-api/internal/global"
	"net/smtp"
)

func SendMail(msg domain.Message, receiverEmail ...string) error {
	// Email credentials
	senderEmail := global.Env.Email
	mailPass := global.Env.EmailPass

	// SMTP server configuration.
	smtpHost := "smtp.zoho.in"
	smtpPort := "587"

	// MIME type
	mime := "Content-Type: text/html; charset=UTF-8\r\n\r\n"

	// template
	emailBody, err := BuildEmailTmpl(msg, "./template/message_tmpl.html")
	if err != nil {
		return err
	}

	email := []byte(
		"From:" + "Portfolio" + "<api@koustav.dev>\r\n" +
			"To: me@koustav.dev\r\n" +
			"Subject: New message from " + msg.Name + "\r\n" +
			mime +
			emailBody.String() + "\r\n")

	// Auth with smtp server
	auth := smtp.PlainAuth("", senderEmail, mailPass, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, receiverEmail, email)
	if err != nil {
		return err
	}

	return nil
}
