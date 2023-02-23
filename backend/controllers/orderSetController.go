package controllers

import (
	"net/http"

	"github.com/Sokke91/food-connections.git/models"
	"github.com/gin-gonic/gin"
)

func CreateOrderSet(c *gin.Context) {
	var input models.OrderSetInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	orderSet := models.OrderSet{
		Name: input.Name,
	}
	savedOrderSet, err := orderSet.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": savedOrderSet})
	}
	c.JSON(http.StatusOK, gin.H{"data": savedOrderSet})
}

func GetAllOrderSets(c *gin.Context) {
	var orderSets []*models.OrderSet

	orderSets, err := models.LoadAllOrderSets()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orderSets})
}

func DeleteOrderSet(c *gin.Context) {
	orderSet, err := models.FindOrderSetById(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	_, err = orderSet.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func UpdateOrderSet(c *gin.Context) {
}
