package weather

import (
	"fmt"
	"goweather/lib/models"
	"goweather/lib/services"
	"log"
)

type IWeather interface {
	GetWeather() string
}

type Weather struct {
	api services.IAPI
	date services.IDateService
}

func NewWeather(api services.IAPI, date services.IDateService) *Weather {
	return &Weather{
		api,
		date,
	}
}

func (w Weather) GetWeather() string {
	t := new(models.Response)
	err := w.api.GetCurrentWeather(t)

	if err != nil {
		log.Fatal(err)
	}

	temp := fmt.Sprint(w.convertKelvinToFahrenheit(t.Main.Temp)) + "°F"
	today := "Last updated " + w.date.GetCurrentDate()
	return temp + " " + today
}

func (w Weather) convertKelvinToFahrenheit(k float64) uint {
	return uint((k - 273.15) * (9/5) + 32)
}




