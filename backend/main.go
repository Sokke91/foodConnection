package main

import (
	"fmt"
	"log"

	"github.com/Sokke91/food-connections.git/controllers"
	"github.com/Sokke91/food-connections.git/database"
	"github.com/Sokke91/food-connections.git/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvVariables()
	connectDatabase()
	serverApp()
}

func connectDatabase() {
	database.ConnectDatabase()
	database.DB.AutoMigrate(&models.Order{}, &models.OrderSet{})
}

func loadEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("Env variables loaded")
	}
}

func serverApp() {
	r := gin.Default()
	protectedRoutes := r.Group("/api")
	protectedRoutes.GET("/orders", controllers.GetOrders)
	protectedRoutes.GET("/order/:id", controllers.GetOrderById)
	protectedRoutes.POST("/orders", controllers.CreateOrder)

	protectedRoutes.GET("/ordersets", controllers.GetAllOrderSets)
	protectedRoutes.POST("/ordersets", controllers.CreateOrderSet)
	r.Run()
	fmt.Println("Server started")
}
