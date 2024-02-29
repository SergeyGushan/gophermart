package app

import (
	"database/sql"
	"gophermart/internal/usecase"

	"gophermart/internal/service/operationservice"
	"gophermart/internal/service/orderservice"
	"gophermart/internal/service/userservice"

	"gophermart/internal/adapter/pgsqlrepo/operationrepo"
	"gophermart/internal/adapter/pgsqlrepo/orderrepo"
	"gophermart/internal/adapter/pgsqlrepo/userrepo"
)

type Container struct {
	pgsql *sql.DB
}

func NewContainer(pgSQLConn *sql.DB) *Container {

	return &Container{
		pgsql: pgSQLConn,
	}
}

func (c *Container) GetUseCase() *usecase.UseCase {

	return usecase.NewUseCase(c.getUserService(), c.getOrderService(), c.getOperationService())
}

func (c *Container) getPgsqlx() *sql.DB {
	return c.pgsql
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

func (c *Container) getUserRepo() *userrepo.Repository {

	return userrepo.NewRepository(c.getPgsqlx())
}

func (c *Container) getOrderRepo() *orderrepo.Repository {

	return orderrepo.NewRepository(c.getPgsqlx())
}

func (c *Container) getOperationRepo() *operationrepo.Repository {

	return operationrepo.NewRepository(c.getPgsqlx())
}
