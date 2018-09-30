# Doggo

Go client for the [Dog API](https://dog.ceo/dog-api/).

### Usage & Supported Endpoints

**Installation**

`go get -u github.com/JonathanWThom/doggo`

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

**Multiple Images By Breed** - `https://dog.ceo/api/breed/dachshund/images/random/3`

  ```
  client := doggo.InitClient()
  client.MultipleImagesByBreed("dachshund", 3)
  fmt.Println(client.Response.Message)

  => [https://images.dog.ceo/breeds/dachshund/Stretched_Dachshund.jpg https://images.dog.ceo/breeds/dachshund/dachshund-2033796_640.jpg https://images.dog.ceo/breeds/dachshund/dachshund-6.jpg]
  ```

**Sub Breeds** - `https://dog.ceo/api/breed/hound/list`

  ```
  client := doggo.InitClient()
  client.SubBreeds("hound")
  fmt.Println(client.Response.Message)

  => [afghan basset blood english ibizan walker]
  ```

**Sub Breed Images** - `https://dog.ceo/api/breed/hound/afghan/images`

  ```
  client := doggo.InitClient()
  client.SubBreedImages("hound", "afghan")
  fmt.Println(client.Response.Message)

  => [https://images.dog.ceo/breeds/hound-afghan/n02088094_1003.jpg https://images.dog.ceo/breeds/hound-afghan/n02088094_1007.jpg ...]
  ```

**Random Sub Breed Image** - `https://dog.ceo/api/breed/hound/afghan/images/random`

  ```
  client := doggo.InitClient()
  client.RandomImageBySubBreed("hound", "afghan")
  fmt.Println(client.Response.Message)

  => https://images.dog.ceo/breeds/hound-afghan/n02088094_7131.jpg
  ```

**Multiple Images by Sub Breed** - `https://dog.ceo/api/breed/hound/afghan/images/random/3`

  ```
  client := doggo.InitClient()
  client.MultipleImagesBySubBreed("hound", "afghan", 3)
  fmt.Println(client.Response.Message)

  => [https://images.dog.ceo/breeds/hound-afghan/n02088094_4635.jpg https://images.dog.ceo/breeds/hound-afghan/n02088094_6690.jpg https://images.dog.ceo/breeds/hound-afghan/n02088094_988.jpg]
  ```

### Todo & Design Questions

* Stub test endpoints (they make real hits right now).
* DRY up tests with helpers, potentially.
* Tests for `client.Response.Message` could be more specific to the response we want.
* Should a client instance hold all responses it has received, or just the most
recent one?
* Should breed and sub-breed live on `Client`?

### License

MIT
