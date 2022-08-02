package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// access log
	log.Println("Starting API...")
	//endpoints
	port := os.Getenv("PORT")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	http.ListenAndServe(":"+port, router) // Test heroku

	//Test localhost:8080
	//http.ListenAndServe(":8080", router)

}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Home endpoint accessed")
	fmt.Fprintf(w, "Aplication is running...")
}

//comandos:
//go mod init
//go mod tidy
//go run main.go
