package http

import (
	"encoding/json"
	"log"
	"net/http"
	"webserver-http/domain"
)

const (
	SERVER_ADDR = ":9090"
)

func ServerRun() {
	http.HandleFunc("/", handleNotFound)

	log.Println("[SERVER] starting server on address", SERVER_ADDR)
	http.ListenAndServe(SERVER_ADDR, nil)
}

func handleNotFound(w http.ResponseWriter, _ *http.Request) {
	result := domain.Result{Error: "resource is not found", Code: "404"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	err := json.NewEncoder(w).Encode(result)
	if nil != err {
		log.Println("[HTTP] failed to process request", err)
	}
}
func handleMethodNotAllowed(w http.ResponseWriter, _ *http.Request) {
	result := domain.Result{Error: "method is not allowed", Code: "405"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	err := json.NewEncoder(w).Encode(result)
	if nil != err {
		log.Println("[HTTP] failed to process request", err)
	}
}

func handleGeneric(w http.ResponseWriter, _ *http.Request) {
	result := domain.Result{Error: "failed to process request", Code: "500"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	err := json.NewEncoder(w).Encode(result)
	if nil != err {
		log.Println("[HTTP] failed to process request", err)
	}
}
