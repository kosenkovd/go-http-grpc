package category

import (
	"context"
	category_service "github.com/kosenkovd/go-http-grpc/category-service/pkg/category-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ICategoryClient interface {
	DoesCategoryExist(ctx context.Context, categoryID uint64) (bool, error)
}

func NewCategoryClient(client category_service.CategoryServiceClient) ICategoryClient {
	return &CategoryClient{
		client: client,
	}
}

type CategoryClient struct {
	client category_service.CategoryServiceClient
}

func (c *CategoryClient) DoesCategoryExist(ctx context.Context, categoryID uint64) (bool, error) {
	_, err := c.client.GetCategoryByIdV1(ctx, &category_service.GetCategoryByIdV1Request{
		Id: categoryID,
	})

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
