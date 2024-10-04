package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v80"
)

func InitStripe() {

	err := godotenv.Load(".env.local")
	if err != nil {
		log.Println("No .env.local file found, proceeding with system environment variables")
	}

	stripeSecretKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripeSecretKey == "" {
		log.Fatal("STRIPE_SECRET_KEY environment variable not set")
	}

	stripe.Key = stripeSecretKey
}
