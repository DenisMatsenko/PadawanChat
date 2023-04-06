package ports

import "database/sql"

type AuthorStorage struct {
	database *sql.DB
}

func NewAuthorStorage(dbConnect *sql.DB) *AuthorStorage {
	return &AuthorStorage{database: dbConnect}
}