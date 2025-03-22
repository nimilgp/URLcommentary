package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version    string
	BaseURL    string
	Dsn        string
	OffsetSize int
}

var Cfg Config = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		Version:    getString("version", "1"),
		BaseURL:    getString("base-url", "localhost:3333"),
		Dsn:        getString("dsn", "user=postgres dbname=urlc"),
		OffsetSize: getInt("offset-size", 5),
	}
}

func getString(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return value
}

func getInt(key string, defaultValue int) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

func getBool(key string, defaultValue bool) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return boolValue
}
