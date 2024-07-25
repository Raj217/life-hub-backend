package main

import (
	"life-hub-backend/config"
	"life-hub-backend/routes"
)

func main() {
	config.Init()
	routes.Handle()
}
