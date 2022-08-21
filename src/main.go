package main

import (
	"gotchaPage/src/config"
	"log"
	"net/http"
)

<<<<<<< HEAD
// Paths.
const (
	// To .css files.
	assetsPath string = "assets"
	// To .js files.
	scriptsPath string = "scripts"
)

var port string
=======
var conf *config.Config
>>>>>>> master

// Handler struct is a custom handler.
type Handler struct {
	http.Handler
	Name string
}

<<<<<<< HEAD
// Setting configurations.
func configs() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	port = os.Getenv("PORT")
}

func enableCORS(rw http.ResponseWriter) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

=======
>>>>>>> master
// ServeHTTP is Handler's main function.
func (hand *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	enableCORS(rw)
	requests(rw, r)
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
