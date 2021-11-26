package main

// https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
}

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	err = http.ListenAndServe(":"+os.Getenv("PORT"), server.router)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}

func NewServer() (*server, error) {
	srv := server{
		router: mux.NewRouter(),
	}
	return &srv, nil
}
