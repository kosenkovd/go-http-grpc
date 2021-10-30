package product_service

import (
	"github.com/kosenkovd/go-http-grpc/product-service/internal/service/product"
	desc "github.com/kosenkovd/go-http-grpc/product-service/pkg/product-service"
)

type Implementation struct {
	desc.UnimplementedProductServiceServer
	productService product.Service
}

func NewProductService(productService product.Service) desc.ProductServiceServer {
	return &Implementation{
		productService: productService,
	}
}
