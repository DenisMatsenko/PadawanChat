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

type AuthorHandler struct {
	authorUsecase *usecases.AuthorUsecase
}

func NewAuthorHandler(authorUsecase *usecases.AuthorUsecase) AuthorHandler {
	return AuthorHandler{authorUsecase: authorUsecase}
}

func (h AuthorHandler) AuthorCreate(rw http.ResponseWriter, r *http.Request)  {
	var author domain.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		sendError(rw, err)
		return
	}

	err = h.authorUsecase.Insert(author)
	if err != nil {
		sendError(rw, err)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}

func (h AuthorHandler) AuthorDelete(rw http.ResponseWriter, r *http.Request) {
	fmt.Print("AuthorDelete")
	vars := mux.Vars(r)
	deleteAuthorId, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		sendError(rw, err)
		return
	}

	err = h.authorUsecase.Delete(int32(deleteAuthorId))
	if err != nil {
		sendError(rw, err)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

func (h AuthorHandler) AuthorUpdate(rw http.ResponseWriter, r *http.Request) {
	var author domain.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		sendError(rw, err)
		return
	}

	err = h.authorUsecase.Update(author)
	if err != nil {
		sendError(rw, err)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (h AuthorHandler) AuthorGetAllMessages(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		sendError(rw, err)
		return
	}

	messages, err := h.authorUsecase.GetAllMessages(int32(authorId))
	if err != nil {
		sendError(rw, err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)  
	json.NewEncoder(rw).Encode(messages)  
}
