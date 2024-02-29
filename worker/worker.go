package worker

import (
	"database/sql"
	"gophermart/internal/adapter/httprepo/accrualrepo"
)

type Worker struct {
	pgSQL *sql.DB
	repo  *accrualrepo.Repository
}

func NewWorker(pgSQL *sql.DB, repo *accrualrepo.Repository) *Worker {
	return &Worker{
		pgSQL: pgSQL,
		repo:  repo,
	}
}
