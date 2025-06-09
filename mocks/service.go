package mocks

import (
	"github.com/EnricoPDG/meli-desafio/model"
	"github.com/google/uuid"
)

type ProductServiceMock struct {
	Error error
	Data  interface{}
}

type SellerServiceMock struct {
	Error error
	Data  interface{}
}

func (m *ProductServiceMock) GetProductByID(id uuid.UUID) (*model.Product, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	data, ok := m.Data.(*model.Product)
	if !ok {
		return nil, m.Error
	}

	return data, nil
}

func (m *ProductServiceMock) GetReviewsByProductID(id uuid.UUID) ([]*model.Review, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	data, ok := m.Data.([]*model.Review)
	if !ok {
		return nil, m.Error
	}

	return data, nil
}

func (m *ProductServiceMock) ListProducts(page, limit int) ([]*model.Product, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	data, ok := m.Data.([]*model.Product)
	if !ok {
		return nil, m.Error
	}

	return data, nil
}

func (m *SellerServiceMock) GetSellerByID(id uuid.UUID) (*model.Seller, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	data, ok := m.Data.(*model.Seller)
	if !ok {
		return nil, m.Error
	}

	return data, nil
}
