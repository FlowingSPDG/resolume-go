package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FlowingSPDG/resolume-go"
)

func main() {
	ctx := context.Background()

	// Create a new client
	client, err := resolume.NewClient("localhost", "8080")
	if err != nil {
		log.Fatal(err)
	}

	// Get available effects
	effects, err := client.GetEffects(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print available video effects
	fmt.Println("Available Video Effects:")
	for _, effect := range effects.Video {
		fmt.Printf("- %s (ID: %s)\n", effect.Name, effect.IDString)
		if len(effect.Presets) > 0 {
			fmt.Println("  Presets:")
			for _, preset := range effect.Presets {
				fmt.Printf("  - %s (ID: %d)\n", preset.Name, preset.ID)
			}
		}
	}

	fmt.Println()

	// Get available sources
	sources, err := client.GetSources(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print available video sources
	fmt.Println("Available Video Sources:")
	for _, source := range sources.Video {
		fmt.Printf("- %s (ID: %s)\n", source.Name, source.IDString)
		if len(source.Presets) > 0 {
			fmt.Println("  Presets:")
			for _, preset := range source.Presets {
				fmt.Printf("  - %s (ID: %d)\n", preset.Name, preset.ID)
			}
		}
	}
}
