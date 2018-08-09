package controller

import (
	"github.com/labstack/echo"
	"github.com/wardana/currency-exchange/helper"
	"github.com/wardana/currency-exchange/services"
)

//Controller initialize controller class
type Controller struct {
	RateService         services.RateInterface
	CurrencyPairService services.CurrencPairInterface
	Helper              helper.Interface
}

//Interface is an interface for controller class
type Interface interface {
	//exchange rate controller
	CreateNewRate(c echo.Context) error
	RemoveRate(c echo.Context) error
	FindAllRates(c echo.Context) error
	FindLatestDataByDate(c echo.Context) error
	//currency controller
	SaveCurrency(c echo.Context) error
	FindAllCurrency(c echo.Context) error
	UpdateCurrency(c echo.Context) error
	DeleteCurrency(c echo.Context) error
}
