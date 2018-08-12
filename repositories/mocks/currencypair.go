package mocks

import (
	"errors"

	"github.com/wardana/currency-exchange/models"
)

//MockCurrencyPair is a mock type for the Interface type
type MockCurrencyPair struct{}

//Create provides a mock function with given fields
func (c MockCurrencyPair) Create(params models.CurrencyPair) (models.CurrencyPair, error) {
	//negative test case for create new currency pair (failed to store into db)
	if params.BaseCurrency == "IDR" && params.CounterCurrency == "IDR" {
		return models.CurrencyPair{}, errors.New("some err")
	}
	return models.CurrencyPair{}, nil
}

//Find provides a mock function with given fields
func (c MockCurrencyPair) Find(params *models.CurrencyPair) ([]models.CurrencyPair, error) {

	if params.ID == 1 {
		return []models.CurrencyPair{models.CurrencyPair{ID: 1}}, nil
	}

	//negative test case for create new currency pair (duplicate currency pair) and also for positive test find one currency pair
	if params.BaseCurrency == "IDR" && params.CounterCurrency == "SGD" {
		return []models.CurrencyPair{models.CurrencyPair{}}, nil
	}

	if params.BaseCurrency == "IDR" && params.CounterCurrency == "GBR" {
		return []models.CurrencyPair{models.CurrencyPair{ID: 6}}, nil
	}

	return []models.CurrencyPair{}, nil

}

//Update provides a mock function with given fields
func (c MockCurrencyPair) Update(id int64, params models.CurrencyPair) (models.CurrencyPair, error) {
	if id == 3 {
		return models.CurrencyPair{}, errors.New("mock some error")
	}

	return models.CurrencyPair{}, nil
}
