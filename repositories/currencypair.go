package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/wardana/currency-exchange/models"
)

type (
	//CurrencyPair initialize currency class
	CurrencyPair struct {
		DB *gorm.DB
	}
	//CurrencyPairInterface is an interface for CurrencyPair entities
	CurrencyPairInterface interface {
		Create(params models.CurrencyPair) (models.CurrencyPair, error)
		Find(params *models.CurrencyPair) ([]models.CurrencyPair, error)
		Update(id int64, params models.CurrencyPair) (models.CurrencyPair, error)
	}
)

//Create is a function to create new record
func (c *CurrencyPair) Create(params models.CurrencyPair) (models.CurrencyPair, error) {
	if errs := c.DB.Create(&params).GetErrors(); len(errs) > 0 {
		return params, errs[0]
	}
	return params, nil
}

//Find is a function to search CurrencyPair data using CurrencyPair code
func (c *CurrencyPair) Find(params *models.CurrencyPair) ([]models.CurrencyPair, error) {
	data := []models.CurrencyPair{}
	if errs := c.DB.Where(params).Find(&data).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}

//Update is a function to update CurrencyPair data
func (c *CurrencyPair) Update(id int64, params models.CurrencyPair) (models.CurrencyPair, error) {
	data := models.CurrencyPair{ID: id}

	if errs := c.DB.Model(&data).UpdateColumns(&params).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}
