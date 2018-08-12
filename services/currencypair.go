package services

import (
	"errors"
	"reflect"
	"time"

	"github.com/wardana/currency-exchange/models"
	"github.com/wardana/currency-exchange/repositories"
)

// CurrencyPair services
type CurrencyPair struct {
	CurrencyPairRepository repositories.CurrencyPairInterface
	RateRepository         repositories.RateInterface
}

// CurrencPairInterface is an interface for currency service
type CurrencPairInterface interface {
	Create(params models.CurrencyPair) (models.CurrencyPair, error)
	FindAll() ([]models.CurrencyPair, error)
	FindOne(params models.CurrencyPair) (models.CurrencyPair, error)
	Update(id int64, params models.CurrencyPair) (models.CurrencyPair, error)
	Delete(id int64) error
}

// Create is a function for create new curency data
func (c *CurrencyPair) Create(params models.CurrencyPair) (models.CurrencyPair, error) {

	opts := &models.CurrencyPair{BaseCurrency: params.BaseCurrency, CounterCurrency: params.CounterCurrency}

	data, _ := c.CurrencyPairRepository.Find(opts)
	if len(data) > 0 {
		return data[0], errors.New("duplicate currency exchange code")
	}

	result, err := c.CurrencyPairRepository.Create(params)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindAll is a function for search available currency
func (c *CurrencyPair) FindAll() ([]models.CurrencyPair, error) {

	data, err := c.CurrencyPairRepository.Find(&models.CurrencyPair{}) //just leave it empty object
	if err != nil {
		return data, err
	}
	return data, nil
}

// FindOne is a function for search available currency
func (c *CurrencyPair) FindOne(params models.CurrencyPair) (models.CurrencyPair, error) {

	resp, err := c.CurrencyPairRepository.Find(&params)
	if err != nil || reflect.DeepEqual([]models.CurrencyPair{}, resp) {
		return models.CurrencyPair{}, errors.New("data not found")
	}
	return resp[0], nil
}

// Update is a function for update currency data
func (c *CurrencyPair) Update(id int64, params models.CurrencyPair) (models.CurrencyPair, error) {

	opts := &models.CurrencyPair{BaseCurrency: params.BaseCurrency, CounterCurrency: params.CounterCurrency}

	data, _ := c.CurrencyPairRepository.Find(opts)
	if len(data) > 0 && data[0].ID != id {
		return models.CurrencyPair{}, errors.New("invalid currency code")
	}

	result, err := c.CurrencyPairRepository.Update(id, params)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete is a function for delete available currency
func (c *CurrencyPair) Delete(id int64) error {

	//validate id
	currency, err := c.CurrencyPairRepository.Find(&models.CurrencyPair{ID: id})
	if err != nil || len(currency) < 1 {
		return errors.New("currency data not found")
	}

	//soft delete currency data
	currentDate := time.Now()
	_, err = c.CurrencyPairRepository.Update(id, models.CurrencyPair{DeletedAt: &currentDate})
	if err != nil {
		return err
	}

	//soft delete rate rate
	err = c.RateRepository.RemoveByPairID(id)
	if err != nil {
		return err
	}

	return nil
}
