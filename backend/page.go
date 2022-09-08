package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strings"
)

func getNetworks(w http.ResponseWriter, r *http.Request) {
	links, err := json.Marshal(conf)
	if err != nil {
		log.Fatalf("Marshal error: %s\n", err.Error())
		return
	}

	_, err = w.Write(links)
	if err != nil {
		log.Fatalf("Write error: %s\n", err.Error())
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	nickname := chi.URLParam(r, "nickname")
	clients := strings.Split(
		strings.TrimRight(
			strings.TrimLeft(
				r.URL.Query().Get("clients"),
				"[",
			),
			"]",
		),
		" ",
	)
	fmt.Println(nickname)
	fmt.Println(clients)

	// Container initialization and execution.
	container := NewRequesterContainer(nickname)
	container.SetUsedLinks(clients...)

	usersInfo := container.GetLinks()

	users, err := json.Marshal(
		struct {
			Users []*UserInfo `json:"users"`
		}{
			Users: usersInfo,
		},
	)
	if err != nil {
		log.Fatalf("Marshal error: %s\n", err.Error())
		return
	}

	_, err = w.Write(users)
	if err != nil {
		log.Fatalf("Write error: %s\n", err.Error())
	}
}
