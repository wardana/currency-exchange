package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/wardana/currency-exchange/models"
)

type (
	//Currency initialize currency class
	Currency struct {
		DB *gorm.DB
	}
	//CurrencyInterface is an interface for currency entities
	CurrencyInterface interface {
		Create(params models.Currency) (models.Currency, error)
		FindAll() ([]models.Currency, error)
		Find(params *models.Currency) ([]models.Currency, error)
		Update(id uint64, params models.Currency) (models.Currency, error)
	}
)

//Create is a function to create new record
func (c *Currency) Create(params models.Currency) (models.Currency, error) {
	if errs := c.DB.Create(&params).GetErrors(); len(errs) > 0 {
		return params, errs[0]
	}
	return params, nil
}

//Find is a function to search currency data using currency code
func (c *Currency) Find(params *models.Currency) ([]models.Currency, error) {
	data := []models.Currency{}
	if errs := c.DB.Where(params).Find(&data).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}

//FindAll is a function to search currency data
func (c *Currency) FindAll() ([]models.Currency, error) {
	data := []models.Currency{}
	if errs := c.DB.Find(&data).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}

//Update is a function to update currency data
func (c *Currency) Update(id uint64, params models.Currency) (models.Currency, error) {
	data := models.Currency{
		ID: id,
	}
	if errs := c.DB.Model(&data).UpdateColumns(&params).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}
