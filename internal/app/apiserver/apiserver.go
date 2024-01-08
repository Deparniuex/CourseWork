package apiserver

import (
	"log"
	"os"
	"os/signal"
	"ripProject/internal/config"
	"ripProject/internal/handler"
	"ripProject/internal/httpserver"
	"ripProject/internal/repository/pgrep"
	"ripProject/internal/service"
	"ripProject/internal/storage/postgres"
)

func Run(cfg *config.Config) error {
	db, err := postgres.ConnectDB(
		postgres.WithUser(cfg.DB.User),
		postgres.WithHost(cfg.DB.Host),
		postgres.WithPort(cfg.DB.Port),
		postgres.WithDBName(cfg.DB.DBName),
		postgres.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("Connection to DB error: %s", err.Error())
		return err
	}
	log.Println("Connection success")
	repository := pgrep.New(db, cfg)
	services := service.New(repository, cfg)
	handler := handler.New(services, cfg)
	server := httpserver.NewServer(handler.InitRouter(), httpserver.WithPort(cfg.HTTP.Port), httpserver.WithReadTimeout(cfg.HTTP.Timeout))
	server.Start()
	log.Println("Server started")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}
	return nil
}
