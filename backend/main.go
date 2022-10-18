package main

import (
	"fmt"
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

func init() {
	conf = new(config.Config)
	err := conf.Init()
	if err != nil {
		log.Fatalf("Config error: %s\n", err.Error())
	}
	log.Println(conf)
	log.Println(*conf.Server)
}

// Here we start.
func main() {
	s := NewServer(conf.Server.Port)
	err := s.Start()
	if err != nil {
		log.Fatalf("Server running error: %s\n", err.Error())
	}
}
