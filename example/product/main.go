package main

import (
	"context"
	"fmt"

	"github.com/FlowingSPDG/resolume-go"
)

func main() {
	c, err := resolume.NewClient("localhost", "8080")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	product, err := c.GetProduct(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("Product:", *product.Name)
	fmt.Printf("Version: %d.%d.%d\n", *product.Major, *product.Minor, *product.Micro)
	fmt.Println("Revision:", *product.Revision)
}
