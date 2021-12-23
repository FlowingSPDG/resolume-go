package resolume

import (
	"context"

	"github.com/FlowingSPDG/resolume-go/models"
)

// GetProduct Retrieve product information
func (c *Client) GetProduct(ctx context.Context) (*models.ProductInfo, error) {
	res, err := c.c.GetProductWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	return (*models.ProductInfo)(res.JSON200), nil
}
