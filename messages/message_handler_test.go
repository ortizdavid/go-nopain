package messages

import (
	"errors"
	"testing"
)

type messageHandlerTest struct {
	err error 
}

func TestLogFailOnError(t *testing.T) {
	testCases := [] messageHandlerTest {
		{ err: errors.New("Error Loading Application") },
		{ err: errors.New("Application Crashed") },
		{ err: errors.New("Error.. Please refresh application") },
	}
	for _, test := range testCases {
		LogFailOnError(test.err)
	}
}

func TestPrintOnError(t *testing.T) {
	testCases := [] messageHandlerTest {
		{ err: errors.New("Error Loading Application") },
		{ err: errors.New("Your Application Crashed") },
		{ err: errors.New("Error.. Please refresh application!") },
		{ err: errors.New("Unknown Error") },
	}
	for _, test := range testCases {
		PrintOnError(test.err)
	}
}


func TestPanicOnError(t *testing.T) {
	testCases := [] messageHandlerTest {
		{ err: errors.New("Error Loading Application") },
		{ err: errors.New("Your Application Crashed") },
		{ err: errors.New("Error.. Please refresh application!") },
		{ err: errors.New("Unknown Error") },
	}
	for _, test := range testCases {
		PanicOnError(test.err)
	}
}