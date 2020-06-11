package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "up", "version": "0.0.1", "environment": "development"}`))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", get).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
