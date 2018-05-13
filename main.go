package main

import (
	"log"
	"net/http"

	"./crud"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", crud.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", crud.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", crud.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", crud.DeletePerson).Methods("DELETE")

	log.Print(http.ListenAndServe(":8080", router))
}
