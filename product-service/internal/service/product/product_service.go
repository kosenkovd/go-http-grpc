package product

import (
	"context"
	"github.com/kosenkovd/product-service/internal/client/category"
	product2 "github.com/kosenkovd/product-service/internal/model/product"
	internal_errors "github.com/kosenkovd/product-service/internal/pkg/errors"
	"github.com/kosenkovd/product-service/internal/repository/product"
)

type Service struct {
	repository            product.ProductRepository
	categoryServiceClient category.ICategoryClient
}

func New(
	repository product.ProductRepository,
	categoryServiceClient category.ICategoryClient,
	) Service {
	return Service{
		repository: repository,
		categoryServiceClient: categoryServiceClient,
	}
}

func (s Service) CreateProduct(
	ctx context.Context,
	name string,
	categoryId uint64) (*product2.Product, error) {

	exists, err := s.categoryServiceClient.DoesCategoryExist(ctx, categoryId)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, internal_errors.ErrNoCategory
	}

	newProduct := &product2.Product{
		ID:         1,
		Name:       name,
		CategoryID: categoryId,
	}

	err = s.repository.SaveProduct(ctx, newProduct)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}
