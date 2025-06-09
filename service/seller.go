package service

import (
	"errors"
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/EnricoPDG/meli-desafio/utils"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// SellerServiceAPI defines methods to interact with sellers.
type SellerServiceAPI interface {
	// GetSellerByID returns a seller by its ID.
	GetSellerByID(id uuid.UUID) (*model.Seller, error)
}

// SellerService implements SellerServiceAPI using a JSON file as data source.
type SellerService struct {
	sellerFilePath string
}

// NewSellerService creates a new SellerService with the given JSON file path.
func NewSellerService(sellerFilePath string) *SellerService {
	return &SellerService{
		sellerFilePath: sellerFilePath,
	}
}

// GetSellerByID returns a seller by its ID.
//
// It loads all sellers from a JSON file and searches for a seller
// with the given UUID. If found, it returns the seller. Otherwise,
// it returns an error.
func (s *SellerService) GetSellerByID(id uuid.UUID) (*model.Seller, error) {
	var sellers []*model.Seller

	log.Debug("loading seller from file", zap.String("file", s.sellerFilePath))
	if err := utils.LoadJSON(s.sellerFilePath, &sellers); err != nil {
		log.Error("error loading json file", zap.Error(err))
		return nil, errors.New("application error")
	}

	log.Debug("searching seller by id", zap.String("id", id.String()))
	for _, seller := range sellers {
		if seller.ID == id {
			log.Info("seller found", zap.String("id", id.String()))
			return seller, nil
		}
	}

	log.Error("seller not found", zap.String("id", id.String()))
	return nil, errors.New("seller not found")
}
