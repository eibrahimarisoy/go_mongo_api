package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	mongoURI := os.Getenv("MONGOURI")
	fmt.Println(mongoURI)
	return mongoURI
}
