package main

import (
	"fmt"
	"log"
	"net/http"
	"observer/pkg/env"
	"os"
)

func runServer(mux *http.ServeMux) {
	env.GetEnv()
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	log.Printf("Starting server on %s", port)
	err := http.ListenAndServe(port, mux)

	log.Fatal(err)
}

func main() {
	mux := http.NewServeMux()
	SetHandlers(mux)
	runServer(mux)
}
