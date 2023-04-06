package handlers

import (
	"Chat/domain"
	"Chat/usecases"
	"encoding/json"
	"net/http"
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

}

func (h AuthorHandler) AuthorGetAllMessages(rw http.ResponseWriter, r *http.Request) {

}
