package crud

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"../data"
)

// DeletePerson deletes a person by id
func DeletePerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range data.People {
		if item.ID == params["id"] {
			data.People = append(data.People[:index], data.People[index+1])
			break
		}
	}

	json.NewEncoder(res).Encode(data.People)
}
