package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type WeatherInfo struct {
	WeatherInfo Main `json:"main"`
}

type Main struct {
	Temperature string `json:"temp"`
}





func GetWeather(lat, lon, appid string) (WeatherInfo, error) {
	logrus.Info(lat, lon, appid)
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s", lat, lon, appid), nil)
	logrus.Info(req.URL)

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

	logrus.Info(resp.StatusCode)

	weatherInfo := WeatherInfo{}

	err = json.Unmarshal(bodyText, &weatherInfo)
	if err != nil {
		return WeatherInfo{}, nil
	}

	return weatherInfo, nil
}
