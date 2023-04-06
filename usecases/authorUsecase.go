package usecases

import (
	"Chat/domain"
	"Chat/ports"
)

type AuthorUsecase struct {
	authorStorage *ports.AuthorStorage
}

func NewAuthorUsecase(pdb *ports.AuthorStorage) *AuthorUsecase {
	return &AuthorUsecase{authorStorage: pdb}
}

func (a AuthorUsecase) Insert(author domain.Author) error {
	err := a.authorStorage.Insert(author)
	if err != nil {
		return err
	}
	return nil
}