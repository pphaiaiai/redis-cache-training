package routes

import (
	"redis-cache-training/internal/handlers"
	"redis-cache-training/internal/repositories"
	"redis-cache-training/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupProductRoutes(db *gorm.DB, appGroup fiber.Router) {

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewHttpProduct(productService)

	productRouter := appGroup.Group("/product")

	productRouter.Get("/:id", productHandler.HandleGetProductByID)
}
