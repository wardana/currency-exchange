package helper

import (
	"net/http"

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
}

// SuccessResponse - stored information for success response
type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse - stored information for success response
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// SetErrorResponse - set all properties error response
func (h *Helper) setErrorResponse(httpStatus int, message string, err interface{}) ErrorResponse {
	errResp := ErrorResponse{
		Status:  httpStatus,
		Message: message,
		Error:   err,
	}
	return errResp
}

// SetSuccessResponse - set all properties success response
func (h *Helper) setSuccessResponse(httpStatus int, message string, data interface{}) SuccessResponse {
	succResp := SuccessResponse{
		Status:  httpStatus,
		Message: message,
		Data:    data,
	}
	return succResp
}

// WriteSuccessResponse - write success response to the client
func (h *Helper) writeSuccessResponse(c echo.Context, resp SuccessResponse) error {
	return c.JSON(http.StatusOK, resp)
}

// WriteErrorResponse - write error response to the client
func (h *Helper) writeErrorResponse(c echo.Context, resp ErrorResponse) error {
	return c.JSON(resp.Status, resp)
}

// HTTPBadRequest - bad request response
func (h *Helper) HTTPBadRequest(c echo.Context, msg interface{}) error {
	return h.writeErrorResponse(c, h.setErrorResponse(http.StatusBadRequest, "Bad request", msg))
}

// HTTPSuccess - success request response
func (h *Helper) HTTPSuccess(c echo.Context, data interface{}) error {
	return h.writeSuccessResponse(c, h.setSuccessResponse(http.StatusOK, "Success", data))
}

// HTTPCreated - success request response
func (h *Helper) HTTPCreated(c echo.Context, data interface{}) error {
	return h.writeSuccessResponse(c, h.setSuccessResponse(http.StatusCreated, "Data has been created", data))
}

// HTTPInternalServerError - internal server error response
func (h *Helper) HTTPInternalServerError(c echo.Context, message string) error {
	return h.writeErrorResponse(c, h.setErrorResponse(http.StatusInternalServerError, message, "internal server error"))
}
