package service

import (
	"fmt"
	"math"
	"net/http"
	"weathcheck/internal/api"
	"weathcheck/internal/types"

	"github.com/sirupsen/logrus"
)

type WeatherInfo struct {
	Appid  string
	Token  string
	Secret string
}

func NewWeatherInfoService(appid, token, secret string) *WeatherInfo {
	return &WeatherInfo{Appid: appid, Token: token, Secret: secret}
}

func (s *WeatherInfo) GetWeatherInfo(address string) (types.ResponseWeatherInfo, error) {

	data, err := api.Geocoder(address, s.Token, s.Secret)
	if err != nil {
		logrus.Error(http.StatusInternalServerError, err.Error())
	}

	weatherInfo, err := api.GetWeather(data[0].Latitude, data[0].Longitude, s.Appid)

	if err != nil {
		logrus.Error(http.StatusInternalServerError, err.Error())
	}

	return types.ResponseWeatherInfo{
		Temperature: fmt.Sprintf("%v", math.Ceil(weatherInfo.TempInfo.Temp)),
		Weather:     weatherInfo.Weather[0].Main,
		Region:      data[0].Region,
	}, nil
}
