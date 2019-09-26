package model

import "github.com/labstack/echo"

// BaseResponse is wrapper of APIs response.
type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse is HTTP Error response structture
type ErrorResponse struct {
	Error     string `json:"error,omitempty"`
	ErrorCode string `json:"error_code,omitempty"`
}

// NewErrorResponse initialises a new error response.
func NewErrorResponse(c echo.Context, err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}
