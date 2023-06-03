package main

import (
	"kpahmadyani/configs"
	"kpahmadyani/routes"
	"os"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDatabase()

}

func main() {
	e := routes.Init()
	e.Start(":" + getPort())
	// e.Start(":8080")
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
