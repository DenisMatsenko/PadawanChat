package ports

import (
	"Chat/domain"
	"Chat/internal/ports/database/gen/PadawanChat/public/model"
	"Chat/internal/ports/database/gen/PadawanChat/public/table"
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
)

type DbStorage struct {
	database *sql.DB
}

func NewDbStorage(dbConnect *sql.DB) *DbStorage {
	return &DbStorage{database: dbConnect}
}

func (ds *DbStorage) Insert(message domain.Message) error {
	// * Create query
	stmt := table.Messages.
		INSERT(
			table.Messages.Content,
			table.Messages.Author).
		VALUES(
			postgres.String(message.Content),
			postgres.String(message.Author))

	// * Execute query
	_, err := stmt.Exec(ds.database)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DbStorage) Delete(messageId int) error {
	// * Create query
	stmt := table.Messages.DELETE().WHERE(table.Messages.ID.EQ(postgres.Int(int64(messageId))))

	// * Execute query
	queryResult, err := stmt.Exec(ds.database)
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

func (ds *DbStorage) GetAll() ([]domain.Message, error) {
	// * Create query
	stmt := table.Messages.SELECT(table.Messages.AllColumns)

	// * Execute query
	messagesModel := []model.Messages{}
	err := stmt.Query(ds.database, &messagesModel)
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
		Author:  *message.Author,
	}
}
