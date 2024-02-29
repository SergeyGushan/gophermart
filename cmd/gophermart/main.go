package main

import (
	"gophermart/internal/app"
	"log"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatal("Fail to create app: ", err)
	}

	if err = a.StartHTTPServer(); err != nil {
		panic(err)
	}
	if err = a.RunWorker(); err != nil {
		panic(err)
	}
}
