package app

import (
	"authService/controllers"
	db "authService/db/repositories"
	"authService/router"
	"authService/services"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr    string
	storage db.Storage
}
type Application struct {
	Config Config
	db     *sql.DB
}

func NewConfig(Addr string) Config {
	return Config{
		Addr: Addr,
	}
}

func NewApplication(cfg Config, _db *sql.DB) *Application {
	return &Application{
		Config: cfg,
		db: _db,
	}
}

func (app *Application) Run() error {
	ur := db.NewUserRepository(app.db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	urouter := router.NewUserRouter(*uc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(urouter), //To do Setup chi router here
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting Server on", app.Config.Addr)

	return server.ListenAndServe()
}
