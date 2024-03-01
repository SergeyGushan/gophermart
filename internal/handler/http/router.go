package http

import (
	"github.com/gorilla/mux"
	"gophermart/pkg/logger"
	"gophermart/pkg/middleware"
)

type HandlerRouter interface {
	AddAuthRoutes(r *mux.Router)
	AddRoutes(r *mux.Router)
}

type Router struct {
	router *mux.Router
}

func NewRouter() *Router {
	return &Router{router: mux.NewRouter()}
}

func (r *Router) WithHandler(h HandlerRouter, logger logger.Logger) *Router {
	api := r.router.PathPrefix("/api/user").Subrouter()
	apiAuth := r.router.PathPrefix("/api/user").Subrouter()

	api.Use(middleware.AuthorizationMiddleware)
	api.Use(middleware.JSONMiddleware)

	h.AddRoutes(api)
	h.AddAuthRoutes(apiAuth)

	return r
}
