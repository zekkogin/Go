package main

import (
	"github.com/joho/godotenv"
	"log"
	"wb-L0/nats/sub"
	"wb-L0/pkg/app"
)

func main() {
	// Load local.env
	loadEnv := godotenv.Load("local.env")
	if loadEnv != nil {
		log.Fatal("Something wrong with local.env file in src")
	}

	myApp := app.NewApp()
	defer myApp.Pool.P.Close()
	myApp.Pool.InitTableOrders()
	sub.SubcribeOrders(myApp)

	select {} // Keep the program running to listen to messages
}
