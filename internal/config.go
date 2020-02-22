package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config config

type config struct {
	NatsUrl string
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	natsHost := os.Getenv("NATS_HOST")
	natsPort := os.Getenv("NATS_PORT")

	Config.NatsUrl = fmt.Sprintf("%s:%s", natsHost, natsPort)
}
