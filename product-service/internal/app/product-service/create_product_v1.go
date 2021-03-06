package product_service

import (
	"context"
	"errors"
	"github.com/kosenkovd/go-http-grpc/product-service/internal/model/product"
	internal_errors "github.com/kosenkovd/go-http-grpc/product-service/internal/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"

	desc "github.com/kosenkovd/go-http-grpc/product-service/pkg/product-service"
)

func (i *Implementation) CreateProductV1(ctx context.Context,
	req *desc.CreateProductV1Request) (*desc.CreateProductV1Response, error) {
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("CreateProductV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	prod, err := i.productService.CreateProduct(ctx, req.GetName(), req.GetCategoryId())
	if err != nil {
		if errors.Is(err, internal_errors.ErrNoCategory) {
			details := &errdetails.BadRequest{
				FieldViolations: []*errdetails.BadRequest_FieldViolation{
					{
						Field:       "categoryId",
						Description: "category with id " + strconv.FormatUint(req.GetCategoryId(), 10) + " does not exist",
					},
				},
			}

			st, stErr := status.New(codes.InvalidArgument, "wrong category").WithDetails(details)
			if stErr != nil {
				return nil, stErr
			}

			return nil, st.Err()
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return convertToProductResponse(prod), nil
}

func convertToProductResponse(product *product.Product) *desc.CreateProductV1Response {
	return &desc.CreateProductV1Response{
		Result: &desc.Product{
			Id:         product.ID,
			Name:       product.Name,
			CategoryId: product.CategoryID,
		},
	}
}
