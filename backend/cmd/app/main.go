package main

import (
	"log"
	"runtime"

	"github.com/olegdayo/gotcha-page/backend/internal/server"
	"github.com/olegdayo/gotcha-page/backend/internal/sharedData"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := sharedData.InitConfig("config.yaml")
	if err != nil {
		log.Fatalf("Config error: %s\n", err.Error())
	}

	log.Printf("Config: %+v\n", sharedData.GetConfig())
}

// Here we start.
func main() {
	s := server.NewServer(sharedData.GetConfig().Server.Port)
	err := s.Start()
	if err != nil {
		log.Fatalf("Server running error: %s\n", err.Error())
	}
}
