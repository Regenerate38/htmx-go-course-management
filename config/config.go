package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return ""
	}
	value, exists := os.LookupEnv(key)
	if !exists {
		fmt.Printf("No value with that key: %s \n", key)
	}
	return value
}
