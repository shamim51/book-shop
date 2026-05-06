package main

import (
	"book-shop/cmd/api/router"
	"book-shop/config"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	conf := config.New()
	//mux := http.NewServeMux()
	var logLevel gormlogger.LogLevel
	if conf.DB.Debug {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(fmtDBString, conf.DB.Host, conf.DB.UserName, conf.DB.Password, conf.DB.DBName, conf.DB.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("DB connection start failure")
		return
	}

	r := router.NewRouter(db)
	//http.HandleFunc("/", hello)
	//mux.HandleFunc("/", hello)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Server.Port),
		Handler:      r,
		ReadTimeout:  conf.Server.TimeoutRead,
		WriteTimeout: conf.Server.TimeoutWrite,
		IdleTimeout:  conf.Server.TimeoutIdle,
	}

	println("Starting the server on port -> ", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		println("something went wrong....")
	}
}
