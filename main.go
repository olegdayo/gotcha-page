package main

import (
	"fmt"
	"net/http"
)

func runServer(port string) {
	var hand *Handler = &Handler{Name: "Handy"}
	http.Handle("/", hand)
	fmt.Println("Start")
	err := http.ListenAndServe(port, hand)
	if err != nil {
		panic(err)
	}
}

func main() {
	runServer(":8080")
}
