package main

import (
	"auth-service/internal/config"
	"auth-service/internal/app"
)

func main() {
	cfg := config.MustLoad()
	
	app.MustRun(cfg)
}
