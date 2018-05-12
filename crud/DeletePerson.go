package crud

import (
	"fmt"
	"log"
	"net/http"
)

// DeletePerson deletes a person by id
func DeletePerson(res http.ResponseWriter, req *http.Response) {
	log.Fatal(fmt.Sprintf("logging a request to /people/%v", req))
}
