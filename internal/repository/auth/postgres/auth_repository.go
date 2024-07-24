package postgres

import (
	"database/sql"

	"github.com/ramailh/technical-test-dbo/internal/repository/auth"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) auth.AuthRepository {
	return &authRepository{
		db,
	}
}
