package logging

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "test-logging")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a logger instance
	logger := NewLogger(tmpDir, "test.log")

	// Check if logger instance is not nil
	if logger == nil {
		t.Error("Expected logger instance, got nil")
	}
}
