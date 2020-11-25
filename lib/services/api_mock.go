package services

import (
	"goweather/lib/models"
)

var (
	API_MOCK_GetCurrentWeather_Response models.Response
	API_MOCK_GetCurrentWeather_Error    error
)

type API_MOCK struct {}

func (a API_MOCK) GetCurrentWeather(res *models.Response) error {
	*res = API_MOCK_GetCurrentWeather_Response
	return API_MOCK_GetCurrentWeather_Error
}


func (a API_MOCK) GivenSuccessGetCurrentWeather(temp float64) {
	API_MOCK_GetCurrentWeather_Response = models.Response{
		Coord: models.Coord{
			Lat: 32.2,
			Lon: 32.2,
		},
		Weather: func() []models.WeatherRes {
			w := make([]models.WeatherRes, 0)
			w = append(w, models.WeatherRes{
				ID: 1,
				Main: "yes",
				Description: "Yes",
				Icon: "Foo",
			})
			return w
		}(),
		Main: models.Main{
			Temp: temp,
			FeelsLike: 72.0,
			TempMin: 72.0,
			TempMax: 72.0,
			Pressure: 1,
			Humidity: 22,
		},
	}
	API_MOCK_GetCurrentWeather_Error = nil

}