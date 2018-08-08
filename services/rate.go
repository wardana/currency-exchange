package services

import (
	"errors"
	"time"

	"github.com/wardana/currency-exchange/models"
	"github.com/wardana/currency-exchange/repositories"
)

// Rate services
type Rate struct {
	RateRepository repositories.RateInterface
}

// RateInterface is an interface for rate service
type RateInterface interface {
	Create(params models.Rate) (models.Rate, error)
	FindAll() ([]models.Rate, error)
	Update(id uint64, params models.Rate) (models.Rate, error)
	Delete(id uint64) error
}

// Create is a function for create new exchange rate data
func (c *Rate) Create(params models.Rate) (models.Rate, error) {

	opts := &models.Rate{
		BaseCurrency:    params.BaseCurrency,
		CounterCurrency: params.CounterCurrency,
		ExchangeDate:    params.ExchangeDate,
	}

	data, _ := c.RateRepository.Find(opts)
	if len(data) > 0 {
		//try to update currency right here
		result, err := c.RateRepository.Update(data[0].ID, params)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	result, err := c.RateRepository.Create(params)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindAll is a function for search available exchange rate
func (c *Rate) FindAll() ([]models.Rate, error) {

	data, err := c.RateRepository.Find(&models.Rate{}) //just leave it empty object
	if err != nil {
		return data, err
	}
	return data, nil
}

// Update is a function for update exchange rate data
func (c *Rate) Update(id uint64, params models.Rate) (models.Rate, error) {

	opts := &models.Rate{
		ExchangeDate:    params.ExchangeDate,
		BaseCurrency:    params.BaseCurrency,
		CounterCurrency: params.CounterCurrency,
	}

	rate, err := c.RateRepository.Find(opts)
	if err != nil {
		return models.Rate{}, err
	}

	if len(rate) > 0 && rate[0].ID != id {
		return models.Rate{}, errors.New("invalid currency code")
	}

	data, err := c.RateRepository.Update(id, params)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Delete is a function for delete available currency
func (c *Rate) Delete(id uint64) error {

	rate, err := c.RateRepository.Find(&models.Rate{ID: id})
	if err != nil || len(rate) < 1 {
		return errors.New("currency data not found")
	}

	currentDate := time.Now()
	_, err = c.RateRepository.Update(id, models.Rate{DeletedAt: &currentDate})
	if err != nil {
		return err
	}
	return nil
}