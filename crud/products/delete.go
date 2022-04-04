package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "DELETE" {

		//connect the database
		db := utils.DbConnect()

		//get url values
		vars := mux.Vars(r)
		id_product := vars["id_product"]

		//create the sql query
		sqlQuery := ("UPDATE product SET is_alive = false WHERE id = " + id_product + ";")
		_, err := db.Exec(sqlQuery)
		utils.CheckErr(err)

		//close the database connection
		db.Close()

		//create the json response
		json.NewEncoder(w).Encode(structures.JsonResponseProduct{
			Type:    "success",
			Message: "Product deleted",
		})
	}
}