package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Port = ":7000"
)

func serverDynamic(w http.ResponseWriter, r *http.Request) {
	response := "the time is now: " + time.Now().String()
	fmt.Fprintln(w, response)
}
func serverStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {
	http.HandleFunc("/static", serverStatic)
	http.HandleFunc("/dynamic", serverDynamic)
	http.ListenAndServe(Port, nil)
}
