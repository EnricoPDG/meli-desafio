package v1

import (
	"github.com/EnricoPDG/meli-desafio/logger"
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/EnricoPDG/meli-desafio/response"
	"github.com/EnricoPDG/meli-desafio/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var log = logger.GetLogger()

type ProductHandler struct {
	service service.ProductServiceAPI
}

func NewProductHandler(service service.ProductServiceAPI) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetProductByID godoc
// @Summary Get product by ID
// @Description Get details of a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.JSONSuccessResultProduct
// @Failure 400 {object} response.JSONErrorResponse
// @Failure 404 {object} response.JSONErrorResponse
// @Failure 500 {object} response.JSONErrorResponse
// @Router /api/v1/products/{id} [get]
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	log.Debug("GetProductByID called")
	strID := c.Param("id")
	id, err := uuid.Parse(strID)
	if err != nil {
		log.Error("error parsing uuid", zap.Error(err))
		response.BadRequestResponse(c, err)
		return
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		log.Error("error getting product", zap.Error(err))
		response.ChooseErrorResponse(c, err)
		return
	}

	response.SuccessResponseProduct(c, *product)
}

// GetReviewsByProductID godoc
// @Summary Get reviews by product ID
// @Description Get reviews
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} response.JSONSuccessResultReview
// @Failure 400 {object} response.JSONErrorResponse
// @Failure 404 {object} response.JSONErrorResponse
// @Failure 500 {object} response.JSONErrorResponse
// @Router /api/v1/products/{id}/reviews [get]
func (h *ProductHandler) GetReviewsByProductID(c *gin.Context) {
	log.Debug("GetReviewsByProductID called")
	strID := c.Param("id")
	id, err := uuid.Parse(strID)
	if err != nil {
		log.Error("error parsing uuid", zap.Error(err))
		response.BadRequestResponse(c, err)
		return
	}

	reviews, err := h.service.GetReviewsByProductID(id)
	if err != nil {
		log.Error("error getting reviews", zap.Error(err))
		response.ChooseErrorResponse(c, err)
		return
	}

	response.SuccessResponseReviews(c, reviews)
}

// ListProducts godoc
// @Summary List products
// @Description List products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} response.JSONSuccessResultProducts
// @Failure 400 {object} response.JSONErrorResponse
// @Failure 404 {object} response.JSONErrorResponse
// @Failure 500 {object} response.JSONErrorResponse
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Router /api/v1/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	var q model.ListProductQueryParameters

	if err := c.ShouldBindQuery(&q); err != nil {
		log.Error("error parsing query parameters", zap.Error(err))
		response.BadRequestResponse(c, err)
		return
	}

	products, err := h.service.ListProducts(q.Page, q.Limit)
	if err != nil {
		log.Error("error getting products", zap.Error(err))
		response.ChooseErrorResponse(c, err)
		return
	}

	response.SuccessResponseProducts(c, products)
}
