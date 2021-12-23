//go:generate go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0
//go:generate oapi-codegen -generate "types,client" -package "internal" -o ./internal/oapi_client.gen.go ./swagger.yaml
//go:generate goimports -w ./internal/oapi_client.gen.go
package resolume
