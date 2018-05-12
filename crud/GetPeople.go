package crud

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetPeople gets all people from within the api
func GetPeople(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(req)
	log.Printf("logging a reqest to /people with %v", req.Body)
}
