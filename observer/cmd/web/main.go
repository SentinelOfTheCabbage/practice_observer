package main

import (
	"log"
	"net/http"
)

func runServer(mux *http.ServeMux) {
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func main() {
	mux := http.NewServeMux()
	SetHandlers(mux)
	runServer(mux)
}
