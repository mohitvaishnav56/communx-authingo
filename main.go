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
	db, err := dbConfig.SetupDB()
	if err != nil {
		panic(err)
	}	
	app := app.NewApplication(cfg, db)
	app.Run()
}
