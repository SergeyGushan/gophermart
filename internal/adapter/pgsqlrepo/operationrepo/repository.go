package operationrepo

import (
	"database/sql"
	"gophermart/internal/adapter/pgsqlrepo/transactor"
)

type Repository struct {
	transactor.Transactor
}

func NewRepository(conn *sql.DB) *Repository {
	return &Repository{
		transactor.NewTransactor(conn),
	}
}
