package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/wardana/currency-exchange/models"
)

type (
	//Rate initialize exchange rate class
	Rate struct {
		DB *gorm.DB
	}
	//RateInterface is an interface for currency entities
	RateInterface interface {
		Create(params models.Rate) (models.Rate, error)
		Find(params *models.Rate) ([]models.Rate, error)
		Update(id int64, params models.Rate) (models.Rate, error)
	}
)

//Create is a function to create new record
func (c *Rate) Create(params models.Rate) (models.Rate, error) {
	if errs := c.DB.Create(&params).GetErrors(); len(errs) > 0 {
		return params, errs[0]
	}
	return params, nil
}

//Find is a function to search currency data using currency code
func (c *Rate) Find(params *models.Rate) ([]models.Rate, error) {
	data := []models.Rate{}
	if errs := c.DB.Where(params).Find(&data).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}

//Update is a function to update currency data
func (c *Rate) Update(id int64, params models.Rate) (models.Rate, error) {
	data := models.Rate{
		ID: id,
	}
	if errs := c.DB.Model(&data).UpdateColumns(&params).GetErrors(); len(errs) > 0 {
		return data, errs[0]
	}
	return data, nil
}
