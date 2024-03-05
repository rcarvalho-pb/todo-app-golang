package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/todo-app-golang/internal/router"
	"github.com/rcarvalho-pb/todo-app-golang/pkg/db/sqlite3db"
)

func main() {
	db, err := sqlite3db.New("db/database.db")
	if err != nil {
		log.Fatal(err)
	}

	r := router.New(db)

	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe(":3300", r))
}
