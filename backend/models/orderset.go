package models

import (
	"github.com/Sokke91/food-connections.git/database"
	"gorm.io/gorm"
)

type OrderSet struct {
	gorm.Model
	Name   string
	orders []Order
}

type OrderSetInput struct {
	Name string `json:"name" binding:"required"`
}

type OrderSetUpdateInput struct {
	Name string `json:"name"`
}

// Methods
func (orderSet *OrderSet) Save() (*OrderSet, error) {
	err := database.DB.Create(&orderSet).Error
	if err != nil {
		return &OrderSet{}, nil
	}
	return orderSet, nil
}

func (orderSet *OrderSet) Update(input OrderSetUpdateInput) (*OrderSet, error) {
	err := database.DB.Model(&orderSet).Updates(input).Error
	if err != nil {
		return &OrderSet{}, err
	}
	return orderSet, nil
}

func (orderSet *OrderSet) Delete() (bool, error) {
	err := database.DB.Delete(&orderSet).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// Functions

func FindOrderSetById(id string) (*OrderSet, error) {
	var orderSet *OrderSet
	err := database.DB.Where("id=?", id).First(&orderSet).Error
	if err != nil {
		return &OrderSet{}, err
	}
	return orderSet, nil
}

func LoadAllOrderSets() ([]*OrderSet, error) {
	var oderSets []*OrderSet

	err := database.DB.Find(&oderSets).Error
	if err != nil {
		return []*OrderSet{}, err
	}
	return oderSets, nil
}
