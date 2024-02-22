package app

import (
	"database/sql"
	"net/http"

	"gophermart/internal/usecase"

	"gophermart/internal/service/accrualservice"
	"gophermart/internal/service/operationservice"
	"gophermart/internal/service/orderservice"
	"gophermart/internal/service/userservice"

	"gophermart/internal/adapter/httprepo/accrualrepo"
	"gophermart/internal/adapter/pgsqlrepo/operationrepo"
	"gophermart/internal/adapter/pgsqlrepo/orderrepo"
	"gophermart/internal/adapter/pgsqlrepo/userrepo"
)

type Container struct {
	pgsql   *sql.DB
	http    *http.Client
	baseURL string

	deps map[string]interface{}
}

func NewContainer(pgSqlConn *sql.DB, httpClient *http.Client, baseURL string) *Container {

	return &Container{
		pgsql:   pgSqlConn,
		http:    httpClient,
		baseURL: baseURL,

		deps: make(map[string]interface{}),
	}
}

func (c *Container) GetUseCase() *usecase.UseCase {

	return usecase.NewUseCase(c.getUserService(), c.getOrderService(), c.getOperationService(), c.getAccrualService())
}

func (c *Container) getPgsqlx() *sql.DB {
	return c.pgsql
}

func (c *Container) getHttp() *http.Client {
	return c.http
}

func (c *Container) getUserService() *userservice.Service {

	return userservice.NewService(c.getUserRepo())
}

func (c *Container) getOrderService() *orderservice.Service {

	return orderservice.NewService(c.getOrderRepo())
}

func (c *Container) getOperationService() *operationservice.Service {

	return operationservice.NewService(c.getOperationRepo())
}

func (c *Container) getAccrualService() *accrualservice.Service {

	return accrualservice.NewService(c.getAccrualRepo())
}

func (c *Container) getUserRepo() *userrepo.Repository {

	return userrepo.NewRepository(c.getPgsqlx())
}

func (c *Container) getOrderRepo() *orderrepo.Repository {

	return orderrepo.NewRepository(c.getPgsqlx())
}

func (c *Container) getOperationRepo() *operationrepo.Repository {

	return operationrepo.NewRepository(c.getPgsqlx())
}

func (c *Container) getAccrualRepo() *accrualrepo.Repository {

	return accrualrepo.NewRepository(c.getHttp(), c.baseURL)
}
