package logging

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_NewLogger(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "test-logging")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)
	logger := NewLogger(tmpDir, "test.log")
	if logger == nil {
		t.Error("Expected logger instance, got nil")
	}
}
