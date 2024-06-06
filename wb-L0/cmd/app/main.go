package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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
	myApp.C.LoadFromDatabase(myApp.Pool)
	defer myApp.Pool.Close()
	myApp.Pool.InitTableOrders()
	sub.SubcribeOrders(myApp)

	myApp.R.GET("/retrieve/:order_uid", RetrieveHandler(myApp))
	myApp.R.Run()
	select {} // Keep the program running to listen to messages
}

func RetrieveHandler(app app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.Param("order_uid")
		model := app.C.GetByID(uid)
		if model != nil {
			c.JSON(http.StatusOK, model)
		} else {
			model = app.Pool.GetRowByID(context.Background(), uid)
			if model == nil {
				c.Status(http.StatusNotFound)
				return
			}
			c.JSON(http.StatusOK, model)
		}
	}
}
