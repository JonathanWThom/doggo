# Doggo

Go client for the [Dog API](https://dog.ceo/dog-api/).

### Usage & Supported Endpoints

**All Breeds** - `https://dog.ceo/api/breeds/list/all`

  ```
  client := doggo.InitClient()
  client.AllBreeds()
  fmt.Println(client.Response.Message)

  => map[bouvier:[] collie:[border] dane:[great] eskimo:[] ...
  ```

**Random Image** - `https://dog.ceo/api/breeds/image/random`

  ```
  client := doggo.InitClient()
  client.RandomImage()
  fmt.Println(client.Response.Message)

  => https://images.dog.ceo/breeds/ridgeback-rhodesian/n02087394_1161.jpg
  ```

**All Images By Breed** - `https://dog.ceo/api/breed/dachshund/images`

  ```
  client := doggo.InitClient()
  client.ImagesByBreed("dachshund")
  fmt.Println(client.Response.Message)

  => [https://images.dog.ceo/breeds/dachshund/Dachshund_rabbit.jpg https://images.dog.ceo/breeds/dachshund/Daschund-2.jpg ...]
  ```

**Random Image By Breed** - `https://dog.ceo/api/breed/dachshund/images/random`

  ```
  client := doggo.InitClient()
  client.RandomImageByBreed("dachshund")
  fmt.Println(client.Response.Message)

  => https://images.dog.ceo/breeds/dachshund/Standard_Wire-hair_Dachshund.jpg
  ```

### TODO

* All the endpoints!
  - `https://dog.ceo/api/breed/hound/images/random/3`
  - `https://dog.ceo/api/breed/hound/list`
  - `https://dog.ceo/api/breed/hound/afgan/images`
  - `https://dog.ceo/api/breed/hound/afghan/images/random`
  - `https://dog.ceo/api/breed/hound/afghan/images/random/3`
* Stub test endpoints (they make real hits right now).
* Tests for `client.Response.Message` could be more specific to the response we want.
* Should a client instance hold all responses it has received, or just the most
recent one?
* Can `ImagesByBreed` and `RandomImageByBreed` be combined somehow?

### License

MIT
