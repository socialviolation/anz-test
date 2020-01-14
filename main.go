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

type versionResponse struct {
	App []version `json:"myapplication"`
}

type version struct {
	Version     string `json:"version"`
	SHA         string `json:"lastcommitsha"`
	Description string `json:"description"`
}

func getVersion() versionResponse {
	sv := versionResponse{}
	sv.App = append(sv.App, version{
		Version:     os.Getenv("APP_VERSION"),
		SHA:         os.Getenv("SHA"),
		Description: description,
	})
	return sv
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(getVersion())
	if err != nil {
		log.Printf("Error occurred encoding response: %v", err)
	}
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
