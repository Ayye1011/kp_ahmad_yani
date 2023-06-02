package main

import (
	"kpahmadyani/configs"
	"kpahmadyani/routes"
)

func Init() {
	configs.LoadEnv()
	configs.ConnectDatabase()
}

func main() {

	e := routes.Init()
	e.Start(":8000")
}
