package main

import (
	category "inprinteBackoffice/crud/category"
	commandLines "inprinteBackoffice/crud/commandLines"
	commands "inprinteBackoffice/crud/commands"
	products "inprinteBackoffice/crud/products"
	rates "inprinteBackoffice/crud/rates"
	roles "inprinteBackoffice/crud/roles"
	users "inprinteBackoffice/crud/users"
	"net/http"

	"github.com/gorilla/mux"
)

func handleUsers(router *mux.Router) {
	router.HandleFunc("/users", users.GetAll).Methods("GET")
	router.HandleFunc("/users/{id_user}", users.GetOne).Methods("GET")
	router.HandleFunc("/users/{id_user}", users.UpdateOne).Methods("OPTIONS", "PUT")
	// router.HandleFunc("/users/{id_user}", users.delete).Methods("DELETE")
}

func handleCommand(router *mux.Router) {
	router.HandleFunc("/commands", commands.GetAll).Methods("GET")
	router.HandleFunc("/commands/{id_command}", commands.GetOne).Methods("GET")
	router.HandleFunc("/commands/{id_command}", commands.UpdateOne).Methods("PUT", "OPTIONS")
}

func handleCommandLine(router *mux.Router) {
	router.HandleFunc("/commandLines", commandLines.GetAll).Methods("GET")
	router.HandleFunc("/commandLines/{id_commandLine}", commandLines.GetOne).Methods("GET")
	router.HandleFunc("/commandLines/{id_commandLine}", commandLines.UpdateOne).Methods("OPTIONS", "PUT")
}

func handleRate(router *mux.Router) {
	router.HandleFunc("/rates", rates.GetAll).Methods("GET")
	router.HandleFunc("/rates/{id_rate}", rates.GetOne).Methods("GET")
	router.HandleFunc("/rates/{id_rate}", rates.UpdateOne).Methods("OPTIONS", "PUT")
}

func handleProducts(router *mux.Router) {
	// router.HandleFunc("/products", products.insert).Methods("POST")
	router.HandleFunc("/products", products.GetAll).Methods("GET")
	router.HandleFunc("/products/{id_product}", products.GetOne).Methods("GET")
	// router.HandleFunc("/products/{id_product}", products.update).Methods("UPDATE")
	// router.HandleFunc("/products/{id_product}", products.delete).Methods("DELETE")
}

func handleCategory(router *mux.Router) {
	router.HandleFunc("/categories", category.InsertOne).Methods("OPTIONS", "POST")
	router.HandleFunc("/categories", category.GetAll).Methods("GET")
	router.HandleFunc("/categories/{id_category}", category.GetOne).Methods("GET")
	router.HandleFunc("/categories/{id_category}", category.UpdateOne).Methods("OPTIONS", "PUT")
	// router.HandleFunc("/category/{id_product}", category.delete).Methods("DELETE")
}

func handleRoles(router *mux.Router) {
	router.HandleFunc("/roles", roles.InsertOne).Methods("OPTIONS", "POST")
	router.HandleFunc("/roles", roles.GetAll).Methods("GET")
	router.HandleFunc("/roles/{id_role}", roles.GetOne).Methods("GET")
	router.HandleFunc("/roles/{id_role}", roles.UpdateOne).Methods("OPTIONS", "PUT")
	// router.HandleFunc("/role/{id_product}", role.delete).Methods("DELETE")
}

func main() {
	//create a new mux router
	router := mux.NewRouter()

	//handle all the paths
	handleUsers(router)
	handleCommand(router)
	handleCommandLine(router)
	handleRate(router)
	handleProducts(router)
	handleCategory(router)
	handleRoles(router)

	http.ListenAndServe(":8080", router)
}
