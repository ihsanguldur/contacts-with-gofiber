package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//Config is a function which load .env file and gets value with given key.
func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("cannot load .env file.")
	}

	return os.Getenv(key)
}
