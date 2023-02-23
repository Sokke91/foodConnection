package controllers

import (
	"net/http"

	"github.com/Sokke91/food-connections.git/models"
	"github.com/gin-gonic/gin"
)

// load orders from database
func GetOrders(c *gin.Context) {
	var orders []models.Order
	models.DB.Find(&orders)
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func GetOrderById(c *gin.Context) {
	var order models.Order
	if err := models.DB.Where("id = ?", c.Param("id")).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
}

func CreateOrder(c *gin.Context) {
	var input models.OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order := models.Order{Name: input.Name, Price: input.Price}
	models.DB.Create(&order)
	c.JSON(http.StatusOK, gin.H{"data": order})
}
