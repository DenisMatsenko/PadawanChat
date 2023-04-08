package usecases

import (
	"Chat/domain"
	"Chat/ports"
)

type AuthorUsecase struct {
	authorStorage  *ports.AuthorStorage
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

func (a AuthorUsecase) Update(author domain.Author) error {
	err := a.authorStorage.Update(author)
	if err != nil {
		return err
	}
	return nil
}

func (a AuthorUsecase) Delete(authorId int32) error {
	_, err := a.authorStorage.GetByID(authorId)
	if err != nil {
		return err
	}

	err = a.messageStorage.DeletetAllByAuthorId(authorId)
	if err != nil {
		return err
	}

	err = a.authorStorage.Delete(authorId)
	if err != nil {
		return err
	}
	return nil
}

func (a AuthorUsecase) GetAllMessages(authorId int32) ([]domain.Message, error) {
	_, err := a.authorStorage.GetByID(authorId)
	if err != nil {
		return nil, err
	}

	messages, err := a.messageStorage.GetAllByAuthorId(authorId)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
