module github.com/kosenkovd/go-http-grpc/product-service

go 1.16

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/jackc/pgx/v4 v4.13.0
	github.com/kosenkovd/go-http-grpc/category-service/pkg/category-service v0.0.0-20211030120227-a15648ef59e4
	github.com/kosenkovd/go-http-grpc/product-service/pkg/product-service v0.0.1
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.3
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rs/zerolog v1.24.0
	google.golang.org/genproto v0.0.0-20211021150943-2b146023228c // indirect
	google.golang.org/grpc v1.41.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/kosenkovd/go-http-grpc/product-service/pkg/product-service => ./pkg/product-service
