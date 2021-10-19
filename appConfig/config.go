package appConfig

import (
	"os"
	"strconv"
)

type Config struct {
	HttpPort         int
	ConnectionString string
	Storage          string // "db" or "mem"
	JwtSecret        string
}

func NewConfig() (Config, error) {
	var cfg Config
	var err error
	httpPort, err := strconv.Atoi(getEnv("HTTP_PORT", "8080"))
	if err != nil {
		return cfg, err
	}
	return Config{
		HttpPort:         httpPort,
		ConnectionString: getEnv("CONNECTION_STRING", "root:toor@tcp(localhost:3306)/alta?charset=utf8&parseTime=True&loc=Local"),
		Storage:          getEnv("STORAGE", "db"), // "db" or "mem"
		JwtSecret:        getEnv("JWT_SECRET", "rahasiaBanget"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
