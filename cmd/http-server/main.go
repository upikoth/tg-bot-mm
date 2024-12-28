package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/app"
	"github.com/upikoth/tg-bot-mm/internal/config"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
)

func main() {
	initCtx := context.Background()
	// Чтение .env файла нужно только при локальной разработке.
	// В других случаях значения переменных окружения уже должны быть установлены.
	// Поэтому ошибку загрузки файла обрабатывать не нужно.
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
	}

	go func() {
		loggerInstance.Info("Запуск приложения")

		if appErr := appInstance.Start(initCtx); !errors.Is(appErr, http.ErrServerClosed) {
			loggerInstance.Fatal(appErr.Error())
		}

		loggerInstance.Info("Приложение перестало принимать новые запросы")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	timeToStopAppInSeconds := 10
	shutdownCtx, shutdownRelease := context.WithTimeout(
		context.Background(),
		time.Duration(timeToStopAppInSeconds)*time.Second,
	)
	defer shutdownRelease()

	if stopErr := appInstance.Stop(shutdownCtx); stopErr != nil {
		loggerInstance.Fatal(fmt.Sprintf("Не удалось корректно остановить сервер, ошибка: %v", stopErr))
	}

	loggerInstance.Info("Приложение остановлено")
}
