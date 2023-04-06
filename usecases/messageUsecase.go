package usecases

import (
	"Chat/domain"
	"Chat/ports"
)

type MessageUsecase struct {
	messageStorage *ports.MessageStorage
}

func NewMessageUsecase(pdb *ports.MessageStorage) *MessageUsecase {
	return &MessageUsecase{messageStorage: pdb}
}

func (msgu MessageUsecase) Insert(message domain.Message) error {
	err := msgu.messageStorage.Insert(message)
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
