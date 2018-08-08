package controller

import (
	"reflect"
	"strconv"

	"github.com/labstack/echo"
	"github.com/wardana/currency-exchange/models"
)

// CreateNewRate is a handler for storing new exchange rate data
func (ctr *Controller) CreateNewRate(c echo.Context) error {

	payload := models.RatePayload{}

	// Throw err response when payload is not valid
	if err := c.Bind(&payload); err != nil {
		return ctr.Helper.HTTPBadRequest(c, err.Error())
	}

	opts := models.Currency{
		BaseCurrency:    payload.BaseCurrency,
		CounterCurrency: payload.CounterCurrency,
	}

	currency, err := ctr.CurrencyService.FindOne(opts)
	if err != nil || reflect.DeepEqual(models.Currency{}, currency) {
		return ctr.Helper.HTTPBadRequest(c, "currency not found")
	}

	params := models.Rate{
		CurrencyID:   currency.ID,
		ExchangeDate: payload.ExchangeDate,
		ExchangeRate: payload.ExchangeRate,
	}

	_, err = ctr.RateService.Create(params)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPCreated(c, payload)
}

// RemoveRate is a handler for soft delete rate data using it's id
func (ctr *Controller) RemoveRate(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return ctr.Helper.HTTPBadRequest(c, "id not found")
	}

	idInt64, errConvert := strconv.ParseInt(id, 10, 64)

	if errConvert != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid type")
	}

	err := ctr.RateService.Delete(idInt64)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, map[string]interface{}{})
}

//FindAllRates is a handler for find all exchange rates
func (ctr *Controller) FindAllRates(c echo.Context) error {

	data, err := ctr.RateService.FindAll()
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, data)
}
