package doggo

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Client holds Dog API's base endpoint as well as the response the client receives
// back
type Client struct {
	BaseURL  string
	Response Response
}

// Response holds the status and message returned from each API call
type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

// InitClient initializes a new client
// client := doggo.InitClient()
func InitClient() Client {
	return Client{BaseURL: "https://dog.ceo/api/"}
}

// AllBreeds fetches all breeds
// client.AllBreeds()
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
func (c *Client) RandomImage() error {
	endpoint := "breeds/image/random"

	return getResponse(c, endpoint)
}

// ImagesByBreed fetches all images available for a particular breed
// client.ImagesByBreed("dachshund")
func (c *Client) ImagesByBreed(breed string) error {
	endpoint := getBreedEndpoint(breed)

	return getResponse(c, endpoint)
}

// RandomImageByBreed fetch a random image url for a particular breed
// client.RandomImageByBreed("dachshund")
func (c *Client) RandomImageByBreed(breed string) error {
	endpoint := getRandomBreedEndpoint(breed)

	return getResponse(c, endpoint)
}

// MultipleImagesByBreed fetches a set number of random images from a particular breed
// client.MultipleImagesByBreed("dachshund", 3)
func (c *Client) MultipleImagesByBreed(breed string, num int) error {
	endpoint := getRandomBreedEndpoint(breed) + "/" + strconv.Itoa(num)

	return getResponse(c, endpoint)
}
