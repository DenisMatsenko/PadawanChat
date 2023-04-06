package handlers

import (
	"Chat/domain"
	"Chat/usecases"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type MessageHandler struct {
	messageUsecase *usecases.MessageUsecase
}

func NewMessageHadler(messageUsecase *usecases.MessageUsecase) MessageHandler {
	return MessageHandler{messageUsecase: messageUsecase}
}

func (h MessageHandler) MessageCreate(rw http.ResponseWriter, r *http.Request) {
	var message domain.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		sendError(rw, err)
		return
	}

	err = h.messageUsecase.Insert(message)
	if err != nil {
		sendError(rw, err)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (h MessageHandler) MessageDelete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deleteMessageId, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		sendError(rw, err)
		return
	}

	err = h.messageUsecase.Delete(int(deleteMessageId))
	if err != nil {
		sendError(rw, err)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func (h MessageHandler) MessageGetAll(rw http.ResponseWriter, r *http.Request) {
	messages, err := h.messageUsecase.GetAll()
	if err != nil {
		sendError(rw, err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)  
	json.NewEncoder(rw).Encode(messages)                
}

func sendError(rw http.ResponseWriter, err error) {
	fmt.Println(err)
	if err == domain.ErrMessageNotFound || err == domain.ErrAuthorNotFound {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
	}
}
