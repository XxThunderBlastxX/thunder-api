package helper

import (
	"net/smtp"

	"github.com/XxThunderBlastxX/thunder-api/domain"
	"github.com/XxThunderBlastxX/thunder-api/internal/gen/smtpconfig"
)

func SendMail(msg domain.Message, smtpConfig *smtpconfig.SMTPConfig, receiverEmail ...string) error {
	// Email credentials
	senderEmail := smtpConfig.User
	mailPass := smtpConfig.Password

	// SMTP server configuration.
	smtpHost := smtpConfig.Host
	smtpPort := smtpConfig.Port

	// MIME type
	mime := "Content-Type: text/html; charset=UTF-8\r\n\r\n"

	// template
	emailBody, err := BuildEmailTmpl(msg, "./pkg/template/email.html")
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
