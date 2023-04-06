package usecases

import (
	"Chat/ports"
)

type AuthorUsecase struct {
	authorStorage *ports.AuthorStorage
}

func NewAuthorUsecase(pdb *ports.AuthorStorage) *AuthorUsecase {
	return &AuthorUsecase{authorStorage: pdb}
}