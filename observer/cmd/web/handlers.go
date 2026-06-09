package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(response http.ResponseWriter, request *http.Request) {
	log.Printf("`%s` request to `%s` was received.", request.Method, request.Host)

	page := request.URL.Query().Get("page")
	if page == "" {
		fmt.Fprintf(response, "Hi mate! It's main page afaik.")
		return
	}

	id, err := strconv.Atoi(page)
	if err != nil || id < 0 {
		http.NotFound(response, request)
		return
	}

	fmt.Fprintf(response, "Hi dawg. It's page #%d", id)
}

func SetHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", home)
}
