package main

import (
	"authService/app"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		fmt.Println("error loading env")
	}
	addr := os.Getenv("PORT")
	cfg := app.NewConfig(addr)
	app := app.NewApplication(cfg)
	app.Run()
}
