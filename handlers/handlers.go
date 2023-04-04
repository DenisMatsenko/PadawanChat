package handlers

import (
	"Chat/domain"
	"Chat/usecases"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	usecase *usecases.MessageUsecase
}

func NewHadler(usecase *usecases.MessageUsecase) Handler {
	return Handler{usecase: usecase}
}

func (h Handler) MessageCreate(w http.ResponseWriter, r *http.Request) {
	// ? Should it be here, or its usecase part?
	var message domain.Message
	err := json.NewDecoder(r.Body).Decode(&message) 
	errorCheck(err)

	err = h.usecase.InsertToDb(message)
	errorCheck(err)

	w.WriteHeader(http.StatusCreated) // # Status
}

func (h Handler) MessageDelete(rw http.ResponseWriter, r *http.Request) {

	// ? Should it be here, or its usecase part?
	vars := mux.Vars(r)
	deleteMessageId, err := strconv.Atoi(vars["id"])
	errorCheck(err)

	successful, err := h.usecase.DeleteFromDb(deleteMessageId)
	errorCheck(err)

	if successful {
		rw.WriteHeader(http.StatusOK) // # Status
	} else {
		rw.WriteHeader(http.StatusNotFound) // # Status
	}
}

func (h Handler) MessageGetAll(rw http.ResponseWriter, r *http.Request) {
	messages, err := h.usecase.GetAllFromDb()
	errorCheck(err)

	rw.Header().Set("Content-Type", "application/json") // # Header
	json.NewEncoder(rw).Encode(messages) // # Body // ? Should it be here, or its usecase part?

	rw.WriteHeader(http.StatusOK) // # Status
	fmt.Println("Get all messages")
}

func errorCheck(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}