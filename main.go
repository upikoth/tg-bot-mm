package main

import (
	"app/config"
	"app/db"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type service struct {
	r      *http.Request
	w      http.ResponseWriter
	bot    *bot.Bot
	db     *db.DB
	update *models.Update
}

func Handler(w http.ResponseWriter, r *http.Request) {
	srv, err := newService(r, w)

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	defer func() {
		cErr := srv.close()

		if cErr != nil {
			log.Println("close error: ", cErr)
			return
		}
	}()

	err = srv.setCommands()

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	err = srv.handleMessages()

	if err != nil {
		log.Println("Error: ", err)
		return
	}

	log.Println("Handler успешно выполнен")
}

func newService(
	r *http.Request,
	w http.ResponseWriter,
) (*service, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	b, err := bot.New(cfg.BotToken)
	if err != nil {
		return nil, err
	}

	ydb, err := db.New(r.Context(), cfg)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	log.Println("Request body: ", string(body))

	var update models.Update
	err = json.Unmarshal(body, &update)
	if err != nil {
		return nil, err
	}

	if update.Message == nil {
		return nil, errors.New("message is empty")
	}

	return &service{
		r:      r,
		w:      w,
		bot:    b,
		db:     ydb,
		update: &update,
	}, nil
}

func (s *service) close() error {
	return s.db.Close(s.r.Context())
}
