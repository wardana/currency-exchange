package controller

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/wardana/currency-exchange/models"
)

//SaveCurrency is a handler for create new currency pair
func (ctr *Controller) SaveCurrency(c echo.Context) error {

	payload := models.CurrencyPair{}

	// Throw err response when payload is not valid
	if err := c.Bind(&payload); err != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid payload")
	}

	data, err := ctr.CurrencyPairService.Create(payload)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPCreated(c, data)
}

// FindAllCurrency is a handler for get all available currency pair
func (ctr *Controller) FindAllCurrency(c echo.Context) error {

	data, err := ctr.CurrencyPairService.FindAll()
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, data)
}

// UpdateCurrency is a handler for get all available currency pair
func (ctr *Controller) UpdateCurrency(c echo.Context) error {

	payload := models.CurrencyPair{}

	id := c.Param("id")

	if id == "" {
		return ctr.Helper.HTTPBadRequest(c, "id not found")
	}

	// Throw err response when payload is not valid
	if err := c.Bind(&payload); err != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid payload")
	}

	idInt64, errConvert := strconv.ParseInt(id, 10, 64)

	if errConvert != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid type")
	}

	data, err := ctr.CurrencyPairService.Update(idInt64, payload)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, data)
}

// DeleteCurrency is a handler for delete some currency pair data using its id
func (ctr *Controller) DeleteCurrency(c echo.Context) error {

	id := c.Param("id")

	if id == "" {
		return ctr.Helper.HTTPBadRequest(c, "id not found")
	}

	idInt64, errConvert := strconv.ParseInt(id, 10, 64)

	if errConvert != nil {
		return ctr.Helper.HTTPBadRequest(c, "invalid type")
	}

	err := ctr.CurrencyPairService.Delete(idInt64)
	if err != nil {
		return ctr.Helper.HTTPInternalServerError(c, err.Error())
	}

	return ctr.Helper.HTTPSuccess(c, map[string]interface{}{})
}
