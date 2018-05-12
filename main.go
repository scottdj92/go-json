package main

import (
	"log"
	"net/http"

	"./crud"
	"./data"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", crud.GetPeople).Methods("GET")
	// router.HandleFunc("/people/{id}", crud.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", crud.UpdatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", crud.DeletePerson).Methods("DELETE")

	// adding mock data
	data.People = append(data.People, data.Person{ID: "1", FirstName: "John", LastName: "Doe", Address: &data.Address{City: "City X", State: "State X"}})
	data.People = append(data.People, data.Person{ID: "2", FirstName: "Koko", LastName: "Doe", Address: &data.Address{City: "City Z", State: "State Y"}})
	data.People = append(data.People, data.Person{ID: "3", FirstName: "Francis", LastName: "Sunday"})

	log.Fatal(http.ListenAndServe(":8080", router))
}
