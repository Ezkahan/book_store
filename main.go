package main

import (
	"github.com/ezkahan/book_store/config"
	router "github.com/ezkahan/book_store/router"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
	// config.SyncDatabase()
}

func main() {
	r := router.Routes()

	r.Run()
}
