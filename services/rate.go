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
	Update(id int64, params models.Rate) (models.Rate, error)
	Delete(id int64) error
	HistoricalDataByDate(date time.Time) (*[]models.RatePayload, error)
}

// Create is a function for create new exchange rate data
func (r *Rate) Create(params models.Rate) (models.Rate, error) {

	opts := &models.Rate{
		CurrencyPairID: params.CurrencyPairID,
		ExchangeDate:   params.ExchangeDate,
	}

	data, _ := r.RateRepository.Find(opts)
	if len(data) > 0 {
		//try to update currency right here
		result, err := r.RateRepository.Update(data[0].ID, params)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	result, err := r.RateRepository.Create(params)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindAll is a function for search available exchange rate
func (r *Rate) FindAll() ([]models.Rate, error) {

	data, err := r.RateRepository.Find(&models.Rate{}) //just leave it empty object
	if err != nil {
		return data, err
	}
	return data, nil
}

// Update is a function for update exchange rate data
func (r *Rate) Update(id int64, params models.Rate) (models.Rate, error) {

	opts := &models.Rate{
		CurrencyPairID: params.CurrencyPairID,
		ExchangeDate:   params.ExchangeDate,
	}

	rate, err := r.RateRepository.Find(opts)
	if err != nil {
		return models.Rate{}, err
	}

	if len(rate) > 0 && rate[0].ID != id {
		return models.Rate{}, errors.New("invalid currency code")
	}

	data, err := r.RateRepository.Update(id, params)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Delete is a function for soft delete exchange rate data
func (r *Rate) Delete(id int64) error {

	data, err := r.RateRepository.Find(&models.Rate{ID: id})
	if err != nil || len(data) < 1 {
		return errors.New("data not found")
	}

	currentDate := time.Now()
	_, err = r.RateRepository.Update(id, models.Rate{DeletedAt: &currentDate})
	if err != nil {
		return err
	}
	return nil
}

//HistoricalDataByDate is a service for get historical data using specify date
func (r *Rate) HistoricalDataByDate(date time.Time) (*[]models.RatePayload, error) {

	data, err := r.RateRepository.HistoricalDataByDate(date)
	if err != nil {
		return data, err
	}
	return data, nil

}
