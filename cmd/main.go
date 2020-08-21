package main

import (
	"github.com/cobbinma/kanye-west-api/cmd/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	client := http.DefaultClient
	r.HandleFunc("/kanye", handlers.Kanye(client))
	log.Println("serving quotes...")
	log.Fatal(http.ListenAndServe(":2024", r))
}
