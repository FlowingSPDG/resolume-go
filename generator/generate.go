//go:generate go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0
//go:generate oapi-codegen -generate "types" -package "internal" -o ../internal/oapi_types.gen.go ../swagger.yaml
//go:generate oapi-codegen -generate "client" -package "internal" -o ../internal/oapi_client.gen.go ../swagger.yaml
//go:generate goimports -w ../internal/oapi_types.gen.go
//go:generate goimports -w ../internal/oapi_client.gen.go

//go:generate go run ./models/main.go -src ../internal/oapi_types.gen.go -dest ../models/models.gen.go
//go:generate goimports -w ../models/models.gen.go
package resolume
