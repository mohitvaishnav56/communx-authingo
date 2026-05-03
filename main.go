package main

import (
	"authService/app"
	config "authService/config/env"
)

func main() {
	config.Load()
	addr := config.GetString("PORT", ":8080")
	cfg := app.NewConfig(addr)
	app := app.NewApplication(cfg)
	app.Run()
}
