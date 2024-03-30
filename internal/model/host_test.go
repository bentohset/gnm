package model

import (
	"reflect"
	"testing"
)

func Test_NewHost(t *testing.T) {
	// Create a new host using the NewHost function
	expectedHost := Host{
		ID:             1,
		Alias:          "TestAlias",
		Description:    "TestDescription",
		User:           "Testuser",
		HostName:       "TestHost",
		PrivateKeyPath: "/path/to/private/key",
	}

	newHost := NewHost(expectedHost.ID, expectedHost.Alias, expectedHost.User, expectedHost.HostName, expectedHost.Description, expectedHost.PrivateKeyPath)

	// Check if the new host matches the expected host
	if !reflect.DeepEqual(newHost, expectedHost) {
		t.Errorf("NewHost function did not create the expected host. Expected: %v, Got: %v", expectedHost, newHost)
	}
}
