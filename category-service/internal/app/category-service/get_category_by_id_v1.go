package category_service

import (
	"context"
	"errors"
	internal_errors "github.com/kosenkovd/category-service/internal/pkg/errors"
	errorWrapper "github.com/pkg/errors"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/kosenkovd/category-service/pkg/category-service"
)

func (i *Implementation) GetCategoryByIdV1(
	ctx context.Context,
	req *desc.GetCategoryByIdV1Request,
) (*desc.GetCategoryByIdV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CategoryMethodV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	cat, err := i.categoryService.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, internal_errors.ErrNoCategory) {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.NotFound, errorWrapper.Wrap(err, "categoryService.GetCategoryByID").Error())
	}

	return &desc.GetCategoryByIdV1Response{
		Value: &desc.Template{
			Id:   cat.ID,
			Name: cat.Name,
		},
	}, nil
}

