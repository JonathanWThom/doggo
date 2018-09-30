package doggo

import (
	"testing"
)

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

func TestSubBreedImages(t *testing.T) {
	tests := []struct {
		breed          string
		subbreed       string
		expectedStatus string
	}{
		{"hound", "afghan", "success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.SubBreedImages(test.breed, test.subbreed)

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("Sub Breed Images status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("Sub Breed Images message was not present")
		}
	}
}

func TestRandomImageBySubBreed(t *testing.T) {
	tests := []struct {
		breed          string
		subbreed       string
		expectedStatus string
	}{
		{"hound", "afghan", "success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.RandomImageBySubBreed(test.breed, test.subbreed)

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("Random Sub Breed Images status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("Random Sub Breed Images message was not present")
		}
	}
}

func TestMultipleImagesBySubBreed(t *testing.T) {
	tests := []struct {
		breed          string
		subbreed       string
		num            int
		expectedStatus string
	}{
		{"hound", "afghan", 3, "success"},
	}

	for _, test := range tests {
		client := InitClient()
		client.MultipleImagesBySubBreed(test.breed, test.subbreed, test.num)

		actualStatus := client.Response.Status
		if actualStatus != test.expectedStatus {
			t.Errorf("Multiple Images by Sub Breed status was incorrect, got: %s, want: %s.",
				actualStatus, test.expectedStatus)
		}

		if client.Response.Message == nil {
			t.Errorf("Multiple Images by Sub Breed message was not present")
		}
	}
}
