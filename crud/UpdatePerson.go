package crud

import (
	"fmt"
	"log"
	"net/http"
)

// UpdatePerson updates a person by id
func UpdatePerson(res http.ResponseWriter, req *http.Request) {
	log.Fatal(fmt.Sprintf("logging a reques to %v", req))
}
