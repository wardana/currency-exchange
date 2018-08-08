package services

import (
	"errors"
	"time"

	"github.com/wardana/currency-exchange/models"
	"github.com/wardana/currency-exchange/repositories"
)

// Currency services
type Currency struct {
	CurrencyRepository repositories.CurrencyInterface
}

// CurrencyInterface is an interface for currency service
type CurrencyInterface interface {
	Create(params models.Currency) (models.Currency, error)
	FindAll() ([]models.Currency, error)
	Update(id uint64, params models.Currency) (models.Currency, error)
	Delete(id uint64) error
}

// Create is a function for create new curency data
func (c *Currency) Create(params models.Currency) (models.Currency, error) {

	opts := &models.Currency{ISOCode: params.ISOCode}

	data, _ := c.CurrencyRepository.Find(opts)
	if len(data) > 0 {
		return data[0], errors.New("duplicate currency code")
	}

	result, err := c.CurrencyRepository.Create(params)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindAll is a function for search available currency
func (c *Currency) FindAll() ([]models.Currency, error) {

	data, err := c.CurrencyRepository.Find(&models.Currency{}) //just leave it empty object
	if err != nil {
		return data, err
	}
	return data, nil
}

// Update is a function for update currency data
func (c *Currency) Update(id uint64, params models.Currency) (models.Currency, error) {

	currency, err := c.CurrencyRepository.Find(&models.Currency{ISOCode: params.ISOCode})
	if err != nil {
		return models.Currency{}, err
	}

	if len(currency) > 0 && currency[0].ID != id {
		return models.Currency{}, errors.New("invalid currency code")
	}

	data, err := c.CurrencyRepository.Update(id, params)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Delete is a function for delete available currency
func (c *Currency) Delete(id uint64) error {

	currency, err := c.CurrencyRepository.Find(&models.Currency{ID: id})
	if err != nil || len(currency) < 1 {
		return errors.New("currency data not found")
	}

	currentDate := time.Now()
	_, err = c.CurrencyRepository.Update(id, models.Currency{DeletedAt: &currentDate})
	if err != nil {
		return err
	}
	return nil
}
