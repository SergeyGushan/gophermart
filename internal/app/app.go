package app

import (
	"database/sql"
	"gophermart/internal/adapter/httprepo/accrualrepo"
	"gophermart/migrations"
	"gophermart/worker"
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
	worker *worker.Worker
}

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

	pgSQLConn, err := app.newPgSQLConnect(cfg.DatabaseURI)

	if err != nil {
		return nil, err
	}

	app.pgsql = pgSQLConn
	httpClient := app.newHTTPClient()
	repo := accrualrepo.NewRepository(httpClient, cfg.AccrualSystemAddress)
	app.http = httpClient
	app.c = NewContainer(app.pgsql, app.http, cfg.AccrualSystemAddress)
	app.worker = worker.NewWorker(pgSQLConn, repo)
	migration(app.pgsql)

	return app, nil
}

func migration(client *sql.DB) {
	_, err := client.Exec(migrations.CreateUsersTableSQL())

	if err != nil {
		panic(err)
	}

	_, err = client.Exec(migrations.CreateOrdersTableSQL())

	if err != nil {
		panic(err)
	}

	_, err = client.Exec(migrations.CreateOperationsTableSQL())

	if err != nil {
		panic(err)
	}
}
