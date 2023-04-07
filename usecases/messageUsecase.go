package usecases

import (
	"Chat/domain"
	"Chat/ports"
)

type MessageUsecase struct {
	messageStorage *ports.MessageStorage
	authorStorage  *ports.AuthorStorage
}

func NewMessageUsecase(ms *ports.MessageStorage, as *ports.AuthorStorage) *MessageUsecase {
	return &MessageUsecase{messageStorage: ms, authorStorage: as}
}

func (msgu MessageUsecase) Insert(message domain.Message) error {

	_, err := msgu.authorStorage.Exist(message.AuthorId)
	if err != nil {
		return err
	}

	err = msgu.messageStorage.Insert(message)
	if err != nil {
		return err
	}
	return nil
}

func (msgu MessageUsecase) Delete(messageId int) error {
	err := msgu.messageStorage.Delete(messageId)
	if err != nil {
		return err
	}
	return nil
}

func (msgu MessageUsecase) GetAll() ([]domain.Message, error) {
	messages, err := msgu.messageStorage.GetAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}
