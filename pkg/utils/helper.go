package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		fmt.Print("Error on load environment variables")
	}

	value, ok := os.LookupEnv(key)

	if !ok {
		fmt.Printf("Envirement variable not found %s", key)
	}

	return value
}
