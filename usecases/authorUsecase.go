package usecases

import (
	"Chat/domain"
	"Chat/ports"
)

type AuthorUsecase struct {
	authorStorage *ports.AuthorStorage
	messageStorage *ports.MessageStorage
}

func NewAuthorUsecase(as *ports.AuthorStorage, ms *ports.MessageStorage) *AuthorUsecase {
	return &AuthorUsecase{authorStorage: as, messageStorage: ms}
}

func (a AuthorUsecase) Insert(author domain.Author) error {
	err := a.authorStorage.Insert(author)
	if err != nil {
		return err
	}
	return nil
}