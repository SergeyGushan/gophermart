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
	app.http = app.newHTTPClient()
	app.c = NewContainer(app.pgsql)
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
