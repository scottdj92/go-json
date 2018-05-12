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
	router.HandleFunc("/people/{id}", crud.UpdatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", crud.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
