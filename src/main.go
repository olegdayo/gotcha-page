package main

import (
	"gotchaPage/src/config"
	"log"
	"net/http"
)

var conf *config.Config

// Handler struct is a custom handler.
type Handler struct {
	http.Handler
	Name string
}

// ServeHTTP is Handler's main function.
func (hand *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	page(rw, r)
}

// Adds root, assets and starts the server.
func runServer() (err error) {
	mux := http.NewServeMux()
	var hand *Handler = &Handler{Name: "Handy"}

	// Adding root.
	mux.Handle("/", hand)

	// Adding CSS files.
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(conf.Paths.Assets))))

	// Adding JS files
	mux.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir(conf.Paths.Scripts))))

	log.Println("Start")
	err = http.ListenAndServe(":"+conf.Port, mux)
	return err
}

// Here we start.
func main() {
	var err error
	conf, err = config.Init()
	if err != nil {
		panic(err)
	}

	err = runServer()
	if err != nil {
		panic(err)
	}
}
