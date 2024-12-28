package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/upikoth/tg-bot-mm/internal/config"
	"github.com/upikoth/tg-bot-mm/internal/controller"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/repositories"
	"github.com/upikoth/tg-bot-mm/internal/services"
)

type App struct {
	config       *config.Config
	logger       logger.Logger
	repositories *repositories.Repository
	services     *services.Services
	controller   *controller.Controller
}

func New(
	cfg *config.Config,
	log logger.Logger,
) (*App, error) {
	repositoriesInstance, err := repositories.New(log, cfg)

	if err != nil {
		log.Error(fmt.Sprintf("Ошибка при инициализации repositories: %s", err))
		return nil, err
	}

	servicesInstance, err := services.New(
		log,
		cfg,
		repositoriesInstance,
	)

	if err != nil {
		log.Error(fmt.Sprintf("Ошибка при инициализации services: %s", err))
		return nil, err
	}

	controllerInstance, err := controller.New(cfg, log, servicesInstance)

	if err != nil {
		log.Error(fmt.Sprintf("Ошибка при инициализации controller: %s", err))
		return nil, err
	}

	return &App{
		config:       cfg,
		logger:       log,
		repositories: repositoriesInstance,
		services:     servicesInstance,
		controller:   controllerInstance,
	}, nil
}

func (s *App) Start(_ context.Context) error {
	err := s.repositories.Connect()

	if err != nil {
		return err
	}

	s.logger.Debug("Подключение к repositories прошло без ошибок")

	return s.controller.Start()
}

func (s *App) ConnectToRepositories() error {
	return s.repositories.Connect()
}

func (s *App) Stop(ctx context.Context) error {
	err := s.repositories.Disconnect()

	if err != nil {
		return err
	}

	s.logger.Debug("Отключение от repositories прошло без ошибок")

	return s.controller.Stop(ctx)
}

func (s *App) MainHandler(w http.ResponseWriter, r *http.Request) {
	s.controller.MainHandler(w, r)
}
