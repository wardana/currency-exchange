package helper

import (
	"github.com/labstack/echo"
)

// Helper initialize Helper class
type Helper struct{}

// Interface initialize Helper interface
type Interface interface {
	HTTPBadRequest(c echo.Context, msg interface{}) error
	HTTPSuccess(c echo.Context, data interface{}) error
	HTTPInternalServerError(c echo.Context, message string) error
	HTTPCreated(c echo.Context, data interface{}) error
	MinMaxAverageInSlices(array []float64) (float64, float64, float64)
}
