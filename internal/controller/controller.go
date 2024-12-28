package controller

import (
	"context"
	"net/http"

	"github.com/upikoth/tg-bot-mm/internal/config"
	controllerhttp "github.com/upikoth/tg-bot-mm/internal/controller/http"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/services"
)

type Controller struct {
	http   *controllerhttp.HTTP
	logger logger.Logger
	config *config.Config
}

func New(
	config *config.Config,
	logger logger.Logger,
	service *services.Services,
) (*Controller, error) {
	httpInstance, err := controllerhttp.New(config, logger, service)

	if err != nil {
		return nil, err
	}

	return &Controller{
		http:   httpInstance,
		logger: logger,
		config: config,
	}, nil
}

func (c *Controller) Start() error {
	return c.http.Start()
}

func (c *Controller) Stop(ctx context.Context) error {
	return c.http.Stop(ctx)
}

func (c *Controller) MainHandler(w http.ResponseWriter, r *http.Request) {
	c.http.MainHandler(w, r)
}
