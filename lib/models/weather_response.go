package models


type Response struct {
	Coord   Coord        `json:"coord"`
	Weather []WeatherRes `json:"weather"`
	Main    Main         `json:"main"`
}

type Main struct {
	Temp float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin float64 `json:"temp_min"`
	TempMax float64 `json:"temp_max"`
	Pressure uint `json:"pressure"`
	Humidity uint `json:"humidity"`
}

type WeatherRes struct {
	ID uint `json:"id"'`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json"icon"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"long"`
}

//{
//  "visibility": 16093,
//  "wind": {
//    "speed": 1.5,
//    "deg": 350
//  },
//  "clouds": {
//    "all": 1
//  },
//  "dt": 1560350645,
//  "sys": {
//    "type": 1,
//    "id": 5122,
//    "message": 0.0139,
//    "country": "US",
//    "sunrise": 1560343627,
//    "sunset": 1560396563
//  },
//  "timezone": -25200,
//  "id": 420006353,
//  "name": "Mountain View",
//  "cod": 200
//  }