package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}