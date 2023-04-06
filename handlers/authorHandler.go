package handlers

import (
	"Chat/usecases"
)

type AuthorHandler struct {
	authorUsecase *usecases.AuthorUsecase
}

func NewAuthorHandler(authorUsecase *usecases.AuthorUsecase) AuthorHandler {
	return AuthorHandler{authorUsecase: authorUsecase}
}