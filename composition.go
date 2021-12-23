package resolume

import (
	"context"

	"github.com/FlowingSPDG/resolume-go/models"
)

// ListComposition Retrieve the complete composition
func (c *Client) ListComposition(ctx context.Context) (*models.Composition, error) {
	res, err := c.c.ListCompositionWithResponse(ctx)
	if err != nil {
		return nil, err
	}

	return (*models.Composition)(res.JSON200), nil
}
