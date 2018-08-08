package services

import (
	"errors"
	"reflect"
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
	FindOne(params models.Currency) (models.Currency, error)
	Update(id int64, params models.Currency) (models.Currency, error)
	Delete(id int64) error
}

// Create is a function for create new curency data
func (c *Currency) Create(params models.Currency) (models.Currency, error) {

	opts := &models.Currency{BaseCurrency: params.BaseCurrency, CounterCurrency: params.CounterCurrency}

	data, _ := c.CurrencyRepository.Find(opts)
	if len(data) > 0 {
		return data[0], errors.New("duplicate currency exchange code")
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

// FindOne is a function for search available currency
func (c *Currency) FindOne(params models.Currency) (models.Currency, error) {

	resp, err := c.CurrencyRepository.Find(&params)
	if err != nil || reflect.DeepEqual([]models.Currency{}, resp) {
		return models.Currency{}, err
	}
	return resp[0], nil
}

// Update is a function for update currency data
func (c *Currency) Update(id int64, params models.Currency) (models.Currency, error) {

	opts := &models.Currency{BaseCurrency: params.BaseCurrency, CounterCurrency: params.CounterCurrency}

	data, _ := c.CurrencyRepository.Find(opts)
	if len(data) > 0 && data[0].ID != id {
		return models.Currency{}, errors.New("invalid currency code")
	}

	result, err := c.CurrencyRepository.Update(id, params)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete is a function for delete available currency
func (c *Currency) Delete(id int64) error {

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
