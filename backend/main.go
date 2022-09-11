package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gotchaPage/config"
	"log"
	"net/http"
)

var conf *config.Config

type Server struct {
	http.Server
}

func NewServer(port uint16) (s *Server) {
	s = new(Server)
	s.Addr = fmt.Sprintf(":%d", port)
	s.Handler = setRouter()
	return
}

func (s *Server) Start() error {
	log.Println("Starting the server")
	return s.ListenAndServe()
}

func setRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", getNetworks)
	r.Get("/{nickname}", getUsers)
	return r
}

func init() {
	conf = new(config.Config)
	err := conf.Init()
	if err != nil {
		panic(err)
	}
	log.Println(conf)
	log.Println(*conf.Server)
}

// Here we start.
func main() {
	s := NewServer(conf.Server.Port)
	err := s.Start()
	if err != nil {
		panic(err)
	}
}
