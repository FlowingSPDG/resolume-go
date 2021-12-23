package resolume

import (
	"context"

	"github.com/FlowingSPDG/resolume-go/internal"
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

// UpdateComposition Update the complete composition
func (c *Client) UpdateComposition(ctx context.Context, body models.ReplaceCompositionJSONRequestBody) error {
	if _, err := c.c.ReplaceCompositionWithResponse(ctx, internal.ReplaceCompositionJSONRequestBody(body)); err != nil {
		return err
	}
	return nil
}

func (c *Client) ResetComposition(ctx context.Context, param string, body models.ResetCompositionParamJSONRequestBody) error {
	if _, err := c.c.ResetCompositionParamWithResponse(ctx, param, internal.ResetCompositionParamJSONRequestBody(body)); err != nil {
		return err
	}

	return nil
}
