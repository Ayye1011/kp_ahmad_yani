package main

import (
	"kpahmadyani/configs"
	"kpahmadyani/routes"
	"os"
)

func Init() {
	configs.LoadEnv()
	configs.ConnectDatabase()
}

func main() {

	e := routes.Init()
	e.Start(":" + getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
