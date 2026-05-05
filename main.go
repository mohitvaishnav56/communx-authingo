package main

import (
	"authService/app"
	dbConfig "authService/config/db"
	envConfig "authService/config/env"
)

func main() {
	envConfig.Load()
	addr := envConfig.GetString("PORT", ":8080")
	cfg := app.NewConfig(addr)
	app := app.NewApplication(cfg)
	dbConfig.SetupDB()
	app.Run()
}
