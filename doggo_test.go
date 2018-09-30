package doggo

import (
	"testing"
)

// TODO: Eventually these responses will need to be stubbed
// TODO: Tests could maybe be DRYed up with helpers.

func TestAllBreeds(t *testing.T) {
	tests := []struct {
		expectedStatus string
	}{
		{"success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.AllBreeds()

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("All Breeds status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("All Breeds message was not present")
		}
	}
}

func TestRandomImage(t *testing.T) {
	tests := []struct {
		expectedStatus string
	}{
		{"success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.RandomImage()

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("Random Image status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("Random Image message was not present")
		}
	}
}

func TestImagesByBreed(t *testing.T) {
	tests := []struct {
		breed          string
		expectedStatus string
	}{
		{"dachshund", "success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.ImagesByBreed(test.breed)

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("Images By Breed status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("Images By Breed message was not present")
		}
	}
}

func TestRandomImageByBreed(t *testing.T) {
	tests := []struct {
		breed          string
		expectedStatus string
	}{
		{"dachshund", "success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.RandomImageByBreed(test.breed)

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("Random Image By Breed status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("Random Image By Breed message was not present")
		}
	}
}

func TestMultipleImagesByBreed(t *testing.T) {
	tests := []struct {
		breed          string
		num            int
		expectedStatus string
	}{
		{"dachshund", 3, "success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.MultipleImagesByBreed(test.breed, test.num)

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("Multiple Images by Breed status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("Multiple Images by Breed message was not present")
		}
	}
}

func TestSubBreeds(t *testing.T) {
	tests := []struct {
		breed          string
		expectedStatus string
	}{
		{"hound", "success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.SubBreeds(test.breed)

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("Sub Breeds status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("Sub Breeds message was not present")
		}
	}
}
