package app

import "time"

func (a *App) RunWorker() error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			a.worker.Accrual()
		}
	}
}
