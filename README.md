# Doggo

Go client for the [Dog API](https://dog.ceo/dog-api/).

### Usage & Supported Endpoints

All Breeds - `https://dog.ceo/api/breeds/list/all`

  ```
  client := doggo.InitClient()
  client.AllBreeds()
  resp := client.Response
  fmt.Println(resp["message"])
  => map[bouvier:[] collie:[border] dane:[great] eskimo:[] ...
  ```

Random Image - `https://dog.ceo/api/breeds/image/random`

  ```
  client := doggo.InitClient()
  client.RandomImage()
  resp := client.Response
  fmt.Println(resp["message"])
  => https://images.dog.ceo/breeds/ridgeback-rhodesian/n02087394_1161.jpg
  ```

### TODO

* All the endpoints!
* Stub test endpoints (they make real hits right now)

### License

MIT
