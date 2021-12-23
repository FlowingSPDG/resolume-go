package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"image/png"

	"github.com/FlowingSPDG/resolume-go"
)

var flagClipID int64

// go run example/thumbnail/main.go -clipid 16213427XXXXX

func main() {
	flag.Int64Var(&flagClipID, "clipid", 0, "Clip ID")
	flag.Parse()
	c, err := resolume.NewClient("localhost", "8080")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// Retrieve all clip IDs
	thumbnail, err := c.GetClipThumbnailByID(ctx, flagClipID)
	if err != nil {
		panic(err)
	}

	// Reader
	r := bytes.NewReader(thumbnail)
	im, err := png.Decode(r)
	if err != nil {
		panic(err)
	}
	fmt.Println("Image Size:", im.Bounds().Size())
	fmt.Println("Binary Size:", len(thumbnail))
}
