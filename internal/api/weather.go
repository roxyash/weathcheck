package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type WeatherInfo struct {
	Temperature string `json:"temp"`
}

func GetWeather(lat, lon, part, appid string) (WeatherInfo, error) {
	logrus.Info(lat, lon, part, appid)
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&exclude=%s&appid=%s", lat, lon, part, appid), nil)
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
