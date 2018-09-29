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

### TODO

* All the endpoints!
* Stub test endpoints (they make real hits right now)

### License

MIT
