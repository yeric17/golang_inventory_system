package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/yeric17/inventory-system/config"
	"github.com/yeric17/inventory-system/controllers"
	"github.com/yeric17/inventory-system/server"
)

func main() {
	var err error
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/products", controllers.ProductController.Update).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.ProductController.Delete).Methods("DELETE")
	router.HandleFunc("/products", controllers.ProductController.GetAll).Methods("GET")
	router.HandleFunc("/orders", controllers.OrderController.GetAll).Methods("GET")
	router.HandleFunc("/orders", controllers.OrderController.Create).Methods("POST")

	s := server.GetServer(router)

	fmt.Printf("Listening app in port: %s", config.APIPort)
	fmt.Println("")
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
