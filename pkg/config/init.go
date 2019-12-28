package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

// Init config init
func Init() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	})
}
