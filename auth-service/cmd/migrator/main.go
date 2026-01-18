package main

import (
	"auth-service/internal/app"
	"auth-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	app.MigrateMustRun(cfg)
}
