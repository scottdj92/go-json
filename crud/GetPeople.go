package crud

import (
	"log"
	"net/http"
)

// GetPeople gets all people from within the api
func GetPeople(res http.ResponseWriter, req *http.Request) {
	log.Fatal("logging a reqest to /people")
}
