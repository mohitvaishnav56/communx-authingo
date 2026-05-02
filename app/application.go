package app

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}
type Application struct {
	Config Config
}

func NewConfig(Addr string) Config {
	return Config{
		Addr: Addr,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      nil, //To do Setup chi router here
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting Server on", app.Config.Addr)

	return server.ListenAndServe()
}
