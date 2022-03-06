package main

import (
	"fmt"
	"net/http"
)

func runServer(port string) {
	mux := http.NewServeMux()
	var hand *Handler = &Handler{Name: "Handy"}

	mux.Handle("/", hand)
	mux.Handle("/assets/", http.StripPrefix("/assets/",
		http.FileServer(http.Dir("assets"))))

	fmt.Println("Start")
	err := http.ListenAndServe(port, mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	runServer(":8080")
}
