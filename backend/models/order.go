package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Nr         uint
	Name       string
	Price      float64
	OrderSetID uint
}

type OrderInput struct {
	Nr    uint    `json:"nr" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}
