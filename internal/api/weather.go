package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherInfo struct {
	TempInfo Main      `json:"main"`
	Weather  []Weather `json:"weather"`
	Wind     Wind      `json:"wind"`
}

type Wind struct {
	Deg   float64 `json:"deg"`
	Gust  float64 `json:"gust"`
	Speed float64 `json:"speed"`
}

type Weather struct {
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	Id          float64 `json:"id"`
	Main        string  `json:"main"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

func GetWeather(lat, lon, appid string) (WeatherInfo, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", lat, lon, appid), nil)

	if err != nil {
		return WeatherInfo{}, nil
	}
	resp, err := client.Do(req)
	if err != nil {
		return WeatherInfo{}, nil
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return WeatherInfo{}, nil
	}

	weatherInfo := WeatherInfo{}

	err = json.Unmarshal(bodyText, &weatherInfo)
	if err != nil {
		return WeatherInfo{}, err
	}

	return weatherInfo, nil
}
