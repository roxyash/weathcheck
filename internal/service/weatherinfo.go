package service

import (
	"fmt"
	"math"
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

	data, err := api.GeocoderGraphhopper(address, s.Apikey)
	if err != nil {
		return types.ResponseWeatherInfo{}, err
	}

	weatherInfo, err := api.GetWeather(data.Latitude, data.Longitude, s.Appid)
	logrus.Infof("%v %v", weatherInfo, err)
	if err != nil {
		return types.ResponseWeatherInfo{}, err
	}

	responseWeatherInfo := types.ResponseWeatherInfo{Temperature: fmt.Sprintf("%v", math.Ceil(weatherInfo.TempInfo.Temp)),
		Weather: weatherInfo.Weather[0].Main,
		Region:  data.Region,
	}

	return responseWeatherInfo, nil
}
