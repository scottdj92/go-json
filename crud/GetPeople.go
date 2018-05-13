package crud

import (
	"encoding/json"
	"net/http"

	"../data"
)

// GetPeople gets all people from within the api
func GetPeople(res http.ResponseWriter, req *http.Request) {
	// TODO: fix db connection singleton
	// rows, err := db.Query("select * from 'People'")

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// json.NewEncoder(res).Encode(rows)
	json.NewEncoder(res).Encode(data.People)
}
