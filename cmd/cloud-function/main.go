package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/app"
	"github.com/upikoth/tg-bot-mm/internal/config"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	_ = godotenv.Load()
	loggerInstance := logger.New()

	cfg, err := config.New()
	if err != nil {
		loggerInstance.Fatal(err.Error())
		return
	}

	if cfg == nil {
		loggerInstance.Fatal(errors.New("Некорректная инициализация конфига приложения").Error())
		return
	}

	loggerInstance.SetPrettyOutputToConsole()

	appInstance, err := app.New(cfg, loggerInstance)
	if err != nil {
		loggerInstance.Fatal(err.Error())
		return
	}

	if appInstance != nil {
		dbErr := appInstance.ConnectToRepositories()

		if dbErr != nil {
			loggerInstance.Fatal(dbErr.Error())
			return
		}

		appInstance.MainHandler(w, r)
	}
}
