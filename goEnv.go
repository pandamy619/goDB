package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func goEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")
	}
}