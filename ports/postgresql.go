package ports

import (
	"Chat/domain"
	"database/sql"
	"fmt"
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

func (ds *DbStorage) DeleteFromDb(messageId int) (bool, error) {
	script := fmt.Sprintf(`DELETE FROM "Messages" WHERE "id" = %d`, messageId)

	queryResult, err := ds.database.Exec(script)
	if err != nil {
		return false, err
	}

	// ? Shold it be here, or its usecase part?
	rowsCount, err := queryResult.RowsAffected()
	if err != nil || rowsCount != 1 {
		return false, err
	} else {
		return true, nil
	}
}

func (ds *DbStorage) GetAllFromDb() (*sql.Rows, error) {
	script := `SELECT * FROM "Messages"`

	rows, err := ds.database.Query(script)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
