package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/FlowingSPDG/resolume-go"
)

func main() {
	// Create a new client
	client, err := resolume.NewClient("localhost", "8080")
	if err != nil {
		log.Fatal(err)
	}

	// Get dummy thumbnail
	dummyThumbnail, err := client.GetDummyThumbnail()
	if err != nil {
		log.Fatal(err)
	}
	defer dummyThumbnail.Close()

	// Save dummy thumbnail to file
	dummyFile, err := os.Create("dummy_thumbnail.png")
	if err != nil {
		log.Fatal(err)
	}
	defer dummyFile.Close()

	if _, err := io.Copy(dummyFile, dummyThumbnail); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved dummy thumbnail to dummy_thumbnail.png")

	// Get clip thumbnail
	clipThumbnail, err := client.GetClipThumbnail(1, 1) // Layer 1, Clip 1
	if err != nil {
		log.Fatal(err)
	}
	defer clipThumbnail.Close()

	// Save clip thumbnail to file
	clipFile, err := os.Create("clip_thumbnail.png")
	if err != nil {
		log.Fatal(err)
	}
	defer clipFile.Close()

	if _, err := io.Copy(clipFile, clipThumbnail); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved clip thumbnail to clip_thumbnail.png")

	// Set custom thumbnail for a clip
	customThumbnail, err := os.Open("custom_thumbnail.png")
	if err != nil {
		log.Fatal(err)
	}
	defer customThumbnail.Close()

	if err := client.SetClipThumbnail(1, 1, customThumbnail); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Set custom thumbnail for Layer 1, Clip 1")

	// Reset thumbnail to default
	if err := client.ResetClipThumbnail(1, 1); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Reset thumbnail for Layer 1, Clip 1 to default")
}
