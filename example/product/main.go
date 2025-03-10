package main

import (
	"fmt"
	"log"

	"github.com/FlowingSPDG/resolume-go"
)

func main() {
	// Create a new client
	client, err := resolume.NewClient("localhost", "8080")
	if err != nil {
		log.Fatal(err)
	}

	// Get product information
	product, err := client.GetProduct()
	if err != nil {
		log.Fatal(err)
	}

	// Print product information
	fmt.Printf("Product: %s\n", product.Name)
	fmt.Printf("Version: %d.%d.%d (revision: %d)\n",
		product.Major,
		product.Minor,
		product.Micro,
		product.Revision)
}
