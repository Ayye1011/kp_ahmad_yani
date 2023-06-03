package main

import (
	"kpahmadyani/configs"
	"kpahmadyani/routes"
)

func main() {
	configs.ConnectDatabase()
	e := routes.Init()
	e.Start(":8080")
}
