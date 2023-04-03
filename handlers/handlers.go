package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"Chat/domain"
	"Chat/usecases"
	"net/http"
	"github.com/gorilla/mux"
)

type Handler struct {
	usecase *usecases.MessageUsecase
}

func NewHadler(usecase *usecases.MessageUsecase) Handler {
	return Handler{usecase: usecase}
}

func (h Handler) MessageCreate(w http.ResponseWriter, r *http.Request) {
	var message domain.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = h.usecase.InsertToDb(message)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	w.WriteHeader(http.StatusCreated) // # Status
}

func (h Handler) MessageDelete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	deleteMessageId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	err = h.usecase.DeleteFromDb(deleteMessageId)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	rw.WriteHeader(http.StatusNoContent) // # Status
}

func (h Handler) MessageGetAll(rw http.ResponseWriter, r *http.Request) {

	messages, err := h.usecase.GetAllFromDb()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	rw.Header().Set("Content-Type", "application/json") // # Header
	json.NewEncoder(rw).Encode(messages)                // # Body

	rw.WriteHeader(http.StatusOK) // # Status
	fmt.Println("Get all messages")
}

// func (h Handler) MessageDelete(w http.ResponseWriter, r *http.Request) {
// 	var message domain.Message
// 	err := json.NewDecoder(r.Body).Decode(&message)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	}

// 	err = h.usecase.DeleteFromDb(message)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	}

// 	w.WriteHeader(http.StatusNoContent) // # Status
// }

// func MessageDelete(rw http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	deleteMessageId, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 	}

// 	rw.WriteHeader(http.StatusNoContent) // # Status
// 	fmt.Println(deleteMessageId)
// }

// func MessageGetAll(rw http.ResponseWriter, r *http.Request) {

// 	var messages []domain.Message
// 	messages = append(messages, domain.Message{Id: 1, Content: "Hello", Author: "John"})
// 	messages = append(messages, domain.Message{Id: 2, Content: "Hi", Author: "Jane"})

// 	rw.Header().Set("Content-Type", "application/json") // # Header
// 	json.NewEncoder(rw).Encode(messages) // # Body

// 	rw.WriteHeader(http.StatusOK) // # Status
// 	fmt.Println("Get all messages")
// }

// func Test(rw http.ResponseWriter, r *http.Request) {
// 	ports.InsertToDb(domain.Message{Id: 1, Content: "Hello", Author: "John"})

// 	rw.WriteHeader(http.StatusOK) // # Status
// 	fmt.Println("Test")
// }
