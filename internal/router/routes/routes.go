package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func ConfigRouter(r *mux.Router, db *sqlx.DB) *mux.Router {
	var routes []Route

	todoRoutes := initTodoRoutes(db)

	routes = append(routes, todoRoutes...)

	for _, route := range routes {
		if route.Authentication {
			// implement
		} else {
			r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
		}
	}

	fileServe := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServe))
	return r
}
