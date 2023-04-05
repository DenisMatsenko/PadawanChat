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

func (msgu MessageUsecase) InsertToDb(message domain.Message) error {
	err := msgu.dbStorage.InsertToDb(message)
	if err != nil {
		return err
	}
	return nil
}

func (msgu MessageUsecase) DeleteFromDb(messageId int) error {
	err := msgu.dbStorage.DeleteFromDb(messageId)
	if err != nil {
		return err
	}
	return nil
}

func (msgu MessageUsecase) GetAllFromDb() ([]domain.Message, error) {
	messages, err := msgu.dbStorage.GetAllFromDb()
	if err != nil {
		return nil, err
	}

	return messages, nil
}
