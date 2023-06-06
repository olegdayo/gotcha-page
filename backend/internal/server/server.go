package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/offluck/gotcha-page/backend/internal/requesters"
	"github.com/offluck/gotcha-page/backend/internal/sharedData"
	"log"
	"net/http"
	"strings"
)

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

func setRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           1000,
	})
	r.Use(c.Handler)

	r.Get("/", getNetworks)
	r.Get("/{nickname}", getUsers)

	return r
}

func getNetworks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	links, err := json.Marshal(sharedData.GetConfig())
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
				r.URL.Query().Get("networks"),
				"[",
			),
			"]",
		),
		" ",
	)
	fmt.Println(nickname)
	fmt.Println(clients)

	// Container initialization and execution.
	container := requesters.NewRequesterContainer(nickname)
	container.SetUsedLinks(clients...)

	usersInfo := container.GetLinks()

	users, err := json.Marshal(
		struct {
			Users []*requesters.UserInfo `json:"users"`
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
