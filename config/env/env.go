package config

import (
	"fmt" // to print the error
	"os" // to get the environment variables
	"github.com/joho/godotenv" // to load the environment variables from the .env file
)

func load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}


func GetString(key string, fallback string) string {
	load()
	value, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println("Environment variable not found")
		return fallback
	}
	return value
}