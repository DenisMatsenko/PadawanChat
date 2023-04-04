package handlers

import (
	"Chat/domain"
	"Chat/usecases"
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	// "os"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	usecase *usecases.MessageUsecase
}

func NewHadler(usecase *usecases.MessageUsecase) Handler {
	return Handler{usecase: usecase}
}

// ! status ok after error ?
func (h Handler) MessageCreate(rw http.ResponseWriter, r *http.Request) {
	var message domain.Message
	err := json.NewDecoder(r.Body).Decode(&message) 
	if err != nil {
		sendError(rw, err)
		return
	}

	err = h.usecase.InsertToDb(message)
	if err != nil {
		sendError(rw, err)
		return
	}

	rw.WriteHeader(http.StatusCreated) // # Status
}

func (h Handler) MessageDelete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deleteMessageId, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(rw, err)
		return
	}

	err = h.usecase.DeleteFromDb(deleteMessageId)
	if err != nil {
		sendError(rw, err)
		return
	}

	rw.WriteHeader(http.StatusNotFound) // # Status
}

func (h Handler) MessageGetAll(rw http.ResponseWriter, r *http.Request) {
	messages, err := h.usecase.GetAllFromDb()
	if err != nil {
		sendError(rw, err)
		return
	}

	rw.Header().Set("Content-Type", "application/json") // # Header
	json.NewEncoder(rw).Encode(messages) // # Body
	rw.WriteHeader(http.StatusOK) // # Status
}

func sendError(rw http.ResponseWriter, err error) {
	fmt.Println(err)
	rw.Header().Set("Content-Type", "application/json") // # Header
	rw.Write([]byte(err.Error()))
	rw.WriteHeader(http.StatusNotFound) // # Status
}