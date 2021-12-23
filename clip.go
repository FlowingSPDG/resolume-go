package resolume

import (
	"context"
	"io"
)

// GetProduct Retrieve product information
func (c *Client) GetClipThumbnailByID(ctx context.Context, cID int64) ([]byte, error) {
	res, err := c.c.ListClipThumbnailById(ctx, cID)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)

	return b, nil
}
