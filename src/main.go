package main

import (
	router "go-api/src/routes"
	"log/slog"
)

func main() {

	slog.Info("Initializing router")
	go router.Init()
}
