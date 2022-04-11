package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/store", createEvent).Methods("POST")
	router.HandleFunc("/store", getAllEvents).Methods("GET")
	router.HandleFunc("/store/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/store/{id}", updateEvent).Methods("PUT")
	router.HandleFunc("/store/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
