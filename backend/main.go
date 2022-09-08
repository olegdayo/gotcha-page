package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

var conf *Config

func setRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", getNetworks)
	r.Get("/{nickname}", getUsers)
	return r
}

// Adds root, assets and starts the server.
func runServer(port string) (err error) {
	log.Println("Start")
	r := setRouter()
	err = http.ListenAndServe(port, r)
	return err
}

func init() {
	conf = new(Config)
	err := conf.Init()
	if err != nil {
		panic(err)
	}
	log.Println(conf)
	log.Println(*conf.Server)
}

// Here we start.
func main() {
	err := runServer(fmt.Sprintf(":%d", conf.Server.Port))
	if err != nil {
		panic(err)
	}
}
