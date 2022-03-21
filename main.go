package main

import (
	"log"
	"net/http"
	"runtime"
)

// Handler struct is a custom handler.
type Handler struct {
	http.Handler
	Name string
}

// Setting configurations.
func configs() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// ServeHTTP is Handler's main function.
func (hand *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	page(rw, r)
}

// Adds root, assets and starts the server.
func runServer(port string) error {
	mux := http.NewServeMux()
	var hand *Handler = &Handler{Name: "Handy"}

	// Adding root.
	mux.Handle("/", hand)
	// Adding CSS files.
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("Start")
	err := http.ListenAndServe(port, mux)
	return err
}

// Here we start.
func main() {
	configs()
	err := runServer(":8080")
	if err != nil {
		panic(err)
	}
}
