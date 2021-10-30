package category

import (
	"context"
	internal_errors "github.com/kosenkovd/go-http-grpc/category-service/internal/pkg/errors"
	"github.com/kosenkovd/go-http-grpc/category-service/internal/service/category/model"
	"github.com/kosenkovd/go-http-grpc/category-service/internal/service/category/repository"
	"github.com/pkg/errors"
)

type Service struct {
	repository repository.CategoryRepository
}

func New(repository repository.CategoryRepository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) GetCategoryByID(ctx context.Context, id uint64) (*model.Category, error)  {
	categories, err := s.repository.FindAllCategories(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.FindAllCategories")
	}

	cat := categories.FilterByID(id)
	if cat == nil {
		return nil, internal_errors.ErrNoCategory
	}

	return cat, nil
}