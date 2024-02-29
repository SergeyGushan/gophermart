package app

import (
	"gophermart/internal/adapter/httprepo/accrualrepo"
	"gophermart/worker"
	"time"
)

func (a *App) RunWorker() error {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			worker.Accrual(a.pgsql, accrualrepo.NewRepository(a.newHTTPClient(), a.cfg.AccrualSystemAddress))
		}
	}()
	return nil
}
