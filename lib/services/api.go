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

type API struct{}

func NewAPI() *API {
	return &API{}
}

func (a API) GetCurrentWeather(res *models.Response) error {
	d:= "https://api.openweathermap.org/data/2.5/weather?lat=45.4889&lon=122.8014&appid=1ed583db37628eeebe28831cb486ee00"
	resp, err := http.Get(d)

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
