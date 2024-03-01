package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) AddRoutes(r *mux.Router) {
	r.HandleFunc("/orders", h.CreateOrder).Methods(http.MethodPost)
	r.HandleFunc("/orders", h.GetOrdersByUserID).Methods(http.MethodGet)
	r.HandleFunc("/balance", h.Balance).Methods(http.MethodGet)
	r.HandleFunc("/balance/withdraw", h.BalanceWithdraw).Methods(http.MethodPost)
	r.HandleFunc("/withdrawals", h.Withdrawals).Methods(http.MethodGet)
}

func (h *Handler) AddAuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", h.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/login", h.Login).Methods(http.MethodPost)
}
