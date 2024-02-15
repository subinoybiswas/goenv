package goenv

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return envGetter(key)
	}
	return GetEnvFrmOS(key)
}
func GetEnvFrmOS(key string) (string, error) {
	return envGetter(key)
}
func GetEnvFrmFile(key string) (string, error) {
	if err := godotenv.Overload(".env"); err != nil {
		return "", errors.New("ENV ERROR")
	}
	return envGetter(key)
}
func envGetter(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", errors.New("ENV ERROR")
}
