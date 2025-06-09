package service

import (
	"errors"
	"github.com/EnricoPDG/meli-desafio/logger"
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/EnricoPDG/meli-desafio/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var log = logger.GetLogger()

// ProductServiceAPI defines methods to interact with products and their reviews.
type ProductServiceAPI interface {
	// GetProductByID returns a product by its ID.
	GetProductByID(id uuid.UUID) (*model.Product, error)

	// GetReviewsByProductID returns all reviews for a given product ID.
	GetReviewsByProductID(id uuid.UUID) ([]*model.Review, error)

	// ListProducts returns a paginated list of products.
	ListProducts(page, limit int) ([]*model.Product, error)
}

// ProductService implements ProductServiceAPI using JSON files as data sources.
type ProductService struct {
	productFilePath string
	reviewFilePath  string
}

// NewProductService creates a new ProductService with paths to product and review files.
func NewProductService(productFilePath, reviewFilePath string) *ProductService {
	return &ProductService{
		productFilePath: productFilePath,
		reviewFilePath:  reviewFilePath,
	}
}

// GetProductByID returns a product by its ID.
//
// It loads all products from a JSON file and searches for a product
// with the given UUID. If found, it returns the product. Otherwise,
// it returns an error.
func (s *ProductService) GetProductByID(id uuid.UUID) (*model.Product, error) {
	var products []*model.Product

	log.Debug("loading products from file", zap.String("file", s.productFilePath))
	if err := utils.LoadJSON(s.productFilePath, &products); err != nil {
		log.Error("error loading json file", zap.Error(err))
		return nil, errors.New("application error")
	}

	log.Debug("searching product by id", zap.String("id", id.String()))
	for _, product := range products {
		if product.ID == id {
			log.Info("product found", zap.String("id", id.String()))
			return product, nil
		}
	}

	log.Error("product not found", zap.String("id", id.String()))
	return nil, errors.New("product not found")
}

// GetReviewsByProductID returns all reviews for a given product ID.
//
// It loads reviews from a JSON file and filters those that match
// the provided product ID. If no reviews are found, it returns an error.
func (s *ProductService) GetReviewsByProductID(id uuid.UUID) ([]*model.Review, error) {
	var reviews []*model.Review

	log.Debug("loading reviews from file", zap.String("file", s.reviewFilePath))
	if err := utils.LoadJSON(s.reviewFilePath, &reviews); err != nil {
		log.Error("error loading json file", zap.Error(err))
		return nil, errors.New("application error")
	}

	var productReviews []*model.Review
	for _, review := range reviews {
		if review.ProductID == id {
			productReviews = append(productReviews, review)
		}
	}

	if len(productReviews) == 0 {
		log.Error("not found any review", zap.String("id", id.String()))
		return nil, errors.New("not found any review")
	}

	return productReviews, nil
}

// ListProducts returns a paginated list of products.
//
// It loads all products from a JSON file and returns the slice
// corresponding to the requested page and limit. If the page is out
// of range, it returns an error.
func (s *ProductService) ListProducts(page, limit int) ([]*model.Product, error) {
	var products []*model.Product

	log.Debug("loading products from file", zap.String("file", s.productFilePath))
	if err := utils.LoadJSON(s.productFilePath, &products); err != nil {
		log.Error("error loading json file", zap.Error(err))
		return nil, errors.New("application error")
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(products) {
		log.Error("page out of range", zap.Int("page", page), zap.Int("limit", limit))
		return nil, errors.New("bad request, page out of range")
	}

	if end > len(products) {
		end = len(products)
	}

	return products[start:end], nil
}
