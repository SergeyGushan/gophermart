package app

import (
	"database/sql"
	"gophermart/migrations"
	"net/http"
	"sync"

	"go.uber.org/zap/zapcore"
	"gophermart/internal/config"
)

type Logger interface {
	Debug(string, ...zapcore.Field)
	Info(string, ...zapcore.Field)
	Error(string, ...zapcore.Field)
	Fatal(string, ...zapcore.Field)
}

type App struct {
	cfg    config.Config
	c      *Container
	cOnce  *sync.Once
	pgsql  *sql.DB
	http   *http.Client
	logger Logger
}

var a *App

func NewApp() (*App, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	app := &App{
		cOnce: &sync.Once{},
		cfg:   cfg,
	}

	app.initLogger()

	pgSqlConn, err := app.newPgSqlConnect(cfg.DatabaseURI)
	if err != nil {
		return nil, err
	}

	app.pgsql = pgSqlConn
	httpClient := app.newHttpClient()
	app.http = httpClient
	app.c = NewContainer(app.pgsql, app.http, cfg.AccrualSystemAddress)

	migration(app.pgsql)

	return app, nil
}

func migration(client *sql.DB) {
	_, err := client.Exec(migrations.CreateUsersTableSql())

	if err != nil {
		panic(err)
	}

	_, err = client.Exec(migrations.CreateOrdersTableSql())

	if err != nil {
		panic(err)
	}

	_, err = client.Exec(migrations.CreateOperationsTableSql())

	if err != nil {
		panic(err)
	}
}
