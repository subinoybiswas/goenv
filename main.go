package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return EnvGetter(key)
	}
	return GetEnvFrmFile(key)
}
func GetEnvFrmFile(key string) (string, error) {
	return EnvGetter(key)
}
func EnvGetter(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", errors.New("ENV ERROR")
}
