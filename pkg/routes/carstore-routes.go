package routes

import (
	"github.com/Artemson/go-carstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterCarStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreareCar).Methods("POST")
	router.HandleFunc("/book/", controllers.GetCar).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetCarById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateCar).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteCar).Methods("DELETE")
}
