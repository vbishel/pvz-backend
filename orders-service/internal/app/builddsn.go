package app

import (
	"orders-service/internal/config"
	"fmt"
)

func buildPostgresDSN(cfg *config.Config) string {
	pg := cfg.Postgres

	if pg.Database == "" || pg.Username == "" || pg.Host == "" || pg.Port == "" {
		panic("database connection parameters are not set. please set correct environment variables")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		pg.Username,
		pg.Password,
		pg.Host,
		pg.Port,
		pg.Database,
	)

	if pg.Schema != "" {
		dsn = fmt.Sprintf("%s&search_path=%s", dsn, pg.Schema)
	}

	return dsn
}