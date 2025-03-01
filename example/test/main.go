package main

import (
	"fmt"
	"log"

	"github.com/FlowingSPDG/resolume-go"
)

func main() {
	// Create a new Resolume client
	client, err := resolume.NewClient("localhost", "8080")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Test 1: Get Product Information
	fmt.Println("=== Testing Product Info ===")
	if err := testProductInfo(client); err != nil {
		log.Printf("Product Info test failed: %v", err)
	}

	// Test 2: Get Effects
	fmt.Println("\n=== Testing Effects ===")
	if err := testEffects(client); err != nil {
		log.Printf("Effects test failed: %v", err)
	}

	// Test 3: Get Sources
	fmt.Println("\n=== Testing Sources ===")
	if err := testSources(client); err != nil {
		log.Printf("Sources test failed: %v", err)
	}
}

func testProductInfo(client *resolume.Client) error {
	product, err := client.GetProduct()
	if err != nil {
		return fmt.Errorf("failed to get product info: %v", err)
	}

	fmt.Printf("Product: %s\n", product.Name)
	fmt.Printf("Version: %d.%d.%d (revision %d)\n",
		product.Major,
		product.Minor,
		product.Micro,
		product.Revision)
	return nil
}

func testEffects(client *resolume.Client) error {
	effects, err := client.GetEffects()
	if err != nil {
		return fmt.Errorf("failed to get effects: %v", err)
	}

	fmt.Printf("Found %d video effects\n", len(effects.Video))
	for i, effect := range effects.Video {
		fmt.Printf("%d. %s (ID: %s)\n", i+1, effect.Name, effect.IDString)
		if len(effect.Presets) > 0 {
			fmt.Printf("   Presets: %d available\n", len(effect.Presets))
			for j, preset := range effect.Presets {
				if j >= 3 { // Only show first 3 presets
					fmt.Printf("   ... and %d more presets\n", len(effect.Presets)-3)
					break
				}
				fmt.Printf("   - %s (ID: %d)\n", preset.Name, preset.ID)
			}
		}
	}
	return nil
}

func testSources(client *resolume.Client) error {
	sources, err := client.GetSources()
	if err != nil {
		return fmt.Errorf("failed to get sources: %v", err)
	}

	fmt.Printf("Found %d video sources\n", len(sources.Video))
	for i, source := range sources.Video {
		fmt.Printf("%d. %s (ID: %s)\n", i+1, source.Name, source.IDString)
		if len(source.Presets) > 0 {
			fmt.Printf("   Presets: %d available\n", len(source.Presets))
			for j, preset := range source.Presets {
				if j >= 3 { // Only show first 3 presets
					fmt.Printf("   ... and %d more presets\n", len(source.Presets)-3)
					break
				}
				fmt.Printf("   - %s (ID: %d)\n", preset.Name, preset.ID)
			}
		}
	}
	return nil
}
