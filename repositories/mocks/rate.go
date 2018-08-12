package mocks

import (
	"errors"
	"time"

	"github.com/wardana/currency-exchange/models"
)

//MockRateRepository is a mock type for the Interface type
type MockRateRepository struct{}

//Create provides a mock function with given fields
func (r MockRateRepository) Create(params models.Rate) (models.Rate, error) {
	if params.CurrencyPairID == 1 {
		return models.Rate{}, errors.New("mock some err")
	}
	return models.Rate{}, nil
}

//Find provides a mock function with given fields
func (r MockRateRepository) Find(params *models.Rate) ([]models.Rate, error) {
	if params.CurrencyPairID == 2 {
		return []models.Rate{models.Rate{ID: 2}, models.Rate{ID: 3}}, nil
	}
	if params.CurrencyPairID == 3 {
		return []models.Rate{models.Rate{ID: 2}}, nil
	}
	if params.CurrencyPairID == 4 {
		return []models.Rate{}, errors.New("mock some err")
	}

	if params.ID == 1 {
		return []models.Rate{models.Rate{ID: 1}}, nil
	}

	return []models.Rate{}, nil
}

//Update provides a mock function with given fields
func (r MockRateRepository) Update(id int64, params models.Rate) (models.Rate, error) {
	if id == 3 {
		return models.Rate{}, errors.New("mock some err")
	}
	if params.ID == 3 {
		return models.Rate{}, errors.New("mock some err")
	}
	if id == 4 {
		return models.Rate{}, errors.New("mock some err")
	}
	return models.Rate{}, nil
}

//RemoveByPairID provides a mock function with given fields
func (r MockRateRepository) RemoveByPairID(id int64) error {
	if id == 2 {
		return errors.New("id not found")
	}
	return nil
}

//TrendDataByCurrency provides a mock function with given fields
func (r MockRateRepository) TrendDataByCurrency(base, counter string) ([]models.ExchangeData, error) {
	if base == "USD" && counter == "IDR" {
		return []models.ExchangeData{}, errors.New("mock some err ")
	}
	return []models.ExchangeData{}, nil
}

//ExchangeDataByDate provides a mock function with given fields
func (r MockRateRepository) ExchangeDataByDate(date time.Time) ([]models.RatePayload, error) {
	dateWithoutTime, _ := time.Parse("2006-01-02", "2018-08-12")

	if dateWithoutTime == date {
		return []models.RatePayload{}, errors.New("mock some err")
	}

	return []models.RatePayload{}, nil
}
