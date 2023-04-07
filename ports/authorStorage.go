package ports

import (
	"Chat/domain"
	"Chat/internal/ports/database/gen/PadawanChat/public/table"
	"database/sql"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type AuthorStorage struct {
	database *sql.DB
}

func NewAuthorStorage(dbConnect *sql.DB) *AuthorStorage {
	return &AuthorStorage{database: dbConnect}
}

func (as *AuthorStorage) Insert(author domain.Author) error {
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

func (as *AuthorStorage) Update(author domain.Author) error {
	stmt := table.Authors.
		UPDATE(table.Authors.Username).
		SET(postgres.String(author.Username)).
		WHERE(table.Authors.ID.EQ(postgres.Int(int64(author.Id))))

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

func (as *AuthorStorage) Delete(authorId int32) error {
	stmt := table.Authors.
		DELETE().
		WHERE(table.Authors.ID.EQ(postgres.Int(int64(authorId))))

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

func (as *AuthorStorage) Exist(authorId int32) (*domain.Author, error) {
	stmt := table.Authors.
		SELECT(table.Authors.AllColumns).
		WHERE(table.Authors.ID.EQ(postgres.Int(int64(authorId))))

	author := domain.Author{}
	err := stmt.Query(as.database, &author)
	if err == qrm.ErrNoRows {
		return nil, domain.ErrAuthorNotFound
	} else if err != nil {
		return nil, err
	}

	return &author, nil
}
