package doggo

import (
	"testing"
)

// TODO: Eventually these responses will need to be stubbed

func TestAllBreeds(t *testing.T) {
	tests := []struct {
		expectedStatus string
	}{
		{"success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.AllBreeds()

		actualStatus := client.Response["status"]
		if actualStatus != test.expectedStatus {
			t.Errorf("All Breeds status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response["message"] == nil {
			t.Errorf("All Breeds message was not present")
		}
	}
}
