package ports

import (
	"Chat/domain"
	"Chat/internal/ports/database/gen/PadawanChat/public/table"
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

	queryResult, err := ds.database.Exec(script)
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

func (ds *DbStorage) GetAllFromDb() (*sql.Rows, error) {
	stmt := table.Messages.SELECT(table.Messages.AllColumns)
	query, args := stmt.Sql()

	fmt.Println(args...)
	fmt.Println(query)
	// `SELECT * FROM "Messages"`

	rows, err := ds.database.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
