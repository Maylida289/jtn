package main

import (
	"jtn/database"
	"jtn/routes"
)

func main() {
	database.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start("localhost:8002"))
}
