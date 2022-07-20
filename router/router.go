package main

import (
	handler "CRUD/handlerFunc"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func init() {
	r.HandleFunc("/", handler.GetUsers).Methods("GET")
	r.HandleFunc("/getuser/{id}", handler.GetUser).Methods("GET")
	r.HandleFunc("/crud", handler.CreateUser).Methods("POST")
	r.HandleFunc("/update/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete/{id}", handler.DeleteUser).Methods("DELETE")
}

func main() {
	log.Fatal(http.ListenAndServe(":8090", r))
	fmt.Println("Sucess...")
}
