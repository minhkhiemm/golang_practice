package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", home).Methods("GET")
	return r
}

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Your request was processed.")
}

func main() {
	log.Print("Starting the service...")
	router := handlers.Router()
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}
