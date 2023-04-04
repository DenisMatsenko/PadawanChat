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

func (msgu MessageUsecase) DeleteFromDb(messageId int) (bool, error) {
	successful, err := msgu.dbStorage.DeleteFromDb(messageId)
	if err != nil {
		return successful, err
	}
	return successful, nil
}

func (msgu MessageUsecase) GetAllFromDb() ([]domain.Message, error) {
	rows, err := msgu.dbStorage.GetAllFromDb()
	if err != nil {
		return nil, err
	}

	var messages []domain.Message

	for rows.Next() {
		var message domain.Message
		err = rows.Scan(&message.Id, &message.Content, &message.Author)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
