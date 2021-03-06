package cookie

import (
	"inprinteBackoffice/utils"
)

func CheckLogin(email, password string) bool {
	var passwordDB string
	var is_alive bool
	var id_role, id_user int

	//connect the database
	db := utils.DbConnect()

	sqlQuery := ("SELECT id, password, is_alive, id_role FROM user WHERE email = '" + email + "'")
	rows, err := db.Query(sqlQuery)
	utils.CheckErr(err)

	//parse the query
	for rows.Next() {
		//retrieve the values and check errors
		err = rows.Scan(&id_user, &passwordDB, &is_alive, &id_role)
		utils.CheckErr(err)
	}

	// close the database connection
	db.Close()

	// check informations
	if passwordDB == password && is_alive && id_role == 1 {
		return true
	}
	return false
}
