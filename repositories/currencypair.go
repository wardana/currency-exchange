package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/wardana/currency-exchange/models"
)

//CurrencyPair initialize currency class
type CurrencyPair struct {
	DB *gorm.DB
}

//Create is a function to create new record
func (c *CurrencyPair) Create(params models.CurrencyPair) (models.CurrencyPair, error) {
	if err := c.DB.Create(&params).Error; err != nil {
		return params, err
	}
	return params, nil
}

//Find is a function to search CurrencyPair data using CurrencyPair code
func (c *CurrencyPair) Find(params *models.CurrencyPair) ([]models.CurrencyPair, error) {
	data := []models.CurrencyPair{}
	if err := c.DB.Where(params).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

//Update is a function to update CurrencyPair data
func (c *CurrencyPair) Update(id int64, params models.CurrencyPair) (models.CurrencyPair, error) {
	data := models.CurrencyPair{ID: id}

	if err := c.DB.Model(&data).UpdateColumns(&params).Error; err != nil {
		return data, err
	}
	return data, nil
}
