package main

import "jtn/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start("localhost:8002"))
}
