package router

import (
	"github/atharvadeshpande/mongoapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/items", controller.GetMyAllItems).Methods("GET")
	router.HandleFunc("/itemcreate", controller.CreateItems).Methods("POST")
	router.HandleFunc("/updateitem/{id}", controller.UpdatedItem).Methods("PUT")
	router.HandleFunc("/delitem", controller.DeleteOneItem).Methods("DELETE")
	router.HandleFunc("/delitems", controller.DeleteAllItem).Methods("DELETE")
	router.HandleFunc("/users", controller.GetMyAllUsers).Methods("GET")
	router.HandleFunc("/createuser", controller.CreateUsers).Methods("POST")
	return router
}
