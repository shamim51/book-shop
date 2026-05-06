package main

import (
	"book-shop/config"
	"fmt"
	"io"
	"net/http"
)

func main() {
	conf := config.New()
	mux := http.NewServeMux()

	//http.HandleFunc("/", hello)
	mux.HandleFunc("/", hello)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Server.Port),
		Handler:      mux,
		ReadTimeout:  conf.Server.TimeoutRead,
		WriteTimeout: conf.Server.TimeoutWrite,
		IdleTimeout:  conf.Server.TimeoutIdle,
	}

	println("Starting the server on port -> 8080....")
	err := server.ListenAndServe()
	/*
		err := http.ListenAndServe(":8080", mux)

	*/
	if err != nil {
		println("something went wrong....")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	//time.Sleep(3 * time.Second)
	io.WriteString(w, "Hello world!")
}
