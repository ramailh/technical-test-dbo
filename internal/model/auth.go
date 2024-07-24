package model

import "time"

type Auth struct {
	AuthID       int64     `json:"auth_id"`
	CustomerID   int64     `json:"customer_id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Session struct {
	SessionID  int64     `json:"session_id"`
	AuthID     int64     `json:"auth_id"`
	CustomerID int64     `json:"customer_id"`
	Token      string    `json:"token"`
	ExpiredAt  time.Time `json:"expired_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type (
	RegisterRequest struct {
		CustomerID int64  `json:"customer_id"`
		Username   string `json:"username"`
		Password   string `json:"password"`
	}

	ChangePasswordRequest struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

type (
	RefreshTokenRequest struct {
		AuthID     int64 `json:"auth_id"`
		CustomerID int64 `json:"customer_id"`
		SessionID  int64 `json:"session_id"`
	}
)
