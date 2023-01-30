package main

import (
	"github.com/ezkahan/book_store/config"
	"github.com/ezkahan/book_store/database/mysql"
	router "github.com/ezkahan/book_store/router"
)

func init() {
	config.LoadEnv()
	mysql.Connection()
	// config.SyncDatabase()
}

func main() {
	r := router.InitRoutes()

	defer mysql.CloseConnection(mysql.DB)

	r.Run()
}
