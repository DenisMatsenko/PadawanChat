package ports

import (
	"Chat/domain"
	"Chat/internal/ports/database/gen/PadawanChat/public/model"
	"Chat/internal/ports/database/gen/PadawanChat/public/table"
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
)

type MessageStorage struct {
	database *sql.DB
}

func NewMessageStorage(dbConnect *sql.DB) *MessageStorage {
	return &MessageStorage{database: dbConnect}
}

func (ms *MessageStorage) Insert(message domain.Message) error {
	// * Create query
	stmt := table.Messages.
		INSERT(
			table.Messages.Content,
			table.Messages.AuthorId).
		VALUES(
			postgres.String(message.Content),
			postgres.String(message.AuthorId))

	// * Execute query
	_, err := stmt.Exec(ms.database)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MessageStorage) Delete(messageId int) error {
	// * Create query
	stmt := table.Messages.DELETE().WHERE(table.Messages.ID.EQ(postgres.Int(int64(messageId))))

	// * Execute query
	queryResult, err := stmt.Exec(ms.database)
	if err != nil {
		return err
	}

	// * Check if row was deleted
	rowsCount, err := queryResult.RowsAffected()
	if err != nil {
		return err
	} else if rowsCount != 1 {
		return domain.ErrMessageNotFound
	}

	return nil
}

func (ms *MessageStorage) GetAll() ([]domain.Message, error) {
	// * Create query
	stmt := table.Messages.SELECT(table.Messages.AllColumns)

	// * Execute query
	messagesModel := []model.Messages{}
	err := stmt.Query(ms.database, &messagesModel)
	if err != nil {
		return nil, err
	}

	// * Map model messages arr to domain messages arr
	messagesDomain := []domain.Message{}
	for _, messageModel := range messagesModel {
		messagesDomain = append(messagesDomain, mapModelToDomainMessage(messageModel))
	}

	return messagesDomain, nil
}

func mapModelToDomainMessage(message model.Messages) domain.Message {
	return domain.Message{
		Id:      message.ID,
		Content: *message.Content,
		AuthorId:  *message.AuthorId,
	}
}
