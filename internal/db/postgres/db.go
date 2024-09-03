package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/skripov-ds-ai/highload_course/internal/config"
)

func NewDB(cfg *config.DBConfig) (*sqlx.DB, error) {
	pool, err := sqlx.Open("pgx", cfg.URI())
	if err != nil {
		return nil, err
	}

	pool.SetMaxOpenConns(cfg.MaxOpenConns)
	pool.SetMaxIdleConns(cfg.MaxIdleConns)
	pool.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	pool.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

	if err := pool.Ping(); err != nil {
		return nil, err
	}
	return pool, nil
}
