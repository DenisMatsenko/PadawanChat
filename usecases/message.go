package usecases

import (
	"Chat/domain"
	"Chat/ports"
)

type MessageUsecase struct {
	dbStorage *ports.DbStorage
}

func NewMessageUsecase(pdb *ports.DbStorage) *MessageUsecase {
	return &MessageUsecase{dbStorage: pdb}
}

func (msgu MessageUsecase) Insert(message domain.Message) error {
	err := msgu.dbStorage.Insert(message)
	if err != nil {
		return err
	}
	return nil
}

func (msgu MessageUsecase) Delete(messageId int) error {
	err := msgu.dbStorage.Delete(messageId)
	if err != nil {
		return err
	}
	return nil
}

func (msgu MessageUsecase) GetAll() ([]domain.Message, error) {
	messages, err := msgu.dbStorage.GetAll()
	if err != nil {
		return nil, err
	}
	return messages, nil
}
