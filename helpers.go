package doggo

import (
	"encoding/json"
	"net/http"
	"strings"
)

func getResponse(c *Client, endpoint string) error {
	resp, err := http.Get(c.BaseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&c.Response)
}

func getBreedEndpoint(breed string) string {
	return "breed/" + strings.ToLower(breed)
}

func getBreedImageEndpoint(breed string) string {
	return getBreedEndpoint(breed) + "/images"
}

func getRandomBreedEndpoint(breed string) string {
	return getBreedImageEndpoint(breed) + "/random"
}

func getSubBreedImageEndpoint(breed, subbreed string) string {
	return getBreedEndpoint(breed) + "/" + subbreed + "/images"
}

func getRandomSubBreedImageEndpoint(breed, subbreed string) string {
	return getSubBreedImageEndpoint(breed, subbreed) + "/random"
}
