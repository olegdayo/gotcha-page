package main

import (
	"fmt"
	"net/http"
)

func runServer(port string) {
	var hand *Handler = &Handler{Name: "Handy"}
	http.Handle("/", hand)
	fmt.Println("Start")
	http.ListenAndServe(port, hand)
}

func main() {
	runServer(":8080")
}
