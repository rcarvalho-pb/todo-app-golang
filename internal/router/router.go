package router

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rcarvalho-pb/todo-app-golang/internal/router/routes"
)

func New(db *sqlx.DB) *mux.Router {
	router := mux.NewRouter()
	routes.ConfigRouter(router, db)
	return router
}
