package app

import (
	"mishaga/internal/server"
	"mishaga/pkg/config"
	"mishaga/pkg/database"
	"os"

	"github.com/rs/zerolog"
)

type ApplicationConfig struct {
	Server *server.Config   `json:"server"`
	DB     *database.Config `json:"db"`
}

func Run(configPath string) error {
	var appConfig ApplicationConfig

	if err := config.Load(&appConfig, configPath); err != nil {
		return err
	}

	logger := zerolog.New(os.Stdout)

	db := database.NewDatabase(appConfig.DB, &logger)
	if err := db.Init(); err != nil {
		return err
	}
	defer db.Close()

	if err := db.Migrate("./migrations"); err != nil {
		return err
	}

	serv := server.NewServer(appConfig.Server, db, &logger)
	return serv.Run()
}
