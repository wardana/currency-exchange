package controller

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/wardana/currency-exchange/helper"
	"github.com/wardana/currency-exchange/models"
	"github.com/wardana/currency-exchange/services"
)

//Controller initialize controller class
type Controller struct {
	CurrencyServices services.CurrencyInterface
	RateServices     services.RateInterface
	Helper           helper.Interface
}

//Interface is an interface for controller class
type Interface interface {
	//currency controller
	SaveCurrency(c echo.Context) error
	FindAllCurrency(c echo.Context) error
	UpdateCurrency(c echo.Context) error
	DeleteCurrency(c echo.Context) error
	//exchange rate controller
	SaveNewRate(c echo.Context) error
}

//SaveCurrency is a handler for create new currency
func (ctr *Controller) SaveCurrency(c echo.Context) error {

	payload := models.Currency{}

	// Throw err response when payload is not valid
	if err := c.Bind(&payload); err != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid payload")
	}

	data, err := ctr.CurrencyServices.Create(payload)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPCreated(c, data)
}

// FindAllCurrency is a handler for get all available currency
func (ctr *Controller) FindAllCurrency(c echo.Context) error {

	data, err := ctr.CurrencyServices.FindAll()
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, data)
}

// UpdateCurrency is a handler for get all available currency
func (ctr *Controller) UpdateCurrency(c echo.Context) error {

	payload := models.Currency{}

	id := c.Param("id")

	if id == "" {
		return ctr.Helper.HTTPBadRequest(c, "id not found")
	}

	// Throw err response when payload is not valid
	if err := c.Bind(&payload); err != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid payload")
	}

	idUint64, errConvert := strconv.ParseUint(id, 10, 64)

	if errConvert != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid type")
	}

	data, err := ctr.CurrencyServices.Update(idUint64, payload)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, data)
}

// DeleteCurrency is a handler for delete some currency data using its id
func (ctr *Controller) DeleteCurrency(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return ctr.Helper.HTTPBadRequest(c, "id not found")
	}

	idUint64, errConvert := strconv.ParseUint(id, 10, 64)

	if errConvert != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid type")
	}

	err := ctr.CurrencyServices.Delete(idUint64)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, map[string]interface{}{})
}

// SaveNewRate is a handler for storing new exchange rate data
func (ctr *Controller) SaveNewRate(c echo.Context) error {

	payload := models.Rate{}

	// Throw err response when payload is not valid
	if err := c.Bind(&payload); err != nil {
		return ctr.Helper.HTTPBadRequest(c, err.Error())
	}

	fmt.Println("payload bind")
	fmt.Println(payload.ExchangeRate)

	data, err := ctr.RateServices.Create(payload)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPCreated(c, data)
}
