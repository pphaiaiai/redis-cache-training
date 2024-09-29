package handlers

import (
	"redis-cache-training/internal/services"
	"redis-cache-training/utils"
	"strconv"

	errorsConstant "redis-cache-training/errors"

	"github.com/gofiber/fiber/v2"
)

type HttpProduct struct {
	service services.ProductService
}

func NewHttpProduct(service services.ProductService) *HttpProduct {
	return &HttpProduct{service: service}
}

func (h *HttpProduct) HandleGetProductByID(c *fiber.Ctx) error {

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return utils.ResponseError(c, errorsConstant.ERROR_CODE_INVALID_ID)
	}

	product, err := h.service.GetProductByID(c, idInt)
	if err != nil {
		return utils.ResponseError(c, errorsConstant.ERROR_CODE_INTERNAL_SERVER_ERROR)
	} else {
		return utils.ResponseSuccess(c, fiber.StatusOK, "Success", product)
	}
}
