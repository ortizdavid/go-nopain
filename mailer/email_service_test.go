package mailer

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	// Initialize EmailService with test SMTP server details
	es := NewEmailService("testuser", "testpassword", "test.smtp.com", 587)

	// Test valid email sending
	err := es.SendPlainEmail("recipient@example.com", "Test Subject", "Test Body")
	if err != nil {
		t.Errorf("Error sending email: %v", err)
	}

	// Test sending email to an invalid recipient
	err = es.SendPlainEmail("", "Test Subject", "Test Body")
	if err == nil {
		t.Error("Expected error for empty recipient address, got nil")
	}

	// Test sending email with an empty message body
	err = es.SendPlainEmail("recipient@example.com", "Test Subject", "")
	if err == nil {
		t.Error("Expected error for empty message body, got nil")
	}
}
