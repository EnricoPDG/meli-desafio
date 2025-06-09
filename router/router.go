package router

import (
	"github.com/EnricoPDG/meli-desafio/logger"
	handler "github.com/EnricoPDG/meli-desafio/router/api/v1"
	"github.com/EnricoPDG/meli-desafio/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var log = logger.GetLogger()

// SetupRouter configures the Gin router with all routes and middleware.
//
// It sets up API versioning under /api/v1, registers handlers for products
// and sellers, and serves the Swagger documentation at /swagger/*any.
//
// Params:
//   - productService: service layer interface for products.
//   - sellerService: service layer interface for sellers.
//
// Returns:
//   - *gin.Engine: configured Gin engine instance.
func SetupRouter(productService service.ProductServiceAPI, sellerService service.SellerServiceAPI) *gin.Engine {
	// Create the default Gin router (includes logger and recovery middleware)
	router := gin.Default()
	log.Info("Starting API...")

	// Initialize handlers with the service dependencies
	productHandler := handler.NewProductHandler(productService)
	sellerHandler := handler.NewSellerHandler(sellerService)

	// API version group
	v1 := router.Group("/api/v1")

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Product routes
	products := v1.Group("/products")
	{
		products.GET("", productHandler.ListProducts)
		products.GET("/:id", productHandler.GetProductByID)
		products.GET("/:id/reviews", productHandler.GetReviewsByProductID)
	}

	// Seller routes
	sellers := v1.Group("/sellers")
	{
		sellers.GET("/:id", sellerHandler.GetSellerByID)
	}

	return router
}
