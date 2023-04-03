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

func (ds *DbStorage) DeleteFromDb(messageId int) error {
	script := fmt.Sprintf(`DELETE FROM "Messages" WHERE "id" = %d`, messageId)

	_, err := ds.database.Exec(script)
	if err != nil {
		return err
	}

	return nil
}

func (ds *DbStorage) GetAllFromDb() (*sql.Rows, error) {
	script := `SELECT * FROM "Messages"`

	rows, err := ds.database.Query(script)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
