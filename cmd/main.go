package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"weathcheck/internal/handler"
	"weathcheck/internal/service"
	"weathcheck/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title Weatherchecker
// @version 1.0
// @description API server for weatherchecker Application 
// @host localhost:8000
// @BasePath /
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	appid := viper.GetString("wapi.appid")

	token := viper.GetString("dadata.token")
	secret := viper.GetString("dadata.secret")

	apikey := viper.GetString("graphhopper.apikey")


	services := service.NewService(appid, token, secret, apikey)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error ocered while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp ShuttingDown")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down :%s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("dev")
	return viper.ReadInConfig()
}
