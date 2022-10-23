package main

import (
	r "praktikum/routes"
)

func main() {

	routes := r.Init()

	routes.Logger.Fatal(routes.Start(":8000"))
}
