package weather

import (
	"goweather/lib/services"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// Given successful api call
	api := new(services.API_MOCK)
	api.GivenSuccessGetCurrentWeather(274.05)

	// Given specific date
	date := new(services.DATE_SERVICE_MOCK)
	date.GivenDateServiceReturns("Saturday 23:21 Nov 21, 2020")

	w := NewWeather(api, date)
	if r := w.GetWeather(); r != "32°F Last updated Saturday 23:21 Nov 21, 2020" {
		t.Errorf("Expected result: 32°F Last updated Saturday 23:21 Nov 21, 2020, but got: %v", r)
	}
}
