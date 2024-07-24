package token

import "github.com/golang-jwt/jwt"

type Claims struct {
	SessionID  int64 `json:"session_id"`
	AuthID     int64 `json:"auth_id"`
	CustomerID int64 `json:"customer_id"`
	jwt.StandardClaims
}
