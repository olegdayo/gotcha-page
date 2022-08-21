package main

import (
	"fmt"
	"log"
	"net/http"
)

var conf *Config

// Handler struct is a custom handler.
type Handler struct {
	http.Handler
	Name string
}

func enableCORS(rw http.ResponseWriter) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// ServeHTTP is Handler's main function.
func (hand *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	enableCORS(rw)
	requests(rw, r)
}

// Adds root, assets and starts the server.
func runServer(port string) (err error) {
	mux := http.NewServeMux()
	var hand *Handler = &Handler{Name: "Handy"}

	// Adding root.
	mux.Handle("/", hand)
	// Adding CSS files.
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(conf.AssetsPath))))
	// Adding JS files
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir(conf.ScriptsPath))))

	log.Println("Start")
	err = http.ListenAndServe(port, mux)
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
