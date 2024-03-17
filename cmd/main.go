package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rcarvalho-pb/todo-app-golang/internal/config"
	"github.com/rcarvalho-pb/todo-app-golang/internal/router"
	"github.com/rcarvalho-pb/todo-app-golang/pkg/db/sqlite3db"
)

func main() {
	config.InitEvenConfigs()
	db, err := sqlite3db.New("db/database.db")
	if err != nil {
		log.Fatal(err)
	}

	r := router.New(db)

	fmt.Printf("Starting server on port %d\n", config.EnvConfigs.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.EnvConfigs.Port), r))
}
