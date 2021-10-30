package product

import (
	"context"
	"github.com/kosenkovd/product-service/internal/model/product"
)

type Repository struct {
	index uint64
}

func New() ProductRepository {
	return &Repository{
		index: 1,
	}
}

func (r *Repository) SaveProduct(ctx context.Context, product *product.Product) error {
	product.ID = r.index
	r.index++

	return nil
}

type ProductRepository interface {
	SaveProduct(ctx context.Context, product *product.Product) error
}
