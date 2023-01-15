package main

import (
	"github.com/ezkahan/golab/config"
	router "github.com/ezkahan/golab/router"
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
