package messages

import (
	"errors"
	"testing"
)

type messageHandlerTest struct {
	err error 
}

func TestLogFailOrError(t *testing.T) {
	testCases := [] messageHandlerTest {
		{ err: errors.New("Error Loading Application") },
		{ err: errors.New("Application Crashed") },
		{ err: errors.New("Error.. Please refresh application") },
	}
	for _, test := range testCases {
		LogFailOrError(test.err)
	}
}

func TestPrintFailOrError(t *testing.T) {
	testCases := [] messageHandlerTest {
		{ err: errors.New("Error Loading Application") },
		{ err: errors.New("Your Application Crashed") },
		{ err: errors.New("Error.. Please refresh application!") },
		{ err: errors.New("Unknown Error") },
	}
	for _, test := range testCases {
		PrintFailOrError(test.err)
	}
}