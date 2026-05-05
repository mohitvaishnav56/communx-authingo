package config

import (
	config "authService/config/env"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error){
	config.Load()
	//get config for mysql
	confg := mysql.NewConfig()
	confg.User = config.GetString("DBUSER", "root")
	confg.Passwd = config.GetString("DBPASSWORD", "root")
	confg.Net = "tcp"
	confg.Addr = "127.0.0.1:3306"
	confg.DBName = "users"
	
	db, err := sql.Open("mysql", confg.FormatDSN())
	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(confg.FormatDSN())
	pingErr := db.Ping()

	if pingErr != nil{
		fmt.Println(pingErr)
		return nil, err
	}
	fmt.Println("connected")

	return db, nil
}