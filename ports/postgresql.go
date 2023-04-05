package ports

import (
	"Chat/domain"
	"Chat/internal/ports/database/gen/PadawanChat/public/table"
	"Chat/internal/ports/database/gen/PadawanChat/public/model"
	"database/sql"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
)

type DbStorage struct {
	database *sql.DB
}

func NewDbStorage(dbConnect *sql.DB) *DbStorage {
	return &DbStorage{database: dbConnect}
}

func (ds *DbStorage) InsertToDb(message domain.Message) error {
	script := fmt.Sprintf(`INSERT INTO "Messages" ("content", "author") VALUES ('%s', '%s')`, message.Content, message.Author)

	_, err := ds.database.Exec(script)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DbStorage) DeleteFromDb(messageId int) error {
	stmt := table.Messages.DELETE().WHERE(table.Messages.ID.EQ(postgres.Int(int64(messageId))))

	queryResult, err := stmt.Exec(ds.database)
	if err != nil {
		return err
	}

	rowsCount, err := queryResult.RowsAffected()
	
	if err != nil {
		return err
	} else if rowsCount != 1 {
		return fmt.Errorf("no message with id: %d", messageId)
	}

	return nil
}

func (ds *DbStorage) GetAllFromDb() ([]domain.Message, error) {
	// * Create query
	stmt := table.Messages.SELECT(table.Messages.AllColumns)
	
	// * Execute query
	messagesModel := []model.Messages{}
	err := stmt.Query(ds.database, &messagesModel)
	if err != nil { 
		return nil, err
	}

	// * Map model messages arr to domain messages arr
	var messagesDomain []domain.Message = []domain.Message{}
	for _, messageModel := range messagesModel {
		messagesDomain = append(messagesDomain, mapModelToDomainMessage(messageModel))
	}

	return messagesDomain, nil
}

func mapModelToDomainMessage(message model.Messages) domain.Message{
	return domain.Message{
		Id:  message.ID,
		Content: *message.Content,
		Author: *message.Author,
	}
}