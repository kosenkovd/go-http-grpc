package main

import (
	"flag"
	category_service "github.com/kosenkovd/go-http-grpc/category-service/pkg/category-service"
	"github.com/kosenkovd/go-http-grpc/product-service/internal/client/category"
	product2 "github.com/kosenkovd/go-http-grpc/product-service/internal/repository/product"
	"github.com/kosenkovd/go-http-grpc/product-service/internal/service/product"
	"google.golang.org/grpc"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/kosenkovd/go-http-grpc/product-service/internal/config"
	"github.com/kosenkovd/go-http-grpc/product-service/internal/server"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	connection, err := grpc.Dial(cfg.CategoryService.BasePath, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	categoryService := product.New(
		product2.New(),
		category.NewCategoryClient(
			category_service.NewCategoryServiceClient(connection)))

	if err := server.NewGrpcServer(categoryService).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
