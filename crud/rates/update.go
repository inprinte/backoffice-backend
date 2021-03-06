package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "PUT" {

		if utils.Securized(w, r) {
			//parse json from put Request
			var oneRate structures.GetRate
			err := json.NewDecoder(r.Body).Decode(&oneRate)
			utils.CheckErr(err)

			//connect the database
			db := utils.DbConnect()

			sqlQuery := ("UPDATE rate SET stars_number = " + strconv.Itoa(oneRate.Stars_number) + " WHERE rate.id = " + strconv.Itoa(oneRate.Id) + ";")

			// execute the sql query
			_, err = db.Exec(sqlQuery)
			utils.CheckErr(err)

			// close the database connection
			db.Close()

			//create the json response
			json.NewEncoder(w).Encode(oneRate)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
