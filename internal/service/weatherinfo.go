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
	Apikey string
}

func NewWeatherInfoService(appid, token, secret, apikey string) *WeatherInfo {
	return &WeatherInfo{Appid: appid, Token: token, Secret: secret, Apikey: apikey}
}

func (s *WeatherInfo) GetWeatherInfo(address string) (types.ResponseWeatherInfo, error) {

	data, err := api.GeocoderGraphhopper(address, s.Appid)
	logrus.Infof("%v", err)
	if err != nil {
		logrus.Info(http.StatusInternalServerError, err.Error())
		return types.ResponseWeatherInfo{
			Error: err.Error(),
		}, nil
	}


	weatherInfo, err := api.GetWeather(data[0].Latitude, data[0].Longitude, s.Appid)

	if err != nil {
		logrus.Info(http.StatusInternalServerError, err.Error())
	}

	return types.ResponseWeatherInfo{
		Temperature: fmt.Sprintf("%v", math.Ceil(weatherInfo.TempInfo.Temp)),
		Weather:     weatherInfo.Weather[0].Main,
		Region:      data[0].Region,
		Error:       err.Error(),
	}, nil
}
