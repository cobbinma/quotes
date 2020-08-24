package main

import (
	"github.com/cobbinma/kanye-west-api/cmd/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	h := handlers.NewHandlers(handlers.WithCustomClient(&http.Client{
		Timeout: time.Duration(time.Second * 10),
	}))
	r.HandleFunc("/kanye", h.Kanye)
	log.Println("serving quotes...")
	log.Fatal(http.ListenAndServe(":2024", r))
}
