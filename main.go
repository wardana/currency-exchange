package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/wardana/currency-exchange/configs"
	"github.com/wardana/currency-exchange/controller"
	"github.com/wardana/currency-exchange/helper"
	"github.com/wardana/currency-exchange/models"
	"github.com/wardana/currency-exchange/repositories"
	"github.com/wardana/currency-exchange/routes"
	"github.com/wardana/currency-exchange/services"
)

func seed() {
	config.Environtment.MySQL.CreateTable(&models.Currency{})
	config.Environtment.MySQL.CreateTable(&models.Rate{})
}

func init() {
	config.ReadConfig(&config.Configuration)
	config.SetupEnvironment(&config.Environtment)
}

func main() {

	seed()

	helper := &helper.Helper{}

	currencyRepository := &repositories.Currency{DB: config.Environtment.MySQL}
	currencyService := &services.Currency{CurrencyRepository: currencyRepository}
	controller := &controller.Controller{CurrencyServices: currencyService, Helper: helper}

	routes := &routes.Routes{Ctr: controller}

	e := routes.NewRoutes()

	e.Logger.Fatal(e.Start(config.Configuration.BindingPort()))
}
