package operationrepo

import (
	"database/sql"
	"gophermart/internal/adapter/pgsqlrepo"
)

type Repository struct {
	pgsqlrepo.Transactor
}

func NewRepository(conn *sql.DB) *Repository {
	return &Repository{
		pgsqlrepo.NewTransactor(conn),
	}
}
