package crud

import (
	"encoding/json"
	"net/http"

	"../data"
)

// GetPeople gets all people from within the api
func GetPeople(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(data.People)
}
