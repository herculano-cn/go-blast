package main

import (
	"blast/app"
	"blast/data"
)

func maikbp() {
	go data.ConnectDB()
	app.BuildNewRoutes()
}
