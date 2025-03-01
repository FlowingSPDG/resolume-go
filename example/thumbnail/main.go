package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/FlowingSPDG/resolume-go"
)

func main() {
	// Create a new Resolume client
	client, err := resolume.NewClient("localhost", "8080")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Test 1: Get Dummy Thumbnail
	fmt.Println("=== Testing Dummy Thumbnail ===")
	if err := testDummyThumbnail(client); err != nil {
		log.Printf("Dummy thumbnail test failed: %v", err)
	}

	// Test 2: Get Clip Thumbnail
	fmt.Println("\n=== Testing Clip Thumbnail ===")
	if err := testClipThumbnail(client, 1, 1); err != nil { // Layer 1, Clip 1
		log.Printf("Clip thumbnail test failed: %v", err)
	}

	// Test 3: Set Custom Thumbnail
	fmt.Println("\n=== Testing Custom Thumbnail ===")
	if err := testSetCustomThumbnail(client, 1, 1, "custom_thumbnail.png"); err != nil {
		log.Printf("Set custom thumbnail test failed: %v", err)
	}

	// Test 4: Reset Thumbnail
	fmt.Println("\n=== Testing Reset Thumbnail ===")
	if err := testResetThumbnail(client, 1, 1); err != nil {
		log.Printf("Reset thumbnail test failed: %v", err)
	}
}

func testDummyThumbnail(client *resolume.Client) error {
	// Get dummy thumbnail data
	resp, err := client.GetDummyThumbnail()
	if err != nil {
		return fmt.Errorf("failed to get dummy thumbnail: %v", err)
	}

	// Save dummy thumbnail to file
	if err := saveThumbnail(resp, "dummy_thumbnail.png"); err != nil {
		return fmt.Errorf("failed to save dummy thumbnail: %v", err)
	}

	fmt.Println("Dummy thumbnail saved as dummy_thumbnail.png")
	return nil
}

func testClipThumbnail(client *resolume.Client, layerIndex, clipIndex int) error {
	// Get clip thumbnail
	resp, err := client.GetClipThumbnail(layerIndex, clipIndex)
	if err != nil {
		return fmt.Errorf("failed to get clip thumbnail: %v", err)
	}

	// Save clip thumbnail to file
	filename := fmt.Sprintf("clip_%d_%d_thumbnail.png", layerIndex, clipIndex)
	if err := saveThumbnail(resp, filename); err != nil {
		return fmt.Errorf("failed to save clip thumbnail: %v", err)
	}

	fmt.Printf("Clip thumbnail saved as %s\n", filename)
	return nil
}

func testSetCustomThumbnail(client *resolume.Client, layerIndex, clipIndex int, thumbnailPath string) error {
	// Open custom thumbnail file
	file, err := os.Open(thumbnailPath)
	if err != nil {
		return fmt.Errorf("failed to open thumbnail file: %v", err)
	}
	defer file.Close()

	// Set custom thumbnail
	if err := client.SetClipThumbnail(layerIndex, clipIndex, file); err != nil {
		return fmt.Errorf("failed to set custom thumbnail: %v", err)
	}

	fmt.Printf("Custom thumbnail set for layer %d, clip %d\n", layerIndex, clipIndex)
	return nil
}

func testResetThumbnail(client *resolume.Client, layerIndex, clipIndex int) error {
	// Reset thumbnail to default
	if err := client.ResetClipThumbnail(layerIndex, clipIndex); err != nil {
		return fmt.Errorf("failed to reset thumbnail: %v", err)
	}

	fmt.Printf("Thumbnail reset for layer %d, clip %d\n", layerIndex, clipIndex)
	return nil
}

func saveThumbnail(data io.Reader, filename string) error {
	// Create output directory if it doesn't exist
	outputDir := "thumbnails"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Create output file
	outputPath := filepath.Join(outputDir, filename)
	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer out.Close()

	// Copy thumbnail data to file
	if _, err := io.Copy(out, data); err != nil {
		return fmt.Errorf("failed to write thumbnail data: %v", err)
	}

	return nil
}
