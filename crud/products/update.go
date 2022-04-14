package crud

import (
	"encoding/json"
	"inprinteBackoffice/structures"
	"inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)
	(w).Header().Set("Access - Control - Allow - Origin", "*")

	if r.Method == "PUT" {

		if utils.Securized(w, r) {
			// parse json from put Request
			var oneProduct structures.GetProduct

			//parse the body
			err := json.NewDecoder(r.Body).Decode(&oneProduct)
			utils.CheckErr(err)

			// connect the database
			db := utils.DbConnect()

			// create the sql query
			sqlQuery := ("UPDATE product SET name = '" + oneProduct.Name + "' , price = " + strconv.Itoa(oneProduct.Price) + " , description = '" + oneProduct.Description + "' , pending_validation = " + strconv.FormatBool(oneProduct.Pending_validation) + " , is_alive = " + strconv.FormatBool(oneProduct.Is_alive) + " , id_category = (SELECT id from category WHERE name = '" + oneProduct.Category + "') WHERE id = " + strconv.Itoa(oneProduct.Id) + ";")

			// execute the sql query
			_, err = db.Exec(sqlQuery)
			utils.CheckErr(err)

			// delete all the fileUrl
			sqlQuery = ("DELETE FROM product_file WHERE id_product = " + strconv.Itoa(oneProduct.Id) + ";")
			_, err = db.Exec(sqlQuery)
			utils.CheckErr(err)

			// delete all the pictureUrl
			sqlQuery = ("DELETE FROM product_picture WHERE id_product = " + strconv.Itoa(oneProduct.Id) + ";")
			_, err = db.Exec(sqlQuery)
			utils.CheckErr(err)

			// create the sql query
			for i := 0; i < len(oneProduct.FileUrl); i++ {
				sqlQuery = ("INSERT INTO product_file (id_product, url) VALUES (" + strconv.Itoa(oneProduct.Id) + ", '" + oneProduct.FileUrl[i] + "');")

				// execute the sql query
				_, err = db.Exec(sqlQuery)
				utils.CheckErr(err)
			}

			// create the sql query for picture table
			for i := 0; i < len(oneProduct.PictureUrl); i++ {
				sqlQuery = ("INSERT INTO picture (url) VALUES ('" + oneProduct.PictureUrl[i] + "');")

				// execute the sql query
				_, err = db.Exec(sqlQuery)
				utils.CheckErr(err)

				// create the sql query
				sqlQuery = ("INSERT INTO product_picture (id_picture, id_product) VALUES ((SELECT id FROM picture ORDER BY id DESC LIMIT 1), '" + strconv.Itoa(oneProduct.Id) + "');")

				// execute the sql query
				_, err = db.Exec(sqlQuery)
				utils.CheckErr(err)
			}

			// close the database connection
			db.Close()

			// create the json response
			json.NewEncoder(w).Encode(oneProduct)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
