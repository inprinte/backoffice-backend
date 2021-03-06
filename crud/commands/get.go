package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCommand(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "GET" {

		if utils.Securized(w, r) {
			//global vars
			var id int
			var status string

			//connect the database
			db := utils.DbConnect()

			//get url values
			vars := mux.Vars(r)
			id_command := vars["id_command"]

			//create the sql query
			sqlQuery := ("SELECT  command.id AS id, command.state FROM command WHERE command.id = " + id_command + ";")

			//execute the sql query
			row := db.QueryRow(sqlQuery)

			//parse the query
			err := row.Scan(&id, &status)
			utils.CheckErr(err)

			//close the database connection
			db.Close()

			//add the values to the response
			oneCommand := structures.GetCommand{
				Id:     id,
				Status: status,
			}

			//create the json response
			json.NewEncoder(w).Encode(oneCommand)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
