package app

import (
	"gophermart/pkg/logger"
	"log"
)

func (a *App) initLogger() {
	l, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	a.logger = l
}
