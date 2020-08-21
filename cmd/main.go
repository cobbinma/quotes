package main

import (
	"github.com/cobbinma/kanye-west-api/cmd/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	h := handlers.NewHandlers(http.DefaultClient)
	r.HandleFunc("/kanye", h.Kanye)
	log.Println("serving quotes...")
	log.Fatal(http.ListenAndServe(":2024", r))
}
