//go:generate go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.9.0
//go:generate oapi-codegen -generate "types,client" -package "internal" -o ./internal/resolume_oapi.gen.go ./swagger.yaml
//go:generate goimports -w ./internal/resolume_oapi.gen.go
package resolume
