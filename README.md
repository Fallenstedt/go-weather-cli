# Go Weather CLI

Heyo you can fetch the current weather and check out a simple example of how to create mocks for your go tests
  
![Weather](./weather.gif)

## Local Develoment

1. You will need an API key from [OpenWeather so you can fetch current weather data](https://openweathermap.org/current)
2. Add the API key as a query param in `api.go`
3. Run `go run main.go`

You can also run the test with mocking examples by running `go test ./...`
