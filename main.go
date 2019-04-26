package main

import (
	"github.com/crud_golang/app"
	"github.com/crud_golang/database"
)

func main() {
	database.StartDB()
	defer database.Finalize()
	app.Start()

}
