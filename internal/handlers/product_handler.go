package handlers

import (
	"redis-cache-training/internal/services"
	"redis-cache-training/logging"
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
	logger := logging.Logger.With().Str("method", "HandleGetProductByID").Logger()

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("invalid ID")
		return utils.ResponseError(c, errorsConstant.ERROR_CODE_INVALID_ID)
	}

	product, err := h.service.GetProductByID(c, idInt)
	if err != nil {
		logger.Error().Stack().Err(err).Msg("error getting product by ID")
		return utils.ResponseError(c, errorsConstant.ERROR_CODE_INTERNAL_SERVER_ERROR)
	} else {
		logger.Info().Msg("success getting product by ID")
		return utils.ResponseSuccess(c, fiber.StatusOK, "Success", product)
	}
}
