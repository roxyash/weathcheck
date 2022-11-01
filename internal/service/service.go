package service

import "weathcheck/internal/types"

type WeatherInfoService interface {
	GetWeatherInfo(address string) (types.ResponseWeatherInfo, error)
}

type Service struct {
	WeatherInfoService
}

func NewService(appid, token, secret, apikey string) *Service {
	return &Service{
		WeatherInfoService:  NewWeatherInfoService(appid, token, secret, apikey),
	}
}
