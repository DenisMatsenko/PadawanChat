package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"Chat/handlers"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/message/create", handlers.MessageCreate).Methods("POST")
	mux.HandleFunc("/message/delete/{id}", handlers.MessageDelete).Methods("DELETE")
	mux.HandleFunc("/message/get/all", handlers.MessageGetAll).Methods("GET")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe("0.0.0.0:8080", mux)
}