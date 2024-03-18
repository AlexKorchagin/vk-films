package main

import (
	"log"

	"github.com/AlexKorchagin/vk-films/cmd"
	"github.com/AlexKorchagin/vk-films/internal/handlers"
	"github.com/AlexKorchagin/vk-films/internal/repo"
	"github.com/spf13/viper"
)

func main() {

	err := initConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repo, err := repo.NewRepository()

	if err != nil {
		log.Fatalf("unable to initialize repository:%s", err.Error())
	}

	handlers.InitHandlers(repo)

	log.Printf("starting server http://localhost:%s", viper.GetString("port"))
	srv := new(cmd.Server)
	if err := srv.Run(viper.GetString("port"), nil); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.SetConfigFile(".yaml")
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
