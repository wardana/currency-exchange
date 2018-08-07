package routes

import (
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/wardana/currency-exchange/controller"
)

//Routes - Initialize route classs
type Routes struct {
	Ctr controller.Interface
}

// NewRoutes - initialize all routes
func (r *Routes) NewRoutes() *echo.Echo {
	//initialization echo framework
	e := echo.New()
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		//config logger
		e.Use(
			middleware.LoggerWithConfig(
				middleware.LoggerConfig{
					Format: "date=${time_rfc3339}, ip=${remote_ip}, method=${method}, url=${uri}, status=${status}, response_time=${latency_human}\n",
				},
			),
		)
		e.Use(middleware.Recover())

		wg.Done()
	}()

	go func() {
		//Define all routes
		e.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"name":    "currency-exchange",
				"version": "1.0.0",
			})
		})

		e.GET("/api/v1/currency", r.Ctr.FindAll)
		e.POST("/api/v1/currency/create", r.Ctr.Create)
		e.PUT("/api/v1/currency/:id", r.Ctr.UpdateCurrency)
		e.DELETE("/api/v1/currency/:id", r.Ctr.DeleteCurrency)

		wg.Done()
	}()
	wg.Wait()
	return e
}
