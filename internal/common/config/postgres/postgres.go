package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConfiguration struct {
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBOptions         string
	MaxConnection     int
	MaxIdleConnection int
	Driver            string
}

func NewSqlConnection(cfg DBConfiguration) (*sql.DB, error) {
	connURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName)

	db, err := sql.Open(cfg.Driver, connURL)
	if err != nil {
		log.Println("failed to connect to db")
		return nil, err
	}

	if cfg.MaxConnection > 0 {
		db.SetMaxOpenConns(cfg.MaxConnection)
	}
	if cfg.MaxIdleConnection > 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConnection)
	}

	if err := db.Ping(); err != nil {
		log.Println("failed to ping db")
		return nil, err
	}

	log.Println("DB Says Pong!, DB connected")

	return db, nil
}
