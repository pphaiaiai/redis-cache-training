package repositories

import (
	"redis-cache-training/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(c *fiber.Ctx, id int) (*models.Product, error)
}

type MasterDataRepositories struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &MasterDataRepositories{DB: db}
}

func (r *MasterDataRepositories) GetProductByID(c *fiber.Ctx, id int) (*models.Product, error) {
	product := &models.Product{}

	if err := r.DB.Table("product").First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}
