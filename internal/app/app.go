package app

import (
	"github.com/6a6ydoping/online-shop/internal/config"
	"github.com/6a6ydoping/online-shop/internal/handler"
	"github.com/6a6ydoping/online-shop/internal/repository/pg"
	"github.com/6a6ydoping/online-shop/internal/service"
	"github.com/6a6ydoping/online-shop/pkg/httpserver"
	"log"
	"os"
	"os/signal"
)

func Run(cfg *config.Config) error {
	db, err := pg.New(
		pg.WithHost(cfg.DB.Host),
		pg.WithPort(cfg.DB.Port),
		pg.WithDBName(cfg.DB.DBName),
		pg.WithUsername(cfg.DB.Username),
		pg.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	log.Println("connection success")

	srvs := service.New(db, cfg)
	hndlr := handler.New(srvs)
	server := httpserver.New(
		hndlr.InitRouter(cfg.Router),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	log.Println("server started")
	server.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
