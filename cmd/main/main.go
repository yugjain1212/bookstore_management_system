package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yugjain1212/book_management_system/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	if err := http.ListenAndServe("localhost:9010", r); err != nil {
		log.Fatal(err)
	}

}
