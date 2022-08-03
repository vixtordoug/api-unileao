package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []Event

var events = allEvents{
	{ID: 1, Title: "Unilão", Description: "Evento 1"},
	{ID: 2, Title: "Unilão", Description: "Evento 2"},
}

func main() {
	// access log
	log.Println("Starting API...")
	//endpoints
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/events", GetAllEvents).Methods("GET")
	http.ListenAndServe(":"+port, router) // Test heroku

	//Test localhost:8080
	//http.ListenAndServe(":8080", router)

}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Home endpoint accessed")
	fmt.Fprintf(w, "Aplication is running...")
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("GettingAllEvents endpoint accessed")
	w.Header().Set("Content-Type", "Aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

//comandos:
//go mod init
//go mod tidy
//go run main.go
