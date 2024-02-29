package app

import "time"

func (a *App) RunWorker() error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		a.worker.Accrual()
	}

	return nil
}
