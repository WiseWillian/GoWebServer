package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

func getValue(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/roulette", getValue).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}