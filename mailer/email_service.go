package mailer

import (
	"fmt"
	"net/smtp"
)

// EmailService represents an email service configuration.
type EmailService struct {
	Username string // Username for authentication with the SMTP server
	Password string // Password for authentication with the SMTP server
	SMTPHost string // SMTP server host address
	SMTPPort int    // SMTP server port
}

// NewEmailService creates a new EmailService instance with the provided configuration.
func NewEmailService(username string, password string, smtpHost string, smtpPort int) EmailService {
	return EmailService{
		Username: username,
		Password: password,
		SMTPHost: smtpHost,
		SMTPPort: smtpPort,
	}
}

// SendPlainEmail sends a plain text email with the provided subject and body to the specified recipient.
func (es EmailService) SendPlainEmail(to, subject, body string) error {
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	return es.sendEmail(to, message)
}

// SendHTMLEmail sends an HTML email with the provided subject and HTML body to the specified recipient.
func (es EmailService) SendHTMLEmail(to, subject, bodyHTML string) error {
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", to, subject, bodyHTML)
	return es.sendEmail(to, message)
}

// sendEmail sends an email with the specified message to the given recipient.
func (es EmailService) sendEmail(to, message string) error {
	auth := smtp.PlainAuth("", es.Username, es.Password, es.SMTPHost)
	addr := fmt.Sprintf("%s:%d", es.SMTPHost, es.SMTPPort)
	err := smtp.SendMail(addr, auth, es.Username, []string{to}, []byte(message))
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}
	return nil
}
