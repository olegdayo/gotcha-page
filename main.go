package main

import (
	"log"
	"net/http"
)

func runServer(port string) {
	mux := http.NewServeMux()
	var hand *Handler = &Handler{Name: "Handy"}

	mux.Handle("/", hand)
	mux.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("assets"))))

	log.Println("Start")
	http.ListenAndServe(port, mux)
}

func main() {
	runServer(":8080")
}
