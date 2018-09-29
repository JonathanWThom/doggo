package doggo

import (
	"reflect"
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
		actualStatus := client.Response

		if !reflect.DeepEqual(actualStatus, test.expectedStatus) {
			t.Errorf("All Breeds status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}
	}
}
