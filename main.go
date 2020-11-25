package main

import (
	"fmt"
	"goweather/lib"
	"goweather/lib/services"
)

func main() {
	a := services.NewAPI()
	d := services.NewDateService()
	fmt.Println(weather.NewWeather(a, d).GetWeather())
}