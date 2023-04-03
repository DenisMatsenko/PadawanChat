package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"Chat/domain"
	"net/http"
	"github.com/gorilla/mux"
)


func MessageCreate(rw http.ResponseWriter, r *http.Request) {
	var message domain.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	rw.WriteHeader(http.StatusCreated) // # Status
	fmt.Println(message)
}

func MessageDelete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	deleteMessageId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	rw.WriteHeader(http.StatusNoContent) // # Status
	fmt.Println(deleteMessageId)
}

func MessageGetAll(rw http.ResponseWriter, r *http.Request) {

	var messages []domain.Message
	messages = append(messages, domain.Message{Id: 1, Content: "Hello", Author: "John"})
	messages = append(messages, domain.Message{Id: 2, Content: "Hi", Author: "Jane"})

	rw.Header().Set("Content-Type", "application/json") // # Header
	json.NewEncoder(rw).Encode(messages) // # Body

	rw.WriteHeader(http.StatusOK) // # Status
	fmt.Println("Get all messages")
}