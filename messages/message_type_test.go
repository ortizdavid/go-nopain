package messages

import (
	"testing"
)

type messageTypeTest struct {
	message string
}

func Test_ErrorMessage(t *testing.T) {
	testCases := [] messageTypeTest {
		{ message: "Error While Loading Module" },
		{ message: "Application Crashed" },
		{ message: "Error.. Please refresh application" },
	}
	for _, test := range testCases {
		Error(test.message)
	}
}

func Test_WarningMessage(t *testing.T) {
	testCases := [] messageTypeTest {
		{ message: "Please, Check your provider!" },
		{ message: "Warning: unautorized" },
		{ message: "Application will stay down in 15 mins" },
	}
	for _, test := range testCases {
		Warning(test.message)
	}
}

func Test_InfoMessage(t *testing.T) {
	testCases := [] messageTypeTest {
		{ message: "This Application is starting!" },
		{ message: "Welcome to Appgen" },
		{ message: "This Admin profile!" },
	}
	for _, test := range testCases {
		Info(test.message)
	}
}

func Test_SuccessMessage(t *testing.T) {
	testCases := [] messageTypeTest {
		{ message: "Project Created!" },
		{ message: "User Logged in Successfully!" },
		{ message: "Application Generated!" },
	}
	for _, test := range testCases {
		Success(test.message)
	}
}