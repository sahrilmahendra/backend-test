package main

import (
	"erajaya/config"
	"erajaya/middlewares"
	"erajaya/routes"
)

// function main
func main() {
	config.InitDB()
	e := routes.New()
	// take notes http method activity
	middlewares.LogMiddlewares(e)
	// start on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
