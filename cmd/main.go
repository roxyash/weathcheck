package main

import (
	"net/http"
	"weathcheck/internal/api"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	appid := viper.GetString("wapi.appid")

	token := viper.GetString("dadata.token")
	secret := viper.GetString("dadata.secret")

	data, err := api.Geocoder("Пресненская набережная 2", token, secret)
	if err != nil {
		logrus.Error(http.StatusInternalServerError, err.Error())
	}

	tempInfo, err := api.GetWeather(data[0].Latitude, data[0].Longitude, appid)

	if err != nil {
		logrus.Error(http.StatusInternalServerError, err.Error())
	}

	logrus.Infof("Temperature now: %v", tempInfo)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
