package services

import (
	"errors"
	"strconv"
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
	Update(id string, params models.Currency) (models.Currency, error)
	Delete(id string) error
}

// Create is a function for create new curency data
func (c *Currency) Create(params models.Currency) (models.Currency, error) {

	opts := &models.Currency{Code: params.Code}

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
	data := []models.Currency{}
	data, err := c.CurrencyRepository.FindAll()
	if err != nil {
		return data, err
	}
	return data, nil
}

// Update is a function for update currency data
func (c *Currency) Update(id string, params models.Currency) (models.Currency, error) {
	idUint64, _ := strconv.ParseUint(id, 10, 64)

	currency, err := c.CurrencyRepository.Find(&models.Currency{Code: params.Code})
	if err != nil {
		return models.Currency{}, err
	}

	if len(currency) > 0 && currency[0].ID != idUint64 {
		return models.Currency{}, errors.New("invalid currency code")
	}

	data, err := c.CurrencyRepository.Update(idUint64, params)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Delete is a function for delete available currency
func (c *Currency) Delete(id string) error {
	idUint64, _ := strconv.ParseUint(id, 10, 64)

	currency, err := c.CurrencyRepository.Find(&models.Currency{ID: idUint64})
	if err != nil {
		return errors.New("invalid currency ID")
	}

	if len(currency) > 0 && currency[0].ID != idUint64 {
		return errors.New("invalid currency code")
	}

	currentDate := time.Now()
	_, err = c.CurrencyRepository.Update(idUint64, models.Currency{DeletedAt: &currentDate})
	if err != nil {
		return err
	}
	return nil
}
