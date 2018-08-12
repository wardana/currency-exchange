package mocks

import (
	"errors"
	"time"

	"github.com/wardana/currency-exchange/models"
)

//MockRateRepository is a mock type for the Interface type
type MockRateRepository struct{}

//Create provides a mock function with given fields
func (r MockRateRepository) Create(params models.Rate) (models.Rate, error) { return models.Rate{}, nil }

//Find provides a mock function with given fields
func (r MockRateRepository) Find(params *models.Rate) ([]models.Rate, error) {
	return []models.Rate{}, nil
}

//Update provides a mock function with given fields
func (r MockRateRepository) Update(id int64, params models.Rate) (models.Rate, error) {
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
	return []models.ExchangeData{}, nil
}

//ExchangeDataByDate provides a mock function with given fields
func (r MockRateRepository) ExchangeDataByDate(date time.Time) ([]models.RatePayload, error) {
	return []models.RatePayload{}, nil
}
