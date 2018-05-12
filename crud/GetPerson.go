package crud

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../data"
)

// GetPerson gets all people from ../data
func GetPerson(req http.ResponseWriter, res *http.Response) {
	var allPeople = json.NewEncoder(req).Encode(data.People)
	log.Printf(fmt.Sprintf("logging a request to %v", allPeople))
}
