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
	endpoint := getBreedImageEndpoint(breed)

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

// SubBreeds fetches a list of sub-breeds for a given breed
// client.SubBreeds("hound")
func (c *Client) SubBreeds(breed string) error {
	endpoint := getBreedEndpoint(breed) + "/list"

	return getResponse(c, endpoint)
}

// SubBreedImages fetches all images for a given sub-breed
// client.SubBreedImages("hound", "afghan")
func (c *Client) SubBreedImages(breed, subbreed string) error {
	endpoint := getSubBreedImageEndpoint(breed, subbreed)

	return getResponse(c, endpoint)
}

// RandomImageBySubBreed fetches a random image for a given sub-breed
// client.RandomImageBySubBreed("hound", "afghan")
func (c *Client) RandomImageBySubBreed(breed, subbreed string) error {
	endpoint := getRandomSubBreedImageEndpoint(breed, subbreed)

	return getResponse(c, endpoint)
}

// MultipleImagesBySubBreed fetches a set number of random images from a particular sub-breed
// client.MultipleImagesBySubBreed("hound", "afghan", 3)
func (c *Client) MultipleImagesBySubBreed(breed, subbreed string, num int) error {
	endpoint := getRandomSubBreedImageEndpoint(breed, subbreed) + "/" + strconv.Itoa(num)

	return getResponse(c, endpoint)
}
