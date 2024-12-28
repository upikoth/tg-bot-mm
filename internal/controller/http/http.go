package http

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/config"
	"github.com/upikoth/tg-bot-mm/internal/controller/http/handler"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/services"
)

type HTTP struct {
	logger    logger.Logger
	appServer *http.Server
	handler   *handler.Handler
}

func New(
	config *config.Config,
	loggerInstance logger.Logger,
	services *services.Services,
) (*HTTP, error) {
	handlerInstance := handler.New(
		loggerInstance,
		config,
		services,
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerInstance.MainHandler)

	appServer := &http.Server{
		Addr:              ":" + config.Port,
		ReadHeaderTimeout: time.Minute,
		Handler:           mux,
	}

	return &HTTP{
		logger:    loggerInstance,
		appServer: appServer,
		handler:   handlerInstance,
	}, nil
}

func (h *HTTP) Start() error {
	return errors.WithStack(h.appServer.ListenAndServe())
}

func (h *HTTP) Stop(ctx context.Context) error {
	return errors.WithStack(h.appServer.Shutdown(ctx))
}

func (h *HTTP) MainHandler(w http.ResponseWriter, r *http.Request) {
	h.handler.MainHandler(w, r)
}
