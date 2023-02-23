package main

import (
	"net/http"

	"github.com/Sokke91/food-connections.git/controllers"
	"github.com/Sokke91/food-connections.git/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	protectedRoutes := r.Group("/api")
	protectedRoutes.GET("/orders", controllers.GetOrders)
	protectedRoutes.GET("/order/:id", controllers.GetOrderById)
	protectedRoutes.POST("/orders", controllers.CreateOrder)

	protectedRoutes.GET("/ordersets", controllers.GetAllOrderSets)
	protectedRoutes.POST("/ordersets", controllers.CreateOrderSet)
	r.Run()
}
