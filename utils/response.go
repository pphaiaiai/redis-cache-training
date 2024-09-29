package utils

import (
	errorsConstant "redis-cache-training/errors"

	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ResponseSuccess returns a success response
func ResponseSuccess(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(SuccessResponse{
		Code:    code,
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// ResponseError returns an error response
func ResponseError(c *fiber.Ctx, code string) error {
	errorData := errorsConstant.GetError(code)

	return c.Status(errorData.HTTPStatus).JSON(ErrorResponse{
		Status:  errorData.HTTPStatus,
		Code:    errorData.Code,
		Message: errorData.Message,
	})
}
