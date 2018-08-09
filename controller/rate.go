package controller

import (
	"reflect"
	"strconv"
	"time"

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

	opts := models.CurrencyPair{
		BaseCurrency:    payload.BaseCurrency,
		CounterCurrency: payload.CounterCurrency,
	}

	currencyPair, err := ctr.CurrencyPairService.FindOne(opts)
	if err != nil || reflect.DeepEqual(models.CurrencyPair{}, currencyPair) {
		return ctr.Helper.HTTPBadRequest(c, "currency not found")
	}

	floatExchangeRate, err := strconv.ParseFloat("0.4", 64)
	if err != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid rate format, rate should be number")
	}

	params := models.Rate{
		CurrencyPairID: currencyPair.ID,
		ExchangeDate:   payload.ExchangeDate,
		ExchangeRate:   floatExchangeRate,
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

//FindExchangeDataByDate find latest data using date as parameter
func (ctr *Controller) FindExchangeDataByDate(c echo.Context) error {

	var date time.Time

	dateStr := c.QueryParam("date")

	if dateStr == "" {
		tmp := time.Now()

		dateWithoutTime, _ := time.Parse("2006-01-02", tmp.Format("2006-01-02"))
		date = dateWithoutTime
	} else {
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			return ctr.Helper.HTTPBadRequest(c, "invalid date format")
		}

		date = t
	}

	data, err := ctr.RateService.ExchangeDataByDate(date)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, data)
}

//TrendDataByCurrency find latest data using iso currency code as parameter
func (ctr *Controller) TrendDataByCurrency(c echo.Context) error {

	var arrRate []float64

	type ExchangeData struct {
		ExchangeDate float64 `json:"exchange_date"`
		ExchangeRate float64 `json:"exchange_rate"`
	}

	base := c.QueryParam("base")
	counter := c.QueryParam("counter")

	if base == "" && counter == "" {
		return ctr.Helper.HTTPBadRequest(c, "you need to specify base and counter currency")
	}

	data, err := ctr.RateService.TrendDataByCurrency(base, counter)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	if len(data) > 0 {
		for _, item := range data {
			arrRate = append(arrRate, item.ExchangeRate)
		}
	}

	min, max, average := ctr.Helper.MinMaxAverageInSlices(arrRate)

	//construct response
	response := &struct {
		BaseCurrency    string                `json:"base_currency"`
		CounterCurrency string                `json:"counter_currency"`
		Average         float64               `json:"average"`
		Variance        float64               `json:"variance"`
		History         []models.ExchangeData `json:"history"`
	}{
		BaseCurrency:    base,
		CounterCurrency: counter,
		Average:         average,
		Variance:        max - min,
		History:         data,
	}

	return ctr.Helper.HTTPSuccess(c, response)
}
