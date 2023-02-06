package main

import (
	"github.com/ezkahan/book_store/database/mysql"
	router "github.com/ezkahan/book_store/router"
)

func init() {
	// mysql.SyncDatabase()
}

func main() {
	db := mysql.Connection()
	defer mysql.CloseConnection(db)

	r := router.InitRoutes()
	r.Run()
}
