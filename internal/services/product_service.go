package services

import (
	"redis-cache-training/internal/models"
	"redis-cache-training/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

type ProductService interface {
	GetProductByID(c *fiber.Ctx, id int) (*models.Product, error)
}

type ProductServiceImpl struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{repo: repo}
}

func (s *ProductServiceImpl) GetProductByID(c *fiber.Ctx, id int) (*models.Product, error) {
	return s.repo.GetProductByID(c, id)
}
