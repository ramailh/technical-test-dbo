package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ramailh/technical-test-dbo/internal/common/config/redis"
	"github.com/ramailh/technical-test-dbo/internal/common/consts"
	"github.com/ramailh/technical-test-dbo/internal/common/env"
	"github.com/ramailh/technical-test-dbo/internal/common/keys"
	"github.com/ramailh/technical-test-dbo/internal/common/response"
	"github.com/ramailh/technical-test-dbo/internal/common/token"
)

const (
	AuthHeader = "Authorization"
)

type AuthMiddleware interface {
	AuthJWT(c *gin.Context)
}

type authJWT struct {
	rds redis.Redis
}

func NewMiddlewareAuth(rds redis.Redis) AuthMiddleware {
	return &authJWT{
		rds: rds,
	}
}

func (md *authJWT) AuthJWT(c *gin.Context) {
	bearerToken := c.GetHeader(AuthHeader)
	if bearerToken == "" {
		c.JSON(http.StatusBadRequest, response.ErrTokenNotExist)
		c.Abort()
		return
	}

	tokens := strings.Split(bearerToken, " ")
	if len(tokens) != 2 {
		c.JSON(http.StatusBadRequest, response.ErrInvalidToken)
		c.Abort()
		return
	}

	tokenJWT := tokens[1]
	claims := &token.Claims{}
	_, err := jwt.ParseWithClaims(tokenJWT, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := env.JWTSigningKey
		return []byte(secret), nil
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrInvalidToken)
		c.Abort()
		return
	}

	key := keys.CacheKeyGenerator(consts.CacheBlacklistSessionIDKey, fmt.Sprint(claims.SessionID))
	_, err = md.rds.Get(context.Background(), key).Result()
	if err == nil {
		c.JSON(http.StatusBadGateway, response.ErrInvalidSession)
		c.Abort()
		return
	}

	c.Set("session_id", claims.SessionID)
	c.Set("customer_id", claims.CustomerID)
	c.Set("auth_id", claims.AuthID)

	c.Next()
}
