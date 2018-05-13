package crud

import (
	"encoding/json"
	"net/http"

	"../data"
	"github.com/gorilla/mux"
)

// CreatePerson creates a person within the slice/db
func CreatePerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person data.Person
	_ = json.NewDecoder(req.Body).Decode(&person)

	person.ID = params["id"]
	data.People = append(data.People, person)
	json.NewEncoder(res).Encode(person)
}
