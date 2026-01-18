package app;

import (
	"orders-service/internal/config"
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateMustRun(cfg *config.Config) {
	var migrationsPath, migrationsTable, command string

	flag.StringVar(&migrationsPath, "migrations-path", "./migrations", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.StringVar(&command, "command", "up", "migration command (up/down)")
	flag.Parse()

	dsn := buildPostgresDSN(cfg)

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		fmt.Sprintf("%s&x-migrations-table=%s", dsn, migrationsTable),
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to create migrator: %v", err))
	}
	defer m.Close()

	switch command {
	case "up":
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")
				return
			}
			panic(err)
		}
	case "down":
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")
				return
			}
			panic(err)
		}
	default:
		panic("unknown command specified")
	}

	fmt.Println("migrations applied successfully")
}
