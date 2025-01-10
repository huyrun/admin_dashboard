package main

import (
	"github.com/huyrun/admin_dashboard/engine"
	_ "github.com/huyrun/go-admin/adapter/gin"                 // web framework adapter
	_ "github.com/huyrun/go-admin/modules/db/drivers/postgres" // sql driver
	_ "github.com/huyrun/themes/sword"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/huyrun/admin_dashboard/tables"
)

func main() {
	startServer()
}

func startServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	eng, err := engine.NewEngine(tables.NewGenerators)
	if err != nil {
		panic(err)
	}

	err = eng.R.Run(":8000")
	if err != nil {
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("closing database connection")
	eng.E.PostgresqlConnection().Close()
}
