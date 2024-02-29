package worker

import (
	"database/sql"
	"gophermart/internal/adapter/httprepo/accrualrepo"
	"time"
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

func (w *Worker) Start() {
	ticker := time.Tick(1 * time.Second) // Create a ticker with a 1-second interval

	for range ticker {
		// Your logic here
	}
}
