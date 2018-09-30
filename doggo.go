package doggo

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Client holds Dog API's base endpoint as well as the response the client receives
// back
type Client struct {
	BaseURL  string
	Response map[string]interface{} // Could this be split up into status/message in a nicer way?
}

// InitClient initializes a new client
// client := doggo.InitClient()
func InitClient() Client {
	return Client{BaseURL: "https://dog.ceo/api/"}
}

// AllBreeds fetches all breeds
// client.AllBreeds()
// resp := client.Response
func (c *Client) AllBreeds() error {
	endpoint := "breeds/list/all"
	resp, err := http.Get(c.BaseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&c.Response)
}

// RandomImage fetches a random image url
// client.RandomImage()
// resp := client.Response
func (c *Client) RandomImage() error {
	endpoint := "breeds/image/random"

	return getResponse(c, endpoint)
}

// ImagesByBreed fetches all images available for a particular breed
// client.ImagesByBreed("dachshund")
// resp := client.Response
func (c *Client) ImagesByBreed(breed string) error {
	endpoint := "breed/" + strings.ToLower(breed) + "/images"

	return getResponse(c, endpoint)
}

func getResponse(c *Client, endpoint string) error {
	resp, err := http.Get(c.BaseURL + endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&c.Response)
}
