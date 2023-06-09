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

//go:generate go run github.com/go-jet/jet/v2/cmd/jet -dsn=postgres://postgres:Dm2016dM@localhost:5432/PadawanChat?sslmode=disable -path=./internal/ports/database/gen

func main() {
	dbConnection, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=Dm2016dM dbname=PadawanChat sslmode=disable")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	messageStorage := ports.NewMessageStorage(dbConnection)
	authorStorage := ports.NewAuthorStorage(dbConnection)
	messageUsecase := usecases.NewMessageUsecase(messageStorage, authorStorage)
	authorUsecase := usecases.NewAuthorUsecase(authorStorage, messageStorage)
	messageHandler := handlers.NewMessageHadler(messageUsecase)
	authorHandler := handlers.NewAuthorHandler(authorUsecase)

	mux := mux.NewRouter()
	mux.HandleFunc("/message/create", messageHandler.MessageCreate).Methods("POST")
	mux.HandleFunc("/message/{id}/delete", messageHandler.MessageDelete).Methods("DELETE")
	mux.HandleFunc("/message/get/all", messageHandler.MessageGetAll).Methods("GET")
	mux.HandleFunc("/author/create", authorHandler.AuthorCreate).Methods("POST")
	mux.HandleFunc("/author/{id}/update", authorHandler.AuthorUpdate).Methods("PUT")
	mux.HandleFunc("/author/{id}/delete", authorHandler.AuthorDelete).Methods("DELETE")
	mux.HandleFunc("/author/{id}/messages", authorHandler.AuthorGetAllMessages).Methods("GET")

	fmt.Println("Server running on port 8080")
	http.ListenAndServe("0.0.0.0:8080", mux)
}
