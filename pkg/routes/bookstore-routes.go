package routes

import (
	"github.com/gorilla/mux"
	"github.com/yugjain1212/book_management_system/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/bookstore", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/bookStores", controllers.GetBook).Methods("GET")
	router.HandleFunc("/bookstore/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/bookstore/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/bookstore/{id}", controllers.DeleteBook).Methods("DELETE")
}
