package main

import (
	"Chat/handlers"
	"Chat/ports"
	"Chat/usecases"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Dm2016dM dbname=PadawanChat sslmode=disable")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v", db)

	dbStorage := ports.NewDbStorage(db)
	usecase := usecases.NewMessageUsecase(dbStorage)
	handler := handlers.NewHadler(usecase)

	mux := mux.NewRouter()
	mux.HandleFunc("/message/create", handler.MessageCreate).Methods("POST")
	mux.HandleFunc("/message/delete/{id}", handler.MessageDelete).Methods("DELETE")
	mux.HandleFunc("/message/get/all", handler.MessageGetAll).Methods("GET")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe("0.0.0.0:8080", mux)
}
