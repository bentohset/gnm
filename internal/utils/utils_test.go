package utils

import (
	"fmt"
	"testing"

	"github.com/bentohset/gnm/internal/mock"
)

func Test_GetInput(t *testing.T) {
	userInput := "test input\n"
	funcDefer, err := mock.MockStdin(t, userInput)
	if err != nil {
		t.Fatal(err)
	}
	defer funcDefer()

	expected := "test input"
	actual := GetInput("Enter input: ")

	if actual != expected {
		t.Errorf("Expected %q, got %q", expected, actual)
	}
}

func Test_ParseBoolean(t *testing.T) {
	testCases := []struct {
		input         string
		expectedValue bool
		expectedError error
	}{
		{"y", true, nil},
		{"n", false, nil},
		{"", true, nil},
		{"invalid", false, fmt.Errorf("invalid input, should be boolean")},
	}

	// Execute and verify
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actualValue, actualError := ParseBoolean(tc.input)

			if actualValue != tc.expectedValue {
				t.Errorf("Expected value: %t, but got: %t", tc.expectedValue, actualValue)
			}

			if (actualError == nil && tc.expectedError != nil) || (actualError != nil && tc.expectedError == nil) || (actualError != nil && tc.expectedError != nil && actualError.Error() != tc.expectedError.Error()) {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedError, actualError)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	err := Execute("", "ls")
	if err != nil {
		t.Error("should return true")
	}

	err = Execute("", "aaa")
	if err == nil {
		t.Error("should return false")
	}
}
