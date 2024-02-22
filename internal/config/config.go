package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
	"log"
)

type Config struct {
	RunAddress           string `env:"RUN_ADDRESS"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	DatabaseURI          string `env:"DATABASE_URI"`
}

func NewConfig() (Config, error) {
	conf := Config{}

	flag.StringVar(&conf.RunAddress, "a", "localhost:8081", "адрес и порт запуска сервиса")
	flag.StringVar(&conf.AccrualSystemAddress, "r", "http://localhost:8080", "адрес системы расчёта начислений")
	flag.StringVar(&conf.DatabaseURI, "d",
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=disable",
			`localhost`,
			`go_user`,
			`go_password`,
			`go_gophermart`,
		),
		"адрес подключения к базе данных",
	)

	flag.Parse()

	err := env.Parse(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf, nil
}
