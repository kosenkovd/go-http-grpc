package repository

import (
	"context"
	"github.com/kosenkovd/category-service/internal/service/category/model"
)

var categories = model.Categories{
	{
		ID: 1,
		Name: "Toys",
	},
	{
		ID: 2,
		Name: "Games",
	},
	{
		ID: 3,
		Name: "Mayhem",
	},
	{
		ID: 4,
		Name: "Murder",
	},
}

type Repository struct {
}

func New() CategoryRepository {
	return &Repository{}
}

func (r *Repository) FindAllCategories(ctx context.Context) (model.Categories, error) {
	return categories, nil
}

type CategoryRepository interface {
	FindAllCategories(ctx context.Context) (model.Categories, error)
}
