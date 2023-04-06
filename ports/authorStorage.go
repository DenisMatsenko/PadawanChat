package ports

import (
	"Chat/domain"
	"Chat/internal/ports/database/gen/PadawanChat/public/table"
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
)

type AuthorStorage struct {
	database *sql.DB
}

func NewAuthorStorage(dbConnect *sql.DB) *AuthorStorage {
	return &AuthorStorage{database: dbConnect}
}

func (as *AuthorStorage) Insert(author domain.Author) error {

	// insertModel := model.Authors{
	// 	Username: &author.Username,
	// }

	// stmt := table.Authors.INSERT(
	// 	table.Authors.AllColumns.Except(table.Authors.ID),
	// 	table.Authors.Username).
	// 	MODELS(insertModel)

	stmt := table.Authors.
		INSERT(
			table.Authors.Username).
		VALUES(
			postgres.String(author.Username))

	_, err := stmt.Exec(as.database)
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthorStorage) Exist(authorId int32) error {
	stmt := table.Authors.SELECT(table.Authors.ID).WHERE(table.Authors.ID.EQ(postgres.Int(int64(authorId))))

	queryResult, err := stmt.Exec(as.database)
	if err != nil {
		return err
	}

	rowsCount, err := queryResult.RowsAffected()
	if err != nil {
		return err
	} else if rowsCount != 1 {
		return domain.ErrAuthorNotFound
	}

	return nil
}