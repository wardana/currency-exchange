package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wardana/currency-exchange/config"
	"github.com/wardana/currency-exchange/controller"
	"github.com/wardana/currency-exchange/helper"
	"github.com/wardana/currency-exchange/models"
	"github.com/wardana/currency-exchange/repositories"
	"github.com/wardana/currency-exchange/routes"
	"github.com/wardana/currency-exchange/services"
)

func init() {
	config.ReadConfig(&config.Configuration)
	config.SetupEnvironment(&config.Environtment)

	//try to create MySQL table
	config.Environtment.MySQL.CreateTable(&models.Currency{})
	config.Environtment.MySQL.CreateTable(&models.Rate{})
}

func main() {

	helper := &helper.Helper{}

	currencyRepository := &repositories.Currency{DB: config.Environtment.MySQL}
	currencyService := &services.Currency{CurrencyRepository: currencyRepository}

	rateRepository := &repositories.Rate{DB: config.Environtment.MySQL}
	rateService := &services.Rate{RateRepository: rateRepository}

	controller := &controller.Controller{
		RateService:     rateService,
		CurrencyService: currencyService,
		Helper:          helper,
	}

	routes := &routes.Routes{Ctr: controller}

	e := routes.NewRoutes()

	e.Logger.Fatal(e.Start(config.Configuration.BindingPort()))
}
