package main

import (
	"github.com/spf13/viper"
	"log"
	specialtribble "special-tribble"
	"special-tribble/pkg/handler"
	"special-tribble/pkg/repository"
	"special-tribble/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initilization configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(specialtribble.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
