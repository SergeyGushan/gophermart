package app

import "time"

func (a *App) RunWorker() error {
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for range ticker.C {
			a.worker.Accrual()
		}
	}()
	return nil
}
