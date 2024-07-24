package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/subosito/gotenv"
)

var (
	jwtSigningKey = "JWT_SIGNING_KEY"

	port = "PORT"

	dbUser        = "DB_USER"
	dbPassword    = "DB_PASSWORD"
	dbHost        = "DB_HOST"
	dbPort        = "DB_PORT"
	dbName        = "DB_NAME"
	dbMaxConn     = "DB_MAX_CONN"
	dbMaxIdleConn = "DB_MAX_IDLE_CONN"

	redisHost      = "REDIS_HOST"
	redisPort      = "REDIS_PORT"
	redisPassword  = "REDIS_PASSWORD"
	redisDB        = "REDIS_DB"
	redisMaxIdle   = "REDIS_MAX_IDLE"
	redisMaxActive = "REDIS_MAX_ACTIVE"
)

var (
	JWTSigningKey string

	Port string

	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	DBMaxConn     int
	DBMaxIdleConn int

	RedisHost      string
	RedisPort      string
	RedisPassword  string
	RedisDB        int
	RedisMaxIdle   int
	RedisMaxActive int
)

func LoadEnv() {
	err := gotenv.Load()
	if err != nil {
		log.Println(err)
	}

	JWTSigningKey, err = getEnvString(jwtSigningKey)
	if err != nil {
		log.Fatal(err)
	}

	Port, err = getEnvString(port)
	if err != nil {
		log.Fatal(err)
	}

	DBUser, err = getEnvString(dbUser)
	if err != nil {
		log.Fatal(err)
	}
	DBPassword, err = getEnvString(dbPassword)
	if err != nil {
		log.Fatal(err)
	}
	DBHost, err = getEnvString(dbHost)
	if err != nil {
		log.Fatal(err)
	}
	DBPort, err = getEnvString(dbPort)
	if err != nil {
		log.Fatal(err)
	}
	DBName, err = getEnvString(dbName)
	if err != nil {
		log.Fatal(err)
	}
	DBMaxConn, err = getEnvInt(dbMaxConn)
	if err != nil {
		log.Fatal(err)
	}
	DBMaxIdleConn, err = getEnvInt(dbMaxIdleConn)
	if err != nil {
		log.Fatal(err)
	}

	RedisHost, err = getEnvString(redisHost)
	if err != nil {
		log.Fatal(err)
	}
	RedisPort, err = getEnvString(redisPort)
	if err != nil {
		log.Fatal(err)
	}
	RedisPassword, err = getEnvString(redisPassword)
	RedisDB, err = getEnvInt(redisDB)
	if err != nil {
		log.Fatal(err)
	}
	RedisMaxIdle, err = getEnvInt(redisMaxIdle)
	if err != nil {
		log.Fatal(err)
	}
	RedisMaxActive, err = getEnvInt(redisMaxActive)
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvString(env string) (string, error) {
	res := os.Getenv(env)
	if res == "" {
		return res, fmt.Errorf("env %s is empty", env)
	}
	return res, nil
}

func getEnvInt(env string) (int, error) {
	res := os.Getenv(env)
	if res == "" {
		return 0, fmt.Errorf("env %s is empty", env)
	}

	resInt, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return resInt, nil
}
