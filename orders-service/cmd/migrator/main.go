package main

import (
	"orders-service/internal/app"
	"orders-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	app.MigrateMustRun(cfg)
}
