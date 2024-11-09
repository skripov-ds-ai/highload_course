package postgres

import (
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/skripov-ds-ai/highload_course/internal/config"
	"math/rand"
)

type DB struct {
	Master *sqlx.DB
	Slaves []*sqlx.DB
}

func (s *DB) ChooseDBForRead() *sqlx.DB {
	// random choose of replicas
	if len(s.Slaves) == 0 {
		return s.Master
	}
	return s.Slaves[rand.Intn(len(s.Slaves))]
}

func newSqlxDB(cfg *config.DBConfig) (*sqlx.DB, error) {
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

func NewDB(cfg *config.DBInstanceConfig) (*DB, error) {
	master, err := newSqlxDB(cfg.Master)
	if err != nil {
		return nil, err
	}

	slaves := make([]*sqlx.DB, 0, len(cfg.Slaves))
	for _, slaveCfg := range cfg.Slaves {
		slave, err := newSqlxDB(slaveCfg)
		if err != nil {
			// TODO: add logging
			continue
		}
		slaves = append(slaves, slave)
	}
	return &DB{
		Master: master,
		Slaves: slaves,
	}, nil
}

func (d *DB) Close() (err error) {
	err = errors.Join(err, d.Master.Close())
	for _, s := range d.Slaves {
		err = errors.Join(err, s.Close())
	}
	return err
}
