package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	appName     = "myapplication"
	description = "pre-interview technical test"
	port        = ":8080"
)

type version struct {
	Version     string `json:"version"`
	SHA         string `json:"lastcommitsha"`
	Description string `json:"description"`
}

func getVersion() map[string][]version {
	sv := make(map[string][]version)
	sv[appName] = append(sv[appName], version{
		Version:     os.Getenv("VERSION"),
		SHA:         os.Getenv("SHA"),
		Description: description,
	})
	return sv
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("%+v", r)
	log.Println("called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getVersion())
}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/version", versionHandler).Methods("GET")

	return r
}

func main() {
	router := setupRoutes()
	log.Printf("%s listening on port %s...", appName, port)

	log.Fatal(http.ListenAndServe(port, router))
}
