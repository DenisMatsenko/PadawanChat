package handlers

import (
	"Chat/usecases"
	"net/http"
)

type AuthorHandler struct {
	authorUsecase *usecases.AuthorUsecase
}

func NewAuthorHandler(authorUsecase *usecases.AuthorUsecase) AuthorHandler {
	return AuthorHandler{authorUsecase: authorUsecase}
}

func (h AuthorHandler) AuthorCreate(rw http.ResponseWriter, r *http.Request) {
	
}

func (h AuthorHandler) AuthorDelete(rw http.ResponseWriter, r *http.Request) {

}

func (h AuthorHandler) AuthorGetAllMessages(rw http.ResponseWriter, r *http.Request) {

}