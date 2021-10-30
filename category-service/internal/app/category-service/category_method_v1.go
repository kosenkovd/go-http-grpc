package category_service

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	desc "github.com/kosenkovd/go-http-grpc/category-service/pkg/category-service"
)

func (i *Implementation) CategoryMethodV1(
	ctx context.Context,
	req *desc.CategoryMethodV1Request,
) (*desc.CategoryMethodV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CategoryMethodV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &desc.CategoryMethodV1Response{
		Value: &desc.Template{
			Id:   req.GetId(),
			Name: "Category",
		},
	}, nil
}
