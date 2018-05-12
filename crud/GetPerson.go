package crud

import (
	"encoding/json"
	"net/http"

	"../data"
	"github.com/gorilla/mux"
)

// GetPerson gets all people from ../data
func GetPerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range data.People {
		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
		}
	}
}
