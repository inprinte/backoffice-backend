package crud

import (
	"encoding/json"
	structures "inprinteBackoffice/structures"
	utils "inprinteBackoffice/utils"
	"net/http"
	"strconv"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	//create cors header
	utils.SetCorsHeaders(&w)

	if r.Method == "POST" {

		if utils.Securized(w, r) {
			// global variables
			var response = structures.JsonResponseUser{}
			var user = structures.CreateUser{}
			var lastInsertID int
			var lastInsertIDAddress int

			// get body
			err := json.NewDecoder(r.Body).Decode(&user)
			utils.CheckErr(err)

			// connect the database
			db := utils.DbConnect()

			// create the sql query
			sql := ("INSERT INTO address (street,city,state,country,zip_code) VALUES (\"" + user.Address.Street + "\", \"" + user.Address.City + "\", \"" + user.Address.State + "\", \"" + user.Address.Country + "\", \"" + user.Address.ZipCode + "\");")

			// execute the sql query
			_, er := db.Exec(sql)
			utils.CheckErr(er)

			// get the last inserted id
			sqlQuer := ("SELECT id FROM address ORDER BY id DESC LIMIT 1;")

			// execute the query
			row := db.QueryRow(sqlQuer)
			err = row.Scan(&lastInsertIDAddress)
			utils.CheckErr(err)

			// create the sql query
			sqlQuery := ("INSERT INTO user (email, first_name, last_name, password, phone, is_alive, id_role, id_address) VALUES ('" + user.Email + "', '" + user.Firstname + "', '" + user.Lastname + "', '" + utils.HashPassword(user.Password) + "', '" + user.Phone + "', " + strconv.FormatBool(user.IsAlive) + ", (SELECT id FROM role WHERE role = '" + user.Role + "'), " + strconv.Itoa(lastInsertIDAddress) + ");")

			// execute the sql query
			_, err = db.Exec(sqlQuery)
			utils.CheckErr(err)

			// get the last inserted id
			sqlQuery = ("SELECT id FROM user ORDER BY id DESC LIMIT 1;")

			// execute the query
			row = db.QueryRow(sqlQuery)
			err = row.Scan(&lastInsertID)
			utils.CheckErr(err)

			// close the database connection
			db.Close()

			// set the response
			response = structures.JsonResponseUser{
				Id:   lastInsertID,
				Type: "success",
				Data: structures.CreateUser{
					Email:     user.Email,
					Firstname: user.Firstname,
					Lastname:  user.Lastname,
					Password:  utils.HashPassword(user.Password),
					Phone:     user.Phone,
					IsAlive:   user.IsAlive,
					Role:      user.Role,
					Address: structures.Address{
						Street:  user.Address.Street,
						City:    user.Address.City,
						State:   user.Address.State,
						Country: user.Address.Country,
						ZipCode: user.Address.ZipCode,
					},
				},
				Message: "New user inserted into DB.",
			}

			//set the response
			json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
