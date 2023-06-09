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
	insertModel := []model.Messages{
		{
			Content:  &message.Content,
			AuthorId: message.AuthorId,
		},
	}

	stmt := table.Messages.INSERT(
		table.Messages.AllColumns.Except(table.Messages.ID)).
		MODELS(insertModel)

	_, err := stmt.Exec(ms.database)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MessageStorage) Delete(messageId int) error {
	stmt := table.Messages.
		DELETE().
		WHERE(table.Messages.ID.EQ(postgres.Int(int64(messageId))))

	queryResult, err := stmt.Exec(ms.database)
	if err != nil {
		return err
	}

	rowsCount, err := queryResult.RowsAffected()
	if err != nil {
		return err
	} else if rowsCount != 1 {
		return domain.ErrMessageNotFound
	}

	return nil
}

func (ms *MessageStorage) GetAll() ([]domain.Message, error) {
	stmt := table.Messages.SELECT(table.Messages.AllColumns)

	messagesModel := []model.Messages{}
	err := stmt.Query(ms.database, &messagesModel)
	if err != nil {
		return nil, err
	}

	messagesDomain := []domain.Message{}
	for _, messageModel := range messagesModel {
		messagesDomain = append(messagesDomain, mapModelToDomainMessage(messageModel))
	}

	return messagesDomain, nil
}

func (ms *MessageStorage) GetAllByAuthorId(authorId int32) ([]domain.Message, error) {
	stmt := table.Messages.
		SELECT(table.Messages.AllColumns).
		WHERE(table.Messages.AuthorId.EQ(postgres.Int32(authorId)))

	messagesModel := []model.Messages{}
	err := stmt.Query(ms.database, &messagesModel)
	if err != nil {
		return nil, err
	}

	messagesDomain := []domain.Message{}
	for _, messageModel := range messagesModel {
		messagesDomain = append(messagesDomain, mapModelToDomainMessage(messageModel))
	}

	return messagesDomain, nil
}

func (ms *MessageStorage) DeletetAllByAuthorId(authorId int32) error {
	stmt := table.Messages.
		DELETE().
		WHERE(table.Messages.AuthorId.EQ(postgres.Int32(authorId)))

	_, err := stmt.Exec(ms.database)
	if err != nil {
		return err
	}

	return nil
}

func mapModelToDomainMessage(message model.Messages) domain.Message {
	var content string
	if message.Content != nil {
		content = *message.Content
	}

	return domain.Message{
		Id:       message.ID,
		Content:  content,
		AuthorId: message.AuthorId,
	}
}
