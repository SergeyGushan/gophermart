package app

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func (a *App) newPgSQLConnect(address string) (*sql.DB, error) {

	db, err := sql.Open("pgx", address)
	if err != nil {
		return nil, err
	}

	return db, err
}
