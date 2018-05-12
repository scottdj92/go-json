package crud

import (
	"fmt"
	"log"
	"net/http"
)

// GetPerson gets a person by id
func GetPerson(res http.ResponseWriter, req *http.Response) {
	log.Fatal(fmt.Sprintf("logging a request to %v", req))
}
