package crud

import (
	"fmt"
	"log"
	"net/http"
)

// UpdatePerson updates a person by id
func UpdatePerson(req http.ResponseWriter, res *http.Request) {
	log.Fatal(fmt.Sprintf("logging a reques to %v", res))
}
