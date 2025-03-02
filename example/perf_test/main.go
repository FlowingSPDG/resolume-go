package main

import (
	"context"
	"log"
	"time"

	"github.com/FlowingSPDG/resolume-go"
)

func main() {
	// Create a new client
	client, err := resolume.NewClient("localhost", "8080")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	log.Println("Starting test...")
	for range 1000 {
		for i := range 3 {
			if err := client.SelectLayerClip(ctx, 1, i+1); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Millisecond * 50)
		}
	}
}
