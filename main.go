package main

import (
	"blast/app"
	"blast/data"
)

func main() {
	go data.ConnectDB()
	app.BuildNewRoutes()
}
