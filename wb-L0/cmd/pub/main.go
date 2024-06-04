package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
	"wb-L0/nats"
	"wb-L0/nats/pub"
)

func main() {

	loadEnv := godotenv.Load("local.env")
	if loadEnv != nil {
		log.Fatal("Something wrong with local.env file in src")
	}

	subject := os.Getenv("SUBJECT")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sc := nats.Connect(ctx, "pub")
	defer sc.Close()

	for {
		pub.PublishOrder(subject, sc)
		time.Sleep(100 * time.Second)
	}
}
