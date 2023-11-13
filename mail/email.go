package helpers

import (
	"fmt"
	"net/smtp"
)

type EmailService struct {
	Username string
	Password string
	SMTPHost string
	SMTPPort int
}

func NewEmailService(username string, password string, smtpHost string, smtpPort int) *EmailService {
	return &EmailService{
		Username: username,
		Password: password,
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
	}
}

func (es *EmailService) SendPlainEmail(to, subject, body string) error {
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	return es.sendEmail(to, message)
}


func (es *EmailService) SendHTMLEmail(to, subject, bodyHTML string) error {
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", to, subject, bodyHTML)
	return es.sendEmail(to, message)
}

func (es *EmailService) sendEmail(to, message string) error {
	auth := smtp.PlainAuth("", es.Username, es.Password, es.SMTPHost)
	addr := fmt.Sprintf("%s:%d", es.SMTPHost, es.SMTPPort)
	err := smtp.SendMail(addr, auth, es.Username, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}
	return nil
}