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

// Here we start.
func main() {
	conf = new(Config)
	err := conf.Init()
	if err != nil {
		panic(err)
	}

	log.Println(conf.AssetsPath)
	log.Println(conf.ScriptsPath)
	log.Println(conf.Port)

	err = runServer(fmt.Sprintf(":%d", conf.Port))
	if err != nil {
		panic(err)
	}
}
