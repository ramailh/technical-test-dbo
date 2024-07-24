package token

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ramailh/technical-test-dbo/internal/common/env"
)

func GenerateJWTCust(sessionID, custID, authID int64, expiredAt time.Time) (string, error) {
	claims := Claims{
		SessionID:  sessionID,
		AuthID:     authID,
		CustomerID: custID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawToken.SignedString([]byte(env.JWTSigningKey))
	if err != nil {
		log.Println(err)
		return "", err
	}

	return token, nil
}
