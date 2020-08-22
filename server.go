package main

import (
	"echo/db"
	"echo/router"
)

func main() {
	db.Init()

	e := router.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
