package services

import (
	"encoding/json"
	"errors"
	"goweather/lib/models"
	"net/http"
)

type IAPI interface {
	GetCurrentWeather(res *models.Response) error
}

type API struct {}

func NewAPI() *API {
	return &API{}
}

func (a API) GetCurrentWeather(res *models.Response) error {
	resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=45.4889&lon=122.8014&appid=d5ffe54323633376b81fa997313518cd")

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("failed to fetch weather")
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(res)
}
