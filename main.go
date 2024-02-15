package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type a struct {
	Hi   string
	Meow string
}

func getEnv(key string) (string, error) {
	if err := godotenv.Load(".env"); err != nil {
		return envGetter(key)
	}
	return getEnvFrmFile(key)
}
func getEnvFrmFile(key string) (string, error) {
	return envGetter(key)
}
func envGetter(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", errors.New("ENV ERROR")
}


func main() {
	var envName string
	fmt.Scanln(&envName)
	value, err := getEnv(envName)
	if err!=nil{
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(value)

}

