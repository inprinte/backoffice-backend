package utils

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func SetCorsHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Max-Age", "86400")
	(*w).Header().Set("Access-Control-Expose-Headers", "*")

}

func SetXTotalCountHeader(w *http.ResponseWriter, totalCount int) {
	(*w).Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	(*w).Header().Set("X-Total-Count", strconv.Itoa(totalCount))
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "inprinteAdmin"
	dbPass := "louisensueur"
	dbProtocol := "tcp"
	dbIp := "178.170.14.134"
	dbPort := "3306"
	dbName := "inprinte"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"("+dbIp+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetAllParams(r *http.Request, table string) (string, string) {
	containsOrder := true
	containsSort := true
	containsStart := true
	containsEnd := true
	rangeBy := ""
	orderBy := ""

	urlOrder, ok := r.URL.Query()["_order"]
	if !ok || len(urlOrder[0]) < 1 {
		log.Println("Url Param 'order' is missing")
		containsOrder = false
	}

	urlSort, ok := r.URL.Query()["_sort"]
	if !ok || len(urlSort[0]) < 1 {
		log.Println("Url Param 'sort' is missing")
		containsSort = true
	}

	urlStart, ok := r.URL.Query()["_start"]
	if !ok || len(urlStart[0]) < 1 {
		log.Println("Url Param 'start' is missing")
		containsStart = false
	}

	urlEnd, ok := r.URL.Query()["_end"]
	if !ok || len(urlEnd[0]) < 1 {
		log.Println("Url Param 'End' is missing")
		containsEnd = false
	}

	if containsOrder && containsSort {
		if table == "rate" && urlSort[0] == "id" {
			urlSort[0] = "rate.id"
		}
		orderBy = " ORDER BY " + urlSort[0] + " " + urlOrder[0]
	}

	if containsStart && containsEnd {
		rangeBy = " LIMIT " + urlStart[0] + "," + urlEnd[0]
	}
	return orderBy, rangeBy

}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	CheckErr(err)
	return string(bytes)
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
