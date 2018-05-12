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
	// router.HandleFunc("/people/{id}", crud.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", crud.UpdatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", crud.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

	// adding mock data
	people = append(people, Person{ID: "1", FirstName: "John", LastName: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", FirstName: "Koko", LastName: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", FirstName: "Francis", LastName: "Sunday"})
}

// Person a person struct
type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address an address struct
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person
