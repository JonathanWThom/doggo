package doggo

import (
	"encoding/json"
	"net/http"
)

// Client holds Dog API's base endpoint as well as the response the client receives
// back
type Client struct {
	BaseURL  string
	Response map[string]interface{}
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

// RandomImage fetches all breeds
// client.RandomImage()
// resp := client.Response
func (c *Client) RandomImage() error {
	endpoint := "breeds/image/random"

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
